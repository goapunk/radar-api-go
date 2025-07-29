package radarapi

import (
	"github.com/goapunk/radar-api-go/location"
	"encoding/json"
	"fmt"
	lang "golang.org/x/text/language"
	"strings"
)

type SearchResultLocations struct {
	Results map[string]*location.Location `json:"result,omitempty"`
	// Number of results. Only the first 500 are actually returned.
	Count  int64                     `json:"count,string"`
	Facets map[string][]*ResultFacet `json:"facets"`
}

// Get the Location associated with uuid. If no fields are specified the default are returned.
func (radar *RadarClient) Location(uuid string, language *lang.Tag, fields ...string) (*location.Location, error) {
	rawUrl := fmt.Sprintf("%slocation/%s.json", baseUrl, uuid)
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
