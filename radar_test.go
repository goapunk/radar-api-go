package radarapi

import (
	"log/slog"
	"os"
	"testing"

	r "github.com/goapunk/radar-api-go"
)

func TestWithTimeout(t *testing.T) {
	var expect = 180
	radarClient := r.NewRadarClient(r.WithTimeout(180))
	if radarClient.GetTimeout() != expect {
		t.Errorf("expected timeout: %d, saw: %d", expect, radarClient.GetTimeout())
	}
}

func TestWithLogger(t *testing.T) {
	var expect = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	radarClient := r.NewRadarClient(r.WithLogger(expect))
	if radarClient.GetLogger() != expect {
		t.Errorf("expected json logger")
	}
}
