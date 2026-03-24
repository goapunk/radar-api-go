package radarapi

import (
	"encoding/json"
	"fmt"
	"github.com/goapunk/radar-api-go/location"
	lang "golang.org/x/text/language"
	"strings"
)

// Get the Location associated with uuid. If no fields are specified the default are returned.
func (radar *RadarClient) Location(uuid string, language *lang.Tag, fields ...string) (*location.Location, error) {
	rawUrl := fmt.Sprintf("location/%s.json", uuid)
	raw, err := radar.prepareAndRunEntityQuery(rawUrl, language, fields)
	dec := json.NewDecoder(strings.NewReader(raw))
	dec.DisallowUnknownFields()
	l := &location.Location{}
	err = dec.Decode(l)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return l, nil
}
