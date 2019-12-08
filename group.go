package radarapi

import (
	"0xacab.org/radarapi/group"
	"encoding/json"
	"fmt"
	lang "golang.org/x/text/language"
	"strings"
)

type SearchResultGroups struct {
	Results map[string]*group.Group `json:"result"`
	// Number of results. Only the first 500 are actually returned.
	Count  int64                    `json:"count"`
	Facets map[string][]*ResultFacet `json:"facets"`
}

// Get the Group associated with uuid. If no fields are specified the default are returned.
func (radar *RadarClient) Group(uuid string, language *lang.Tag, fields ...string) (*group.Group, error) {
	rawUrl := fmt.Sprintf("%snode/%s.json", baseUrl, uuid)
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
