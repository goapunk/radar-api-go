package radarapi

import (
	e "github.com/goapunk/radar-api-go/event"
	l "github.com/goapunk/radar-api-go/location"
	"testing"
)

func TestResolveField(t *testing.T) {
	var expect = e.FieldCategory + ":" + e.FieldTitle + ":" + l.FieldAddress + ":" + l.FieldId
	resolved := ResolveField(e.FieldCategory, e.FieldTitle, l.FieldAddress, l.FieldId)
	if resolved != expect {
		t.Errorf("response didn't match")
	}
}
