package radarapi

import (
	"github.com/goapunk/radar-api-go/event"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchEventsCategoryDate(t *testing.T) {
	const result = `{"result":{"519308":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/1e1543ec-18eb-411f-9bc7-357e82dddc92","id":"1e1543ec-18eb-411f-9bc7-357e82dddc92","resource":"location","title":"Regenbogencafe Lausitzer Str. 22a  Berlin Deutschland"}],"title":"Mittwochscaf\u00e9"},"496473":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/fbb94a74-7bc0-4392-b708-b9a3f642fd66","id":"fbb94a74-7bc0-4392-b708-b9a3f642fd66","resource":"location","title":"JUP Florastr. 84  Berlin Deutschland"}],"title":"K\u00fcche f\u00fcr Alle (jeden Mittwoch)"},"538487":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/78d7792b-84f1-41ed-b2b5-fd3986324b2c","id":"78d7792b-84f1-41ed-b2b5-fd3986324b2c","resource":"location","title":"JUP Florastra\u00dfe 84  Berlin Germany"}],"title":"Kiezk\u00fcche f\u00fcr Alle - Kiezteam Pankow"},"542917":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/2ea82d85-d598-48b7-84f0-962de758e860","id":"2ea82d85-d598-48b7-84f0-962de758e860","resource":"location","title":"Wagenburg Lohm\u00fchle Lohm\u00fchlenstr. 17  Berlin Deutschland"}],"title":"Soli - Kinoabend: Not Just Your Picture"},"533537":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/a232175b-8c59-44be-82fa-1ad600100dd5","id":"a232175b-8c59-44be-82fa-1ad600100dd5","resource":"location","title":"Kadterschmiede Rigaer Str. 94  Berlin Deutschland"}],"title":"SOBER VOxK\u00dc"},"542374":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/a232175b-8c59-44be-82fa-1ad600100dd5","id":"a232175b-8c59-44be-82fa-1ad600100dd5","resource":"location","title":"Kadterschmiede Rigaer Str. 94  Berlin Deutschland"}],"title":"Infoevent on the Situation of Rigaer 94 + K\u00fcfa"},"501974":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/5696c22e-aba7-493c-950d-51d175449055","id":"5696c22e-aba7-493c-950d-51d175449055","resource":"location","title":"KuBiZ Bernkasteler Str. 78  Berlin Deutschland"}],"title":"K\u00fcfa (vegan)"},"543330":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/5696c22e-aba7-493c-950d-51d175449055","id":"5696c22e-aba7-493c-950d-51d175449055","resource":"location","title":"KuBiZ Bernkasteler Str. 78  Berlin Deutschland"}],"title":"VideoKino - Open Air"}},"count":8,"facets":{}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/1.2/search/events.json" {
			t.Errorf("Expected to request '/api/1.2/search/events.json?facets[category][]=food&facets[city][]=Berlin&facets[date][]=2025-09-03&fields=title,offline', got: %s", r.URL.Path)
		}
		if r.URL.RawQuery != "facets%5Bcategory%5D%5B%5D=food&facets%5Bcity%5D%5B%5D=Berlin&facets%5Bdate%5D%5B%5D=2025-09-03&fields=title%2Coffline" {
			t.Errorf("Expected to request raw query 'facets%%5Bcategory%%5D%%5B%%5D=food&facets%%5Bcity%%5D%%5B%%5D=Berlin&facets%%5Bdate%%5D%%5B%%5D=2025-09-03&fields=title%%2Coffline', got: %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}))
	defer server.Close()
	radar := NewRadarClient()
	radar.SetBaseUrl(server.URL + "/api/1.2")
	facets := make([]Facet, 3)
	facets[0] = Facet{event.FacetCity, "Berlin"}
	facets[1] = Facet{event.FacetCategory, "food"}
	facets[2] = Facet{event.FacetDate, "2025-09-03"}
	sb := radar.NewSearchBuilder()
	sb.Facets(facets...)
	sb.Fields(event.FieldTitle, event.FieldOffline)
	results, err := radar.SearchEvents(sb)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	if results.Count != 8 {
		t.Errorf("Expected count to be 8, got: %d", results.Count)
	}
}

func TestSearchEventsNoResult(t *testing.T) {
	const result = `{"result":false,"count":0,"facets":null}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/1.2/search/events.json" {
			t.Errorf("Expected to request '/api/1.2/search/events.json?facets[category][]=nonexisting', got: %s", r.URL.Path)
		}
		if r.URL.RawQuery != "facets%5Bcategory%5D%5B%5D=nonexisting" {
			t.Errorf("Expected to request raw query 'facets%%5Bcategory%%5D%%5B%%5D=nonexisting', got: %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}))
	defer server.Close()
	radar := NewRadarClient()
	radar.SetBaseUrl(server.URL + "/api/1.2")
	facets := make([]Facet, 1)
	facets[0] = Facet{event.FacetCategory, "nonexisting"}
	sb := radar.NewSearchBuilder()
	sb.Facets(facets...)
	results, err := radar.SearchEvents(sb)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	if results != nil {
		t.Errorf("Expected no results, got: %d", results.Count)
	}
}
