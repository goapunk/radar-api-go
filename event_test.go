package radarapi

import (
	"testing"
)

// Test does a real GET, should be mocked
func TestEvent(t *testing.T) {
	//https://radar.squat.net/api/1.2/node/49c0b0e4-7b1a-408c-85f8-1d6f306f277f.json
	radar := NewRadarClient()
	_, err := radar.Event("49c0b0e4-7b1a-408c-85f8-1d6f306f277f", nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
