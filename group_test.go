package radarapi

import (
	"testing"
)

// Test does a real GET, should be mocked or will fail eventually.
func TestGroup(t *testing.T) {
	radar := NewRadarClient()
	_, err := radar.Group("1662899c-ea08-431b-8238-ad775e9ecea6", nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
