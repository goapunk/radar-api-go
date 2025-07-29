package radarapi

import (
	"github.com/goapunk/radar-api-go/event"
	"encoding/json"
	"fmt"
	lang "golang.org/x/text/language"
	"log"
	"strings"
)

type SearchResultEvents struct {
	Results map[string]*event.Event `json:"result"`
	// Number of results. Only the first 500 are actually returned.
	Count  int64                     `json:"count"`
	Facets map[string][]*ResultFacet `json:"facets"`
}

// Get the Event associated with uuid. If no fields are specified the default are returned.
func (radar *RadarClient) Event(uuid string, language *lang.Tag, fields ...string) (*event.Event, error) {
	rawUrl := fmt.Sprintf("%snode/%s.json", baseUrl, uuid)
	raw, err := radar.prepareAndRunEntityQuery(rawUrl, language, fields)
	log.Print(raw)
	dec := json.NewDecoder(strings.NewReader(raw))
	//dec.DisallowUnknownFields()
	e := &event.Event{}
	err = dec.Decode(e)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return e, nil
}
