package radarapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLocation(t *testing.T) {
	const result = `{"address":{"country":"DE","name_line":"K\u00d8PI","first_name":"K\u00d8PI","last_name":"","organisation_name":null,"administrative_area":null,"sub_administrative_area":null,"locality":"Berlin","dependent_locality":null,"postal_code":"10179","thoroughfare":"K\u00f6penicker Stra\u00dfe 137","premise":""},"directions":null,"map":{"geom":"POINT (13.425979525618 52.507708524737)","geo_type":"point","lat":"52.507708524737","lon":"13.425979525618","left":"13.425979525618","top":"52.507708524737","right":"13.425979525618","bottom":"52.507708524737","srid":null,"latlon":"52.507708524737,13.425979525618","schemaorg_shape":""},"timezone":"Europe/Berlin","squat":"squat","id":"25","type":"location","title":"K\u00d8PI K\u00f6penicker Stra\u00dfe 137  Berlin Germany","uuid":"f3632e1a-a091-4a3f-9a9e-9c979bf7aea7","feed_nid":null}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/1.2/location/f3632e1a-a091-4a3f-9a9e-9c979bf7aea7.json" {
			t.Errorf("Expected to request '/api/1.2/location/f3632e1a-a091-4a3f-9a9e-9c979bf7aea7.json', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}))
	defer server.Close()

	radar := NewRadarClient()
	radar.SetBaseUrl(server.URL + "/api/1.2")

	value, err := radar.Location("f3632e1a-a091-4a3f-9a9e-9c979bf7aea7", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if value.UUID != "f3632e1a-a091-4a3f-9a9e-9c979bf7aea7" {
		t.Errorf("Expected UUID f3632e1a-a091-4a3f-9a9e-9c979bf7aea7, got: %s", value.UUID)
	}
	if value.Title != "KØPI Köpenicker Straße 137  Berlin Germany" {
		t.Errorf("Expected title 'KØPI Köpenicker Straße 137  Berlin Germany', got: %s", value.Title)
	}

}
