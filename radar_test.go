package radarapi

import (
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestWithTimeout(t *testing.T) {
	var expect = 180
	radarClient := NewRadarClient(WithTimeout(180))
	if radarClient.GetTimeout() != expect {
		t.Errorf("expected timeout: %d, saw: %d", expect, radarClient.GetTimeout())
	}
}

func TestWithLoggerLevel(t *testing.T) {
	var expect = slog.LevelWarn
	var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	radarClient := NewRadarClient(WithLogger(logger))
	if !radarClient.GetLogger().Enabled(nil, expect) {
		t.Errorf("expected log level warn")
	}
	if radarClient.GetLogger().Enabled(nil, slog.LevelDebug) {
		t.Errorf("expected a max log level of warn, got debug")
	}
}

func TestWithClient(t *testing.T) {
	var (
		expect = 1312 * time.Second
		client = &http.Client{Timeout: 1312 * time.Second}
	)
	radarClient := NewRadarClient(WithClient(client))
	if radarClient.GetClient().Timeout != expect {
		t.Errorf("expected timeout of 1312 seconds")
	}
}
