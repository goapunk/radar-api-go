package radarapi

import (
	"testing"
)

// Test does a real GET, should be mocked or will fail eventually.
func TestTerm(t *testing.T) {
	radar := NewRadarClient()
	_, err := radar.Term("2a7f6975-4c01-4777-8611-dffe0306c06f", nil)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
