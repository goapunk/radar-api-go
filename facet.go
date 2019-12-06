package radarapi

// Facets provide an easy way to filter results by entity properties.
//
// For each entity there's a predefined set of facets. The list of available facets can be found in the corresponding
// subpackages (see subdirectories).
//
// Facets can be passed to the server by using the facets parameter:
//  https://radar.squat.net/api/1.2/search/events.json?facets[facet][]=value
// Example
//  https://radar.squat.net/api/1.2/search/events.json?facets[city][]=Berlin
// Multiple facets can be passed, even the same with different values. When passing the
// same facet with different values the logical operation (AND, OR) depends on the actual
// facet. Each facet is annotated with its operator (see Constants in the corresponding subpackages)
//
// Example
//  category operator is AND, this call returns events which have both the bar/cafe AND theater category
//  https://radar.squat.net/api/1.2/search/events.json?facets[category][]=bar-cafe&facets[category][]=theater
//  group operator is OR, this call returns events organized by group 339006 OR 1599
//  https://radar.squat.net/api/1.2/search/events.json?facets[group][]=339006&fields=*&facets[group][]=1599
type Facet struct {
	// One of event/group/location/term FacetXXXs. Must match the entity you are searching for,
	// e.g. when searching for an event it must be from event.FacetXXXX
	Key   string
	Value string
}
