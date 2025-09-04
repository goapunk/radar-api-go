package radarapi

import (
	"log/slog"
	"os"
	"testing"
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
