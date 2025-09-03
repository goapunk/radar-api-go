// Package radarapi provides bindings for the radar.squat.net event API.
//
// A very brief documention of the API can be found at https://radar.squat.net/en/api
//
// This document will give a more detailed description and provide examples for both the web
// api and the go bindings found in this package.
//
// # General
//
// The radar database consists of the following entities:
//
//   - events : any kind of event, e.g. a party, workshop, etc.
//   - groups : a group organizes events. Some groups just act as promoters and add events to the
//     database without actually organzing them, e.g. "Stressfaktor". An event can be linked to multiple groups.
//   - location : a place where an event happens.
//   - taxonomy term : a category for events, groups and locations, e.g. "party",
//     "workshop", "bar/cafe", etc. Any of the above entities can be linked to
//     multiple taxonomy terms.
//
// Each entity comes with a set of fields. The list of available fields for each entity can be found in the corresponding
// subpackages (see subdirectories). All entities have a common field UUID which
// uniquely identifies an entity instance.
//
// # Endpoints
//
// The rest API provides endpoints for each of those entities:
//
//	event: https://radar.squat.net/api/1.2/node/[uuid].json
//	file:  https://radar.squat.net/api/1.2/file/[uuid].json
//	group: https://radar.squat.net/api/1.2/node/[uuid].json
//	location: https://radar.squat.net/api/1.2/location/[uuid].json
//	term: https://radar.squat.net/api/1.2/taxonomy_term/[uuid].json
//
// # Entity parameters
//
// The entity endpoints support two parameters:
//   - fields: the fields to be returned.
//   - language: use the specified language, if available.
//
// By default, all available fields for an entity are returned. Alternatively, a comma-separated list of fields can be passed:
//
//	https://radar.squat.net/api/1.2/node/[uuid].json?fields=field1,field2,field3
//
// The default language is english (or UND???). To change the language pass an ISO language code like this:
//
//	https://radar.squat.net/api/1.2/node/[uuid].json?language=de
//
// The corresponding go bindings for these endpoints provided by this package
// are named Event(), Group(), Location() and Term() and can be found below.
//
// # Search Endpoints
//
// The API supports searching the database for entities through the following endpoints:
//
//	event: https://radar.squat.net/api/1.2/search/events.json
//	group: https://radar.squat.net/api/1.2/search/groups.json
//	location: https://radar.squat.net/api/1.2/search/location.json
//	term: https://radar.squat.net/api/1.2/search/term.json
//
// Without any parameters, events will return all events with an end time greater than (now - 2h).
// group, location and term will simply return all.
//
// In most cases you'd only want a subset of the data, therefore the API provides different means of filtering the results.
// See Facet and Filter below for info and examples on how to filter the results.
//
// Apart from facets and filters, the search API supports the following other parameters:
//   - keys: string used for full-text search.
//   - fields: see Entity Parameters above, with the extension that reference fields can be auto resolved (see ResolveField below).
//   - language: see Entity Parameters above.
//   - limit: return at most n results, see Limit() below.
//   - sort: sorts the results by the specified field in the given order (asc or desc), see Sort() below.
//
// Searching with this library is done via the SearchBuilder type (see below for a description and examples).
package radarapi
