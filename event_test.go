package radarapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test does a real GET, should be mocked
func TestEvent(t *testing.T) {
	const result = `{"body":{"value":"<p>Vegan Food at 7pm. WTF Queer is a DIY Space for the Community</p>\n","summary":"","format":"plain_text"},"category":[],"group_content_access":"0","og_group_ref":[{"uri":"https://radar.squat.net/api/1.2/node/8542609b-42d7-4d76-902b-5a919ff3bc95","id":"8542609b-42d7-4d76-902b-5a919ff3bc95","resource":"node"}],"og_group_request":[],"date_time":[{"value":"1756918800","value2":"1756940400","duration":21600,"time_start":"2025-09-03T19:00:00+02:00","time_end":"2025-09-04T01:00:00+02:00","rrule":null}],"image":[],"price":null,"email":null,"link":[{"url":"https://vrankrijk.org/events/wtf-queer-wednesday-123/","attributes":[],"display_url":null}],"offline":[{"uri":"https://radar.squat.net/api/1.2/location/b64778f7-d25c-472e-86ae-29413bc01b3a","id":"b64778f7-d25c-472e-86ae-29413bc01b3a","resource":"location"}],"phone":null,"topic":[{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/55574c13-14e8-43e3-87c9-d737018d3b93","id":"55574c13-14e8-43e3-87c9-d737018d3b93","resource":"taxonomy_term"},{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/95284ec4-ec5f-44cc-9d66-e1a89f07afe5","id":"95284ec4-ec5f-44cc-9d66-e1a89f07afe5","resource":"taxonomy_term"}],"title_field":"WTF Queer Wednesday","price_category":[],"event_status":"cancelled","flyer":[],"og_membership":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/637308","id":637308,"resource":"og_membership"}],"og_membership__1":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/637308","id":637308,"resource":"og_membership"}],"og_membership__2":[],"og_membership__3":[],"og_group_ref__og_membership":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/637308","id":637308,"resource":"og_membership"}],"og_group_ref__og_membership__1":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/637308","id":637308,"resource":"og_membership"}],"og_group_ref__og_membership__2":[],"og_group_ref__og_membership__3":[],"og_group_request__og_membership":[],"og_group_request__og_membership__1":[],"og_group_request__og_membership__2":[],"og_group_request__og_membership__3":[],"nid":"538905","vid":"792655","is_new":false,"type":"event","title":"WTF Queer Wednesday","language":"en","url":"https://radar.squat.net/en/event/amsterdam/vrankrijk/2025-09-03/wtf-queer-wednesday","edit_url":"https://radar.squat.net/en/node/538905/edit","status":"1","promote":"0","sticky":"0","created":"1753221605","changed":"1756653278","feed_nid":"493826","feed_node":{"uri":"https://radar.squat.net/api/1.2/node/493826","id":"493826","resource":"node"},"flag_abuse_node_user":[],"flag_abuse_whitelist_node_user":[],"uuid":"854d6e68-17d0-45cc-940e-624e571e0b93","vuuid":"d5aedfce-2f72-42e4-806b-5ca90b6c767d"}`
	//https://radar.squat.net/api/1.2/node/c866676b-c7c2-4eba-bb43-4713243eb09b.json
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/1.2/node/c866676b-c7c2-4eba-bb43-4713243eb09b.json" {
			t.Errorf("Expected to request '/api/1.2/node/c866676b-c7c2-4eba-bb43-4713243eb09b.json', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}))
	defer server.Close()

	radar := NewRadarClient()
	radar.SetBaseUrl(server.URL + "/api/1.2")

	value, err := radar.Event("c866676b-c7c2-4eba-bb43-4713243eb09b", nil)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	if value.UUID != "854d6e68-17d0-45cc-940e-624e571e0b93" {
		t.Errorf("Expected UUID 854d6e68-17d0-45cc-940e-624e571e0b93, got: %s", value.UUID)
	}
	if value.Title != "WTF Queer Wednesday" {
		t.Errorf("Expected title WTF Queer Wednesday, got: %s", value.Title)
	}

}
