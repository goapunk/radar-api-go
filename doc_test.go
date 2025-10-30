package radarapi

import (
	"fmt"

	"github.com/goapunk/radar-api-go/event"
	"github.com/goapunk/radar-api-go/group"
	"github.com/goapunk/radar-api-go/location"
	"golang.org/x/text/language"
)

// Get all events in Berlin and include the address field of the locations.
func ExampleResolveField() {
	// https://radar.squat.net/api/1.2/search/events.json?facets[city][]=Berlin&fields=title,offline,offline:address
	radar := NewRadarClient()
	sb := radar.NewSearchBuilder()
	sb.Facets(Facet{event.FacetCity, "Berlin"})
	sb.Fields(event.FieldTitle, event.FieldOffline, ResolveField(event.FieldOffline, location.FieldAddress))
	result, err := radar.SearchEvents(sb)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	if result == nil {
		fmt.Println("No results")
	}
	fmt.Printf("%v", result)
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
	fmt.Printf("%v", result.Title)
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
	fmt.Printf("title: %s, status: %s", result.Title, result.EventStatus)
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
	fmt.Printf("title: %s, description: %s", result.Title, result.Body.Value)
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
	fmt.Printf("name: %s, street: %s", result.Address.NameLine, result.Address.Thoroughfare)
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
	fmt.Printf("name: %s, description: %s", result.Name, result.Description)
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
	fmt.Printf("name: %s, description: %s", result.Name, result.Description)
}

// Get all events in Berlin with the category work-space/diy happening on the 2020-11-24
func ExampleFacet() {
	radar := NewRadarClient()
	sb := radar.NewSearchBuilder()
	sb.Facets(
		Facet{event.FacetCity, "Berlin"},
		Facet{event.FacetCategory, "work-space-diy"},
		Facet{event.FacetDate, "2020-11-24"})
	result, err := radar.SearchEvents(sb)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	if result == nil {
		fmt.Println("No results")
	}
	fmt.Printf("%v", result)
}

// Get all events in Berlin which are not organized/promoted by group "Stressfaktor" (id: 1599)
func ExampleCreateFilter() {
	// https://radar.squat.net/api/1.2/search/events.json?facets[city][]=Berlin&filter[~and][og_group_ref][~ne]=1599
	radar := NewRadarClient()
	sb := radar.NewSearchBuilder()
	sb.Facets(Facet{event.FacetCity, "Berlin"})
	sb.Filters(CreateFilter(OperatorAnd, event.FieldOgGroupRef, ComparatorNEQ, "1599"))
	result, err := radar.SearchEvents(sb)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	if result == nil {
		fmt.Println("No results")
	}
	fmt.Printf("%v", result)
}

// Get all events in Berlin with a start date between 2020-11-03 00:00:00 and 2020-11-25 00:00:00
func ExampleCreaterRangeFilter() {
	//  https://radar.squat.net/api/1.2/search/events.json?facets[city][]=Berlin&filter[~and][search_api_aggregation_1][~gt]=1604358000&filter[~or][search_api_aggregation_1][~lt]=1606258800
	radar := NewRadarClient()
	sb := radar.NewSearchBuilder()
	sb.Facets(Facet{event.FacetCity, "Berlin"})
	sb.Filters(CreaterRangeFilter(FilterEventStartDateTime, "1604358000", "1606258800"))
	result, err := radar.SearchEvents(sb)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	if result == nil {
		fmt.Println("No results")
	}
	fmt.Printf("%v", result)
}

// Get all events in Berlin with the category food.
func ExampleRadarClient_search() {
	// https://radar.squat.net/api/1.2/search/events.json?facets[city][]=Berlin&facets[category][]=food
	radar := NewRadarClient()
	sb := radar.NewSearchBuilder()
	sb.Facets(Facet{event.FacetCity, "Berlin"}, Facet{event.FacetCategory, "food"})
	result, err := radar.SearchEvents(sb)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	if result == nil {
		fmt.Println("No results")
	}
	fmt.Printf("%v", result)
}

// Get all groups in Berlin with the category bar/cafe.
func ExampleRadarClient_search_group() {
	// https://radar.squat.net/api/1.2/search/groups.json?facets[city][]=berlin&facets[category][]=bar-cafe
	radar := NewRadarClient()
	sb := radar.NewSearchBuilder()
	sb.Facets(Facet{group.FacetCity, "berlin"}, Facet{group.FacetCategory, "bar-cafe"})
	result, err := radar.SearchGroup(sb)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	if result == nil {
		fmt.Println("No results")
	}
	fmt.Printf("%v", result)
}

// Get all locations in Berlin.
func ExampleRadarClient_search_location() {
	// https://radar.squat.net/api/1.2/search/location.json?facets[locality][]=Berlin
	radar := NewRadarClient()
	sb := radar.NewSearchBuilder()
	sb.Facets(Facet{location.FacetLocality, "Berlin"})
	sb.Fields(location.FieldAll)
	result, err := radar.SearchLocation(sb)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	if result == nil {
		fmt.Println("No results")
	}
	fmt.Printf("%v", result)
}
