package radarapi

import (
	"encoding/json"
	"fmt"
	lang "golang.org/x/text/language"
	"log"
	"net/url"
	"strconv"
	"strings"
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

func (radar *RadarClient) NewSearchBuilder() *SearchBuilder {
	return &SearchBuilder{baseUrl: radar.baseUrl}
}

// Search the radar database for Events. Returns nil if no results were found.
func (radar *RadarClient) SearchEvents(sb *SearchBuilder) (*SearchResultEvents, error) {
	sb.entity = typeEvent
	res, err := radar.search(sb)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(*SearchResultEvents), nil
}

// Search the radar database for Groups. Returns nil if no results were found.
func (radar *RadarClient) SearchGroup(sb *SearchBuilder) (*SearchResultGroups, error) {
	sb.entity = typeGroup
	res, err := radar.search(sb)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(*SearchResultGroups), nil
}

// Search the radar database for Locations. Returns nil if no results were found.
//
// Note: By default, no fields are returned for the locations. Use SearchBuilder.Fields() to get a meaningful response.
func (radar *RadarClient) SearchLocation(sb *SearchBuilder) (*SearchResultLocations, error) {
	sb.entity = typeLocation
	res, err := radar.search(sb)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(*SearchResultLocations), nil
}

// Search the radar database for Terms. Returns nil if no result were found.
func (radar *RadarClient) SearchTerm(sb *SearchBuilder) (*SearchResultTerms, error) {
	sb.entity = typeTerm
	res, err := radar.search(sb)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(*SearchResultTerms), nil
}

func (radar *RadarClient) search(sb *SearchBuilder) (interface{}, error) {
	u, err := prepareSearchUrl(sb)
	if err != nil {
		return "", err
	}
	raw, err := radar.runQuery(u)
	if err != nil {
		return nil, err
	}
	if raw == `{"result":false,"count":0,"facets":null}` {
		return nil, nil
	}
	dec := json.NewDecoder(strings.NewReader(raw))
	//	dec.DisallowUnknownFields()
	var buf interface{}
	switch sb.entity {
	case typeEvent:
		buf = &SearchResultEvents{}
	case typeGroup:
		buf = &SearchResultGroups{}
	case typeLocation:
		buf = &SearchResultLocations{}
	case typeTerm:
		buf = &SearchResultTerms{}
	}
	err = dec.Decode(buf)

	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return buf, nil
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
		log.Printf("warning: max. limit of 500 exceeded :%d. Limit set to 500", limit)
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
			switch f.(type) {
			case filter:
				nf := f.(filter)
				query.Add(fmt.Sprintf("filter[%s][%s][~%s]", nf.operator, nf.field, nf.comparator), nf.value)
			case rangeFilter:
				rf := f.(rangeFilter)
				query.Add(fmt.Sprintf("filter[~and][%s][~gte]", rf.field), rf.from)
				query.Add(fmt.Sprintf("filter[~or][%s][~lte]", rf.field), rf.to)
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
