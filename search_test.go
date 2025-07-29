package radarapi

import (
	"github.com/goapunk/radar-api-go/event"
	"github.com/goapunk/radar-api-go/group"
	"github.com/goapunk/radar-api-go/location"
	"testing"
)

// Test does a real GET, should be mocked or will fail eventually.
func TestSearchEventsCategoryDate(t *testing.T) {
	radar := NewRadarClient()
	facets := make([]Facet, 3)
	facets[0] = Facet{event.FacetCity, "Berlin"}
	facets[1] = Facet{event.FacetCategory, "work-space-diy"}
	facets[2] = Facet{event.FacetDate, "2020-11-24"}
	sb := &SearchBuilder{}
	sb.Facets(facets...)
	sb.Fields(event.FieldTitle, event.FieldOffline)
	raw, err := radar.SearchEvents(sb)
	t.Log(raw)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}

func TestRadarClient_SearchGroup(t *testing.T) {
	radar := NewRadarClient()
	sb := &SearchBuilder{}
	sb.Facets(Facet{group.FacetCity, "berlin"}, Facet{group.FacetCategory, "bar-cafe"})
	result, err := radar.SearchGroup(sb)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if result == nil {
		t.Log("No results")
	}
}

// Get all locations in Berlin.
func TestRadarClient_SearchLocation(t *testing.T) {
	// https://radar.squat.net/api/1.2/search/location.json?facets[locality][]=Berlin
	radar := NewRadarClient()
	sb := &SearchBuilder{}
	sb.Facets(Facet{location.FacetLocality, "Berlin"})
	result, err := radar.SearchLocation(sb)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if result == nil {
		t.Log("No results")
	}
}

// Test group logo reference.
func TestRadarClient_SearchGroupLogo(t *testing.T) {
	// https://radar.squat.net/api/1.2/search/location.json?facets[locality][]=Berlin
	radar := NewRadarClient()
	sb := &SearchBuilder{}
	sb.Facets(Facet{group.FacetCity, "berlin"})
	sb.Fields(group.FieldGroupLogo)
	result, err := radar.SearchGroup(sb)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if result == nil {
		t.Log("No results")
	}
	sama := result.Results["1614"]
	if sama.Logo == nil {
		t.Error("logo is not supposed to be nil")
	}
}
