package radarapi

import (
	"0xacab.org/radarapi/event"
	"0xacab.org/radarapi/group"
	"0xacab.org/radarapi/location"
	"0xacab.org/radarapi/term"
	"encoding/json"
	"fmt"
	"golang.org/x/text/language"
)

// Get all events in Berlin and include the address field of the locations.
func ExampleResolveField() {
	// https://radar.squat.net/api/1.2/search/events.json?facets[city][]=Berlin&fields=title,offline,offline:address
	radar := NewRadarClient()
	sb := NewSearch(EVENT)
	sb.Facets(Facet{event.FacetCity, "Berlin"})
	sb.Fields(event.FieldTitle, event.FieldOffline, ResolveField(event.FieldOffline, location.FieldAddress))
	result, err := radar.Search(sb)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	var events map[string]interface{}
	err = json.Unmarshal([]byte(result), &events)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("%v", events)
}

// Get event "e81b15f9-663f-4270-b94b-269176cd9f3f".
func ExampleRadarClient_Event() {
	// https://radar.squat.net/api/1.2/node/e81b15f9-663f-4270-b94b-269176cd9f3f.json
	radar := NewRadarClient()
	result, err := radar.Event("e81b15f9-663f-4270-b94b-269176cd9f3f", nil)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	var ev map[string]interface{}
	err = json.Unmarshal([]byte(result), &ev)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("%v", ev[event.FieldTitle])
	// Output: Seawatch Soli Fete Tag 1
}

// Get event "e81b15f9-663f-4270-b94b-269176cd9f3f" with fields title and event status.
func ExampleRadarClient_Event_fields() {
	// https://radar.squat.net/api/1.2/node/e81b15f9-663f-4270-b94b-269176cd9f3f.json?fields=title,event_status
	radar := NewRadarClient()
	result, err := radar.Event("e81b15f9-663f-4270-b94b-269176cd9f3f", nil, event.FieldTitle, event.FieldEventStatus)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	var ev map[string]interface{}
	err = json.Unmarshal([]byte(result), &ev)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("title: %s, status: %s", ev[event.FieldTitle], ev[event.FieldEventStatus])
	// Output: title: Seawatch Soli Fete Tag 1, status: confirmed
}

// Get group "Liebig34" : "7a8de6aa-5951-467f-a4a4-1d558eea5d3c" and print title and description.
func ExampleRadarClient_Group() {
	// https://radar.squat.net/api/1.2/node/7a8de6aa-5951-467f-a4a4-1d558eea5d3c.json
	radar := NewRadarClient()
	result, err := radar.Group("7a8de6aa-5951-467f-a4a4-1d558eea5d3c", nil)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	var grp map[string]interface{}
	err = json.Unmarshal([]byte(result), &grp)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	title := grp[group.FieldTitle]
	body := grp[group.FieldBody].(map[string]interface{})
	fmt.Printf("title: %s, description: %s", title, body["value"])
	// Output: title: Liebig34, description: <p>Anarcha-Queer-Feminist houseproject in Berlin-Friedrichshain.</p>
	//<p>Non-smoking bar.</p>
	//<p> </p>
}

// Get location "2130ff33-8af2-4c76-a693-75d3b978a2aa" and print name and street.
func ExampleRadarClient_Location() {
	// https://radar.squat.net/api/1.2/location/2130ff33-8af2-4c76-a693-75d3b978a2aa.json
	radar := NewRadarClient()
	result, err := radar.Location("2130ff33-8af2-4c76-a693-75d3b978a2aa", nil)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	var loc map[string]interface{}
	err = json.Unmarshal([]byte(result), &loc)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	address := loc[location.FieldAddress].(map[string]interface{})
	fmt.Printf("name: %s, street: %s", address["name_line"], address["thoroughfare"])
	// Output: name: Mensch Meier, street: Storkower Str. 121
}

