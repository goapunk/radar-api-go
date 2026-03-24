package radarapi

import (
	"encoding/json"
	"fmt"
	"github.com/goapunk/radar-api-go/group"
	lang "golang.org/x/text/language"
	"strings"
)

// Get the Group associated with uuid. If no fields are specified the default are returned.
func (radar *RadarClient) Group(uuid string, language *lang.Tag, fields ...string) (*group.Group, error) {
	rawUrl := fmt.Sprintf("node/%s.json", uuid)
	raw, err := radar.prepareAndRunEntityQuery(rawUrl, language, fields)
	dec := json.NewDecoder(strings.NewReader(raw))
	dec.DisallowUnknownFields()
	g := &group.Group{}
	err = dec.Decode(g)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return g, nil
}
