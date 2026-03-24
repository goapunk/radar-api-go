package radarapi

import (
	"encoding/json"
	"fmt"
	"github.com/goapunk/radar-api-go/event"
	"github.com/goapunk/radar-api-go/group"
	"github.com/goapunk/radar-api-go/location"
	"github.com/goapunk/radar-api-go/term"
	"log/slog"
	"net/url"
	"strconv"

	lang "golang.org/x/text/language"
)

const (
	searchEventUrl    = "search/events.json"
	searchGroupUrl    = "search/groups.json"
	searchLocationUrl = "search/location.json"
	searchTermUrl     = "search/term.json"
)

const (
	typeEvent = iota
	typeGroup
	typeLocation
	typeTerm
)

// A SearchBuilder is used to create a search request which can be passed to RadarClient.Search().
// Use NewSearchBuilder() to create this instead of &SearchBuilder.
//
// The zero value is ready to use.
type SearchBuilder struct {
	entity   int
	key      string
	language lang.Tag
	limit    uint64
	sort     string
	desc     bool
	facets   []Facet
	fields   []string
	filters  []Filter
	baseUrl  string
	log      *slog.Logger
}

// The response to a search request will return a list of available facets.
type ResultFacet struct {
	Filter    string `json:"filter"`
	Count     int64  `json:"count"`
	Formatted string `json:"formatted"`
}

func (e *ResultFacet) UnmarshalJSON(data []byte) error {
	if len(data) < 3 {
		return nil
	}
	var buf map[string]interface{}
	err := json.Unmarshal(data, &buf)
	if err != nil {
		return err
	}
	e.Filter = buf["filter"].(string)
	switch buf["count"].(type) {
	case int64:
		e.Count = buf["count"].(int64)
	case int:
		e.Count = int64(buf["count"].(int))
	case string:
		v, err := strconv.ParseInt(buf["count"].(string), 10, 64)
		if err != nil {
			return err
		}
		e.Count = v
	}
	return nil
}

// Wrapper for the count field in the search result, because it can be either a string or an int in the API response.
type FlexInt64 int64

func (fi *FlexInt64) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	var i int64
	if err := json.Unmarshal(b, &i); err == nil {
		*fi = FlexInt64(i)
		return nil
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return json.Unmarshal([]byte(s), &i)
}

type SearchResult[T any] struct {
	RawResults json.RawMessage           `json:"result"`
	Results    map[string]*T             `json:"-"`
	Count      FlexInt64                 `json:"count"`
	Facets     map[string][]*ResultFacet `json:"facets"`
}

func (radar *RadarClient) NewSearchBuilder() *SearchBuilder {
	return &SearchBuilder{
		baseUrl: radar.baseUrl,
		log:     radar.log,
	}
}

// Search the radar database for Events. Returns nil if no results were found.
func (radar *RadarClient) SearchEvents(sb *SearchBuilder) (*SearchResult[event.Event], error) {
	sb.entity = typeEvent
	return performSearch[event.Event](radar, sb)
}

// Search the radar database for Groups. Returns nil if no results were found.
func (radar *RadarClient) SearchGroups(sb *SearchBuilder) (*SearchResult[group.Group], error) {
	sb.entity = typeGroup
	return performSearch[group.Group](radar, sb)
}

// Search the radar database for Locations. Returns nil if no results were found.
//
// Note: By default, no fields are returned for the locations. Use SearchBuilder.Fields() to get a meaningful response.
func (radar *RadarClient) SearchLocation(sb *SearchBuilder) (*SearchResult[location.Location], error) {
	sb.entity = typeLocation
	return performSearch[location.Location](radar, sb)
}

// Search the radar database for Terms. Returns nil if no result were found.
func (radar *RadarClient) SearchTerm(sb *SearchBuilder) (*SearchResult[term.Term], error) {
	sb.entity = typeTerm
	return performSearch[term.Term](radar, sb)
}

func performSearch[T any](radar *RadarClient, sb *SearchBuilder) (*SearchResult[T], error) {
	u, err := prepareSearchUrl(sb)
	if err != nil {
		return nil, err
	}
	raw, err := radar.runQuery(u)
	if err != nil {
		return nil, err
	}

	var res SearchResult[T]
	if err := json.Unmarshal([]byte(raw), &res); err != nil {
		return nil, err
	}

	if string(res.RawResults) == "false" {
		return nil, nil
	}

	if err := json.Unmarshal(res.RawResults, &res.Results); err != nil {
		return nil, fmt.Errorf("error decoding results map: %v", err)
	}

	return &res, nil
}

// Sets a key for full-text search.
func (sb *SearchBuilder) Key(key string) {
	sb.key = key
}

// Sets the language for the response.
func (sb *SearchBuilder) Language(language lang.Tag) {
	sb.language = language
}

// Sets the amount of results returned.
//
// Must be 0 <= limit <= 500
func (sb *SearchBuilder) Limit(limit uint64) {
	if limit > 500 {
		sb.log.Warn(fmt.Sprintf("max. limit of 500 exceeded: %d. Limit set to 500", limit))
	} else {
		sb.limit = limit
	}
}

// Sets the field used to sort the results.
//
// Only works for indexed fields.
func (sb *SearchBuilder) Sort(field string, descending bool) {
	sb.sort = field
	sb.desc = descending
}

// Sets the facets which will be used to filter the results. See Facet above.
func (sb *SearchBuilder) Facets(facets ...Facet) {
	if facets != nil {
		sb.facets = facets
	}
}

// Sets the fields which will be returned for each result.
//
// Fields must come from the entity you are searching for,
// e.g. for Event event.FieldXXXX. See subpackages for the available fields for each entity.
func (sb *SearchBuilder) Fields(fields ...string) {
	if fields != nil {
		sb.fields = fields
	}
}

// Sets the filters which will be used to filter the results. See Filter above.
func (sb *SearchBuilder) Filters(filters ...Filter) {
	if filters != nil {
		sb.filters = filters
	}
}

func prepareSearchUrl(sb *SearchBuilder) (*url.URL, error) {
	var addr string
	switch sb.entity {
	case typeEvent:
		addr = searchEventUrl
	case typeGroup:
		addr = searchGroupUrl
	case typeLocation:
		addr = searchLocationUrl
	case typeTerm:
		addr = searchTermUrl
	}
	u, err := url.Parse(sb.baseUrl + "/" + addr)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if sb.facets != nil {
		for _, facet := range sb.facets {
			query.Add(fmt.Sprintf("facets[%s][]", facet.Key), facet.Value)
		}
	}
	if sb.filters != nil {
		for _, f := range sb.filters {
			switch v := f.(type) {
			case *filter:
				query.Add(fmt.Sprintf("filter[%s][%s][~%s]", v.operator, v.field, v.comparator), v.value)
			case *rangeFilter:
				query.Add(fmt.Sprintf("filter[~and][%s][~gte]", v.field), v.from)
				query.Add(fmt.Sprintf("filter[~or][%s][~lte]", v.field), v.to)
			}
		}
	}
	if sb.limit > 0 {
		query.Set("limit", strconv.FormatUint(sb.limit, 10))
	}
	if sb.key != "" {
		query.Set("keys", sb.key)
	}
	if len(sb.fields) != 0 {
		query.Set("fields", fieldsToCommaString(sb.fields))
	}
	u.RawQuery = query.Encode()
	return u, nil
}
