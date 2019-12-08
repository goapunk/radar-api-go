package radarapi

import (
	e "0xacab.org/radarapi/event"
	l "0xacab.org/radarapi/location"
	"testing"
)

func TestResolveField(t *testing.T) {
	var expect = e.FieldCategory + ":" + e.FieldTitle + ":" + l.FieldAddress + ":" + l.FieldId
	resolved := ResolveField(e.FieldCategory, e.FieldTitle, l.FieldAddress, l.FieldId)
	if resolved != expect {
		t.Errorf("response didn't match")
	}
}