// Get taxonomy term "2a56c4d7-eb98-4f96-9ac6-d383a1af5ce8" and print name and description.
func ExampleRadarClient_Term() {
	// https://radar.squat.net/api/1.2/taxonomy_term/2a56c4d7-eb98-4f96-9ac6-d383a1af5ce8.json
	radar := NewRadarClient()
	result, err := radar.Term("2a56c4d7-eb98-4f96-9ac6-d383a1af5ce8", nil)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	var t map[string]interface{}
	err = json.Unmarshal([]byte(result), &t)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("name: %s, description: %s", t[term.FieldName], t[term.FieldDescription])
	// Output: name: bar/cafe, description: <p>somewhere were you can go for a drink and to meet people</p>
}

// Get taxonomy term "2a56c4d7-eb98-4f96-9ac6-d383a1af5ce8" and print name and description in spanish.
func ExampleRadarClient_Term_spanish() {
	// https://radar.squat.net/api/1.2/taxonomy_term/2a56c4d7-eb98-4f96-9ac6-d383a1af5ce8.json?language=es
	radar := NewRadarClient()
	result, err := radar.Term("2a56c4d7-eb98-4f96-9ac6-d383a1af5ce8", &language.Spanish)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	var t map[string]interface{}
	err = json.Unmarshal([]byte(result), &t)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("name: %s, description: %s", t[term.FieldName], t[term.FieldDescription])
	// Output: name: bar/café, description: <p>para tomarse una copita, conocer gente maja...</p>
}

// Get all events in Berlin with the category work-space/diy happening on the 2020-11-24
func ExampleFacet() {
	radar := NewRadarClient()
	sb := NewSearch(EVENT)
	sb.Facets(
		Facet{event.FacetCity, "Berlin"},
		Facet{event.FacetCategory, "work-space-diy"},
		Facet{event.FacetDate, "2020-11-24"})
	result, err := radar.Search(sb)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	var events map[string]interface{}
	err = json.Unmarshal([]byte(result), &events)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("%v", events)
}

// Get all events in Berlin which are not organized/promoted by group "Stressfaktor" (id: 1599)
func ExampleCreateFilter() {
	// https://radar.squat.net/api/1.2/search/events.json?facets[city][]=Berlin&filter[~and][og_group_ref][~ne]=1599
	radar := NewRadarClient()
	sb := NewSearch(EVENT)
	sb.Facets(Facet{event.FacetCity, "Berlin"})
	sb.Filters(CreateFilter(OperatorAnd, event.FieldOgGroupRef, ComparatorNEQ, "1599"))
	result, err := radar.Search(sb)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	var events map[string]interface{}
	err = json.Unmarshal([]byte(result), &events)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("%v", events)
}

// Get all events in Berlin with a start date between 2020-11-03 00:00:00 and 2020-11-25 00:00:00
func ExampleCreaterRangeFilter() {
	//  https://radar.squat.net/api/1.2/search/events.json?facets[city][]=Berlin&filter[~and][search_api_aggregation_1][~gt]=1604358000&filter[~or][search_api_aggregation_1][~lt]=1606258800
	radar := NewRadarClient()
	sb := NewSearch(EVENT)
	sb.Facets(Facet{event.FacetCity, "Berlin"})
	sb.Filters(CreaterRangeFilter(FilterEventStartDateTime, "1604358000", "1606258800"))
	result, err := radar.Search(sb)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	var events map[string]interface{}
	err = json.Unmarshal([]byte(result), &events)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("%v", events)
}

// Get all events in Berlin with the category food.
func ExampleRadarClient_Search() {
	// https://radar.squat.net/api/1.2/search/events.json?facets[city][]=Berlin&facets[category][]=food
	radar := NewRadarClient()
	sb := NewSearch(EVENT)
	sb.Facets(Facet{event.FacetCity, "Berlin"}, Facet{event.FacetCategory, "food"})
	result, err := radar.Search(sb)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	var events map[string]interface{}
	err = json.Unmarshal([]byte(result), &events)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("%v", events)
}
