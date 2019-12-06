package radarapi

import (
	"testing"
)

// Test does a real GET, should be mocked
func TestEvent(t *testing.T) {
	//https://radar.squat.net/api/1.2/node/49c0b0e4-7b1a-408c-85f8-1d6f306f277f.json
	const testEvent = `{"body":{"value":"<p>\"Ella Chord &amp; Gary Flanell\" (Charming Covers aus Berlin) und \"Wayne Lostsoul\" (Folk, Punk Singer aus Berlin)</p>\n","summary":"","format":"rich_text_editor"},"category":[{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/ff3e2872-6140-4645-98f0-784d656a9c5c","id":"ff3e2872-6140-4645-98f0-784d656a9c5c","resource":"taxonomy_term"}],"group_content_access":"0","og_group_ref":[{"uri":"https://radar.squat.net/api/1.2/node/1662899c-ea08-431b-8238-ad775e9ecea6","id":"1662899c-ea08-431b-8238-ad775e9ecea6","resource":"node"},{"uri":"https://radar.squat.net/api/1.2/node/9e43dac6-e1da-4f60-8428-de9f32ac9eb0","id":"9e43dac6-e1da-4f60-8428-de9f32ac9eb0","resource":"node"}],"og_group_request":[],"date_time":[{"value":"1415898000","value2":"1415898000","duration":0,"time_start":"2014-11-13T18:00:00+01:00","time_end":"2014-11-13T18:00:00+01:00","rrule":null}],"image":[],"price":null,"email":null,"link":[],"offline":[{"uri":"https://radar.squat.net/api/1.2/location/a1a36ff1-eb8f-4c87-be13-6af347af039b","id":"a1a36ff1-eb8f-4c87-be13-6af347af039b","resource":"location"}],"phone":null,"topic":[],"title_field":"Konzert","price_category":[],"event_status":null,"flyer":[],"og_membership":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/2222","id":2222,"resource":"og_membership"}],"og_membership__1":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/2222","id":2222,"resource":"og_membership"}],"og_membership__2":[],"og_membership__3":[],"og_group_ref__og_membership":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/2222","id":2222,"resource":"og_membership"}],"og_group_ref__og_membership__1":[{"uri":"https://radar.squat.net/api/1.2/entity_og_membership/2222","id":2222,"resource":"og_membership"}],"og_group_ref__og_membership__2":[],"og_group_ref__og_membership__3":[],"og_group_request__og_membership":[],"og_group_request__og_membership__1":[],"og_group_request__og_membership__2":[],"og_group_request__og_membership__3":[],"nid":"1605","vid":"1640","is_new":false,"type":"event","title":"Konzert","language":"de","url":"https://radar.squat.net/en/node/1605","edit_url":"https://radar.squat.net/en/node/1605/edit","status":"1","promote":"0","sticky":"0","created":"1415723921","changed":"1415725799","feed_nid":null,"flag_abuse_node_user":[],"flag_abuse_whitelist_node_user":[],"uuid":"49c0b0e4-7b1a-408c-85f8-1d6f306f277f","vuuid":"68ccf813-1b3c-4b31-803d-7790d747ac58"}`
	radar := NewRadarClient()
	raw, err := radar.Event("49c0b0e4-7b1a-408c-85f8-1d6f306f277f", nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if raw != testEvent {
		t.Errorf("response didn't match")
	}
}
