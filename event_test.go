package radarapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func TestEventNoImage(t *testing.T) {
	const result = `{"category":[{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/8463bb01-e974-4785-9c2d-b95d87c9ee2d","id":"8463bb01-e974-4785-9c2d-b95d87c9ee2d","resource":"taxonomy_term"}],"body":{"value":"<p>\u201cWat betekent het om \u2018gek\u2019 te zijn in een gestoorde wereld? Hoe raken psychiatrie, racisme en seksisme elkaar? Wat als we stigma niet langer zien als een individueel probleem, maar als een vorm van maatschappelijke uitsluiting? In de Mad Studies-leesgroep gaan we samen aan de slag met uitdagende teksten over psychiatrie en gekte. We onderzoeken nieuwe perspectieven die je helpen om je eigen ervaringen en idee\u00ebn te verwoorden rondom onderwerpen als diagnoses, medicatie, stigma, ervaringskennis, hulpverlening en Mad Pride.\u00a0</p>\n<p>Elke sessie lezen we een of twee teksten in het Nederlands of Engels. We starten met een korte inleiding of video, waarna we samen in gesprek gaan over de tekst en onze eigen ervaringen. Tijdens vier bijeenkomsten ben jij als deelnemer medeverantwoordelijk voor de themakeuze en invulling. We vormen een doorlopende groep waarin we samen leren en verdiepen. Door regelmatig deel te nemen groeien we als groep en verdiepen we onze inzichten.</p>\n<p>De leesgroep werkt volgens het principe \u2018pay what you can/pay what you want\u2019. Financi\u00ebn mogen geen belemmering zijn, dus je betaalt wat haalbaar is voor jou. Deze activiteit van stichting Perceval wordt dit jaar voor de tiende keer aangeboden.\u201d</p>\n<p>\u00a0</p>\n<p>\u00a0</p>\n<p>\u00a0</p>\n<p>\u00a0</p>\n<p>\u00a0</p>\n","summary":"","format":"rich_text_editor"},"date_time":[{"value":"1762354800","value2":"1762363800","duration":9000,"time_start":"2025-11-05T16:00:00+01:00","time_end":"2025-11-05T18:30:00+01:00","rrule":"RRULE:FREQ=WEEKLY;INTERVAL=3;BYDAY=WE;COUNT=8;WKST=MO"}],"image":[],"price":null,"group_content_access":"0","og_group_ref":[{"uri":"https://radar.squat.net/api/1.2/node/c825a595-ffc2-48f0-adb4-f2a84d3cc37f","id":"c825a595-ffc2-48f0-adb4-f2a84d3cc37f","resource":"node"}],"og_group_request":[],"email":null,"link":[],"offline":[{"uri":"https://radar.squat.net/api/1.2/location/792644d9-c3d1-407d-bd2c-bea33c391420","id":"792644d9-c3d1-407d-bd2c-bea33c391420","resource":"location"}],"phone":null,"topic":[],"title_field":"Mad Studies","price_category":[],"event_status":"confirmed","flyer":[],"og_membership":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/643480","id":643480,"resource":"og_membership"}],"og_membership__1":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/643480","id":643480,"resource":"og_membership"}],"og_membership__2":[],"og_membership__3":[],"og_group_ref__og_membership":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/643480","id":643480,"resource":"og_membership"}],"og_group_ref__og_membership__1":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/643480","id":643480,"resource":"og_membership"}],"og_group_ref__og_membership__2":[],"og_group_ref__og_membership__3":[],"og_group_request__og_membership":[],"og_group_request__og_membership__1":[],"og_group_request__og_membership__2":[],"og_group_request__og_membership__3":[],"nid":"543232","vid":"802121","is_new":false,"type":"event","title":"Mad Studies","language":"en","url":"https://radar.squat.net/en/event/rotterdam/snackbar-frieda/2025-11-05/mad-studies","edit_url":"https://radar.squat.net/en/node/543232/edit","status":"1","promote":"0","sticky":"0","created":"1756739636","changed":"1756739636","feed_nid":null,"flag_abuse_node_user":[],"flag_abuse_whitelist_node_user":[],"uuid":"9041dc19-6954-4d23-bc57-1c57da33da3b","vuuid":"3f0d9497-9567-4a94-8f3e-f4faa8d852e1"}`
	//https://radar.squat.net/api/1.2/node/c866676b-c7c2-4eba-bb43-4713243eb09b.json
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/1.2/node/9041dc19-6954-4d23-bc57-1c57da33da3b.json" {
			t.Errorf("Expected to request '/api/1.2/node/9041dc19-6954-4d23-bc57-1c57da33da3b.json', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}))
	defer server.Close()
	radar := NewRadarClient()
	radar.SetBaseUrl(server.URL + "/api/1.2")
	value, err := radar.Event("9041dc19-6954-4d23-bc57-1c57da33da3b", nil)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	if value.UUID != "9041dc19-6954-4d23-bc57-1c57da33da3b" {
		t.Errorf("Expected UUID 9041dc19-6954-4d23-bc57-1c57da33da3b, got: %s", value.UUID)
	}
	if value.Title != "Mad Studies" {
		t.Errorf("Expected title 'Mad Studies', got: %s", value.Title)
	}
	if value.Image.File != nil {
		t.Errorf("Expected Image.File to be nil', got: %v", value.Image)
	}
}
