package radarapi

import (
	"0xacab.org/radarapi/event"
	"testing"
)

// Test does a real GET, should be mocked or will fail eventually.
func TestSearchEventsCategoryDate(t *testing.T) {
	// https://radar.squat.net/api/1.2/search/events.json?facets[city][]=Berlin&facets[category][]=work-space-diy&facets[date][]=2020-11-24&fields=title,offline
	const expect = `{"result":{"346823":{"offline":[{"uri":"https://radar.squat.net/api/1.2/location/8b8b63b2-64e5-4831-a91d-7f77bb510890","id":"8b8b63b2-64e5-4831-a91d-7f77bb510890","resource":"location","title":"K19 Caf\u00e9 Kreutzigerstr. 19  Berlin Deutschland"}],"title":"CryptoParty"}},"count":1,"facets":{"field_offline:field_address:postal_code":[{"filter":"10247","count":1,"formatted":"10247"}],"city":[{"filter":"Berlin","count":1,"formatted":"Berlin"}],"country":[{"filter":"DE","count":1,"formatted":"DE"}],"date":[{"filter":"1606240800","count":1,"formatted":"1606240800"}],"price":[{"filter":"free-121","count":1,"formatted":"free"}],"group":[{"filter":"1599","count":1,"formatted":"Stressfaktor"},{"filter":"7186","count":1,"formatted":"K19 Caf\u00e9"}],"category":[{"filter":"course-workshop","count":1,"formatted":"course/workshop"},{"filter":"work-space-diy","count":1,"formatted":"work space/diy"}]}}`
	radar := NewRadarClient()
	facets := make([]Facet, 3)
	facets[0] = Facet{event.FacetCity, "Berlin"}
	facets[1] = Facet{event.FacetCategory, "work-space-diy"}
	facets[2] = Facet{event.FacetDate, "2020-11-24"}
	sb := NewSearch(EVENT)
	sb.Facets(facets...)
	sb.Fields(event.FieldTitle, event.FieldOffline)
	raw, err := radar.Search(sb)
	t.Log(raw)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if raw != expect {
		t.Log(expect)
		t.Errorf("response didn't match")
	}
}
