package radarapi

import (
	"testing"
)

// If KØPI dies this api is dead
// Test does a real GET, should be mocked or will fail eventually.
func TestLocation(t *testing.T) {
	radar := NewRadarClient()
	_, err := radar.Location("f3632e1a-a091-4a3f-9a9e-9c979bf7aea7", nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
