package radarapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test does a real GET, should be mocked
func TestEvent(t *testing.T) {
	const result = `{"body":{"value":"removed","summary":"<h3><strong>Disgra\u00e7a \u2013 a story about an anarchist social centre</strong></h3>\n","format":"rich_text_editor"},"category":[{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/e97f372b-29bc-460b-bff6-35d2462411ff","id":"e97f372b-29bc-460b-bff6-35d2462411ff","resource":"taxonomy_term"}],"group_content_access":"0","og_group_ref":[{"uri":"https://radar.squat.net/api/1.2/node/eecf916f-6218-458b-b18c-315348a956e5","id":"eecf916f-6218-458b-b18c-315348a956e5","resource":"node"}],"og_group_request":[],"date_time":[{"value":"1729983900","value2":"1759273500","duration":29289600,"time_start":"2024-10-27T00:05:00+01:00","time_end":"2025-10-01T00:05:00+01:00","rrule":null}],"image":{"file":{"uri":"https://radar.squat.net/api/1.2/file/1d978950-9016-4451-84ba-4f7d5817abe2","id":"1d978950-9016-4451-84ba-4f7d5817abe2","resource":"file"}},"price":null,"email":"disgraca@riseup.net ","link":[{"url":"https://www.gofundme.com/f/disgraca","attributes":[],"display_url":null}],"offline":[{"uri":"https://radar.squat.net/api/1.2/location/0db82018-d5e8-419d-8fe4-ec6f945415b3","id":"0db82018-d5e8-419d-8fe4-ec6f945415b3","resource":"location"}],"phone":null,"topic":[],"title_field":"Disgra\u00e7a: help us buy our anarchist social centre in Lisbon ","price_category":[{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/9d943d0c-e2bf-408e-9110-4bfb044f60c0","id":"9d943d0c-e2bf-408e-9110-4bfb044f60c0","resource":"taxonomy_term"}],"event_status":"confirmed","flyer":[],"callout":{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/1aba342d-b730-4afd-823b-ece4822db5fe","id":"1aba342d-b730-4afd-823b-ece4822db5fe","resource":"taxonomy_term"},"og_membership":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/585603","id":585603,"resource":"og_membership"}],"og_membership__1":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/585603","id":585603,"resource":"og_membership"}],"og_membership__2":[],"og_membership__3":[],"og_group_ref__og_membership":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/585603","id":585603,"resource":"og_membership"}],"og_group_ref__og_membership__1":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/585603","id":585603,"resource":"og_membership"}],"og_group_ref__og_membership__2":[],"og_group_ref__og_membership__3":[],"og_group_request__og_membership":[],"og_group_request__og_membership__1":[],"og_group_request__og_membership__2":[],"og_group_request__og_membership__3":[],"nid":"499022","vid":"718803","is_new":false,"type":"event","title":"Disgra\u00e7a: help us buy our anarchist social centre in Lisbon ","language":"en","url":"https://radar.squat.net/en/event/lisboa/disgraca/2024-10-27/disgraca-help-us-buy-our-anarchist-social-centre-lisbon","edit_url":"https://radar.squat.net/en/node/499022/edit","status":"1","promote":"0","sticky":"0","created":"1730045186","changed":"1730131722","feed_nid":null,"flag_abuse_node_user":[],"flag_abuse_whitelist_node_user":[],"uuid":"29dc524b-76f6-4094-bcbf-eda0db3d1fa8","vuuid":"474afe1d-009c-430a-9d89-081a02ba6030"}`
	//https://radar.squat.net/api/1.2/node/c866676b-c7c2-4eba-bb43-4713243eb09b.json
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/1.2/node/29dc524b-76f6-4094-bcbf-eda0db3d1fa8.json" {
			t.Errorf("Expected to request '/api/1.2/node/29dc524b-76f6-4094-bcbf-eda0db3d1fa8.json', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}))
	defer server.Close()
	radar := NewRadarClient()
	radar.SetBaseUrl(server.URL + "/api/1.2")
	value, err := radar.Event("29dc524b-76f6-4094-bcbf-eda0db3d1fa8", nil)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	if value.UUID != "29dc524b-76f6-4094-bcbf-eda0db3d1fa8" {
		t.Errorf("Expected UUID 29dc524b-76f6-4094-bcbf-eda0db3d1fa8, got: %s", value.UUID)
	}
	if value.Title != "Disgraça: help us buy our anarchist social centre in Lisbon " {
		t.Errorf("Expected title 'Disgraça: help us buy our anarchist social centre in Lisbon ', got: %s", value.Title)
	}
	if value.Image.File.ReferenceId != "1d978950-9016-4451-84ba-4f7d5817abe2" {
		t.Errorf("Expected Image File Id '1d978950-9016-4451-84ba-4f7d5817abe2', got: %s", value.Image.File.ReferenceId)
	}

}
