package radarapi

import (
	"encoding/json"
	"fmt"
	"github.com/goapunk/radar-api-go/term"
	lang "golang.org/x/text/language"
	"strings"
)

type SearchResultTerms struct {
	Results map[string]*term.Term `json:"result"`
	// Number of results. Only the first 500 are actually returned.
	Count  int64                     `json:"count"`
	Facets map[string][]*ResultFacet `json:"facets"`
}

// Get the Term associated with uuid. If no fields are specified the default are returned.
func (radar *RadarClient) Term(uuid string, language *lang.Tag, fields ...string) (*term.Term, error) {
	rawUrl := fmt.Sprintf("taxonomy_term/%s.json", uuid)
	raw, err := radar.prepareAndRunEntityQuery(rawUrl, language, fields)
	dec := json.NewDecoder(strings.NewReader(raw))
	dec.DisallowUnknownFields()
	t := &term.Term{}
	err = dec.Decode(t)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return t, nil
}
