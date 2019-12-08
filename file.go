package radarapi

import (
	"0xacab.org/radarapi/file"
	"encoding/json"
	"fmt"
	lang "golang.org/x/text/language"
	"strings"
)

// Get the File associated with uuid. If no fields are specified the default are returned.
func (radar *RadarClient) File(uuid string, language *lang.Tag, fields ...string) (*file.File, error) {
	rawUrl := fmt.Sprintf("%sfile/%s.json", baseUrl, uuid)
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
