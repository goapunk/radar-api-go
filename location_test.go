package radarapi

import (
	"testing"
)

// If KØPI dies this api is dead
// Test does a real GET, should be mocked or will fail eventually.
func TestLocation(t *testing.T) {
	//https://radar.squat.net/api/1.2/location/f3632e1a-a091-4a3f-9a9e-9c979bf7aea7.json
	const testLocation = `{"address":{"country":"DE","name_line":"K\u00d8PI","first_name":"K\u00d8PI","last_name":"","organisation_name":null,"administrative_area":null,"sub_administrative_area":null,"locality":"Berlin","dependent_locality":null,"postal_code":"10179","thoroughfare":"K\u00f6penicker Stra\u00dfe 137","premise":""},"directions":null,"map":{"geom":"POINT (13.425979525618 52.507708524737)","geo_type":"point","lat":"52.507708524737","lon":"13.425979525618","left":"13.425979525618","top":"52.507708524737","right":"13.425979525618","bottom":"52.507708524737","srid":null,"latlon":"52.507708524737,13.425979525618","schemaorg_shape":""},"timezone":"Europe/Berlin","squat":"squat","id":"25","type":"location","title":"K\u00d8PI K\u00f6penicker Stra\u00dfe 137  Berlin Germany","uuid":"f3632e1a-a091-4a3f-9a9e-9c979bf7aea7","feed_nid":null}`
	radar := NewRadarClient()
	raw, err := radar.Location("f3632e1a-a091-4a3f-9a9e-9c979bf7aea7", nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if raw != testLocation {
		t.Errorf("response didn't match")
	}
}
