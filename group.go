package radarapi

import (
	"fmt"
	lang "golang.org/x/text/language"
)

// Get the group associated with uuid. If no fields are specified the default are returned.
//
// The returned string is the raw json response. See the examples on how to umarshal it.
func (radar *RadarClient) Group(uuid string, language *lang.Tag, fields ...string) (string, error) {
	rawUrl := fmt.Sprintf("%snode/%s.json", baseUrl, uuid)
	return radar.prepareAndRunEntityQuery(rawUrl, language, fields)
}
