package radarapi

import (
	"encoding/json"
	"fmt"
	"github.com/goapunk/radar-api-go/file"
	lang "golang.org/x/text/language"
	"strings"
)

// Get the File associated with uuid. If no fields are specified the default are returned.
func (radar *RadarClient) File(uuid string, language *lang.Tag, fields ...string) (*file.File, error) {
	rawUrl := fmt.Sprintf("file/%s.json", uuid)
	raw, err := radar.prepareAndRunEntityQuery(rawUrl, language, fields)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(strings.NewReader(raw))
	dec.DisallowUnknownFields()
	f := &file.File{}
	err = dec.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return f, nil
}
