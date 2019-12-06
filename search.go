package radarapi

import (
	"fmt"
	lang "golang.org/x/text/language"
	"log"
	"net/url"
)

const (
	searchEventUrl    = baseUrl + "search/events.json"
	searchGroupUrl    = baseUrl + "%search/groups.json"
	searchLocationUrl = baseUrl + "search/location.json"
	searchTermUrl     = baseUrl + "search/term.json"
)

const (
	// Entity types
	EVENT = iota
	GROUP
	LOCATION
	TERM
)

type search struct {
	entity   int
	key      string
	language lang.Tag
	limit    uint16
	sort     string
	desc     bool
	facets   []Facet
	fields   []string
	filters  []Filter
}

// A SearchBuilder is used to create a search request which can be passed to RadarClient.Search().
//
// Don't use this directly. Create a new SearchBuilder with NewSearch().
type SearchBuilder struct {
	search *search
}

// Creates a new SearchBuilder which allows searching for the given entity.
//
// Possible values are EVENT, GROUP, LOCATION, TERM (see Constants above).
func NewSearch(entity int) *SearchBuilder {
	return &SearchBuilder{search: &search{entity: entity}}
}

// Sets a key for full-text search.
func (sb *SearchBuilder) Key(key string) {
	sb.search.key = key
}

// Sets the language for the response.
func (sb *SearchBuilder) Language(language lang.Tag) {
	sb.search.language = language
}

// Sets the amount of results returned.
//
// Must be 0 <= limit <= 500
func (sb *SearchBuilder) Limit(limit uint16) {
	if limit > 500 {
		log.Printf("warning: max. limit of 500 exceeded :%d. Limit set to 500", limit)
	} else {
		sb.search.limit = limit
	}
}

// Sets the field used to sort the results.
//
// Only works for indexed fields.
func (sb *SearchBuilder) Sort(field string, descending bool) {
	sb.search.sort = field
	sb.search.desc = descending
}

// Sets the facets which will be used to filter the results. See Facet above.
func (sb *SearchBuilder) Facets(facets ...Facet) {
	if facets != nil {
		sb.search.facets = facets
	}
}
// Sets the fields which will be returned for each result.
//
// Fields must come from the entity which has been passed to NewSearch(),
// e.g. for EVENT event.FieldXXXX. See subpackages for the available fields for each entity.
func (sb *SearchBuilder) Fields(fields ...string) {
	if fields != nil {
		sb.search.fields = fields
	}
}

// Sets the filters which will be used to filter the results. See Filter above.
func (sb *SearchBuilder) Filters(filters ...Filter) {
	if filters != nil {
		sb.search.filters = filters
	}
}

func prepareSearchUrl(sb *SearchBuilder) (*url.URL, error) {
	var addr string
	switch sb.search.entity {
	case EVENT:
		addr = searchEventUrl
	case GROUP:
		addr = searchGroupUrl
	case LOCATION:
		addr = searchLocationUrl
	case TERM:
		addr = searchTermUrl
	default:
		return nil, fmt.Errorf("error: entity must be one of EVENT, GROUP, LOCATION or TERM")
	}
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if sb.search.facets != nil {
		for _, facet := range sb.search.facets {
			query.Add(fmt.Sprintf("facets[%s][]", facet.Key), facet.Value)
		}
	}
	if sb.search.filters != nil {
		for _, f := range sb.search.filters {
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
	if sb.search.limit > 0 {
		query.Set("limit", string(sb.search.limit))
	}
	if sb.search.key != "" {
		query.Set("keys", sb.search.key)
	}
	if len(sb.search.fields) != 0 {
		query.Set("fields", fieldsToCommaString(sb.search.fields))
	}
	u.RawQuery = query.Encode()
	return u, nil
}
