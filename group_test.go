package radarapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGroup(t *testing.T) {
	const result = `{"body":{"value":"<p>enchanting international DIY live acts from 7 pm till 10 pm, afterwards free entry &amp; exquisite DJs, partying, hanging out, open till whenever! every 1st sunday cult kakaoke w/ Kj der K\u00e4pt'n</p>\n","summary":"","format":"rich_text_editor"},"category":[{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/2a56c4d7-eb98-4f96-9ac6-d383a1af5ce8","id":"2a56c4d7-eb98-4f96-9ac6-d383a1af5ce8","resource":"taxonomy_term"},{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/ff3e2872-6140-4645-98f0-784d656a9c5c","id":"ff3e2872-6140-4645-98f0-784d656a9c5c","resource":"taxonomy_term"},{"uri":"https://radar.squat.net/api/1.2/taxonomy_term/b89e871f-c923-4f66-a7ba-a34bc9ccce5b","id":"b89e871f-c923-4f66-a7ba-a34bc9ccce5b","resource":"taxonomy_term"}],"group_group":true,"group_logo":{"file":{"uri":"https://radar.squat.net/api/1.2/file/d0ae403f-ce43-4854-91d9-ecf86b92af53","id":"d0ae403f-ce43-4854-91d9-ecf86b92af53","resource":"file"}},"image":[],"email":null,"link":[{"url":"http://www.schokoladen-mitte.de","attributes":[],"display_url":null}],"offline":[{"uri":"https://radar.squat.net/api/1.2/location/a1a36ff1-eb8f-4c87-be13-6af347af039b","id":"a1a36ff1-eb8f-4c87-be13-6af347af039b","resource":"location"}],"opening_times":{"value":"<p>Mo &amp; Mi. - Sa. 19h krass gute internationale Konzerte!<br />ausser... Di. 19h LSD - Liebe Stadt Drogen (Lesung)<br />Mo. 22h Strange Tunes on Monday - Erlesene Djs<br />Der 1. Sonntag: Karaoke-Nacht mit KJ Der K\u00e4pt'n<br />Der 3. Sonntag: Schokoladen Open Stage - Play your Songs unplugged!</p>\n","format":"rich_text_editor"},"phone":null,"topic":[],"notifications":[],"active":false,"flyer":[],"members":[{"uri":"https://radar.squat.net/api/1.2/user/162","id":"162","resource":"user"}],"members__1":[{"uri":"https://radar.squat.net/api/1.2/user/162","id":"162","resource":"user"}],"members__2":[],"members__3":[],"radar_group_listed_by":[{"uri":"https://radar.squat.net/api/1.2/node/1599","id":1599,"resource":"node"}],"nid":"1604","vid":"703532","is_new":false,"type":"group","title":"Schokoladen","language":"de","url":"https://radar.squat.net/en/node/1604","edit_url":"https://radar.squat.net/en/node/1604/edit","status":"1","promote":"0","sticky":"0","created":"1415723268","changed":"1724713688","feed_nid":null,"flag_abuse_node_user":[],"flag_abuse_whitelist_node_user":[],"uuid":"1662899c-ea08-431b-8238-ad775e9ecea6","vuuid":"c80d1c43-790b-45d0-b2f6-4d657beb925c"}`
	//https://radar.squat.net/api/1.2/node/1662899c-ea08-431b-8238-ad775e9ecea6.json
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/1.2/node/1662899c-ea08-431b-8238-ad775e9ecea6.json" {
			t.Errorf("Expected to request '/api/1.2/node/1662899c-ea08-431b-8238-ad775e9ecea6.json', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}))
	defer server.Close()

	radar := NewRadarClient()
	radar.SetBaseUrl(server.URL + "/api/1.2")

	group, err := radar.Group("1662899c-ea08-431b-8238-ad775e9ecea6", nil)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	if group.Title != "Schokoladen" {
		t.Errorf("Expected title to be 'Schokoladen', got: %s", group.Title)
	}
}
