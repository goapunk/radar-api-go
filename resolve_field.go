package radarapi

import (
	"strings"
)

// When searching for an entity, the result may contain a field which is
// a reference to another entity. Normally, one would have to do another API call to fetch this
// referenced entity. The radar api, however, allows us to specify fields which
// should be resolved automatically.
//
// In the radar rest api this is done like this
//  fields=referenceField:targetField
// or even
//  fields=referenceField:anotherReference:targetField
// The Go equivalent is using this method
//  ResolveField(referenceField, targetField)
// or
//  ResolveField(referenceField, anotherReference, targetField)
// The arguments can be any of the event/group/location/term FieldXXXs
//
// Example
//
// Search result for events showing the title and the field 'offline', which is
// a reference to the location entity.
//
// Using the radar rest api
//  https://radar.squat.net/api/1.2/search/events.json?fields=title,offline
// In Go
//  radar := NewRadarClient()
//  sb := NewSearch(EVENT)
//  sb.Fields(event.FieldTitle, event.FieldOffline)
//  raw, err := radar.Search(sb)
// Result
//  "346823": {
//      "offline": [
//        {
//          "uri": "https://radar.squat.net/api/1.2/location/8b8b63b2-64e5-4831-a91d-7f77bb510890",
//          "id": "8b8b63b2-64e5-4831-a91d-7f77bb510890",
//          "resource": "location",
//          "title": "K19 Café Kreutzigerstr. 19  Berlin Deutschland"
//        }
//      ],
//      "title": "CryptoParty"
//    }
// Clicking the uri will reveal the following structure of the location entity:
//  {
//   "address": {
//     "country": "DE",
//     "name_line": "K19 Café",
//     "first_name": "K19",
//     "last_name": "Café",
//     "organisation_name": null,
//     "administrative_area": null,
//     "sub_administrative_area": null,
//     "locality": "Berlin",
//     "dependent_locality": null,
//     "postal_code": "10247",
//     "thoroughfare": "Kreutzigerstr. 19",
//     "premise": ""
//   },
//   "directions": "U-Bhf. Samariterstrasse",
//   "map": {
//     "geom": "POINT (13.4600646 52.5131015)",
//     "geo_type": "point",
//     "lat": "52.513101500000",
//     "lon": "13.460064600000",
//     "left": "13.460064600000",
//     "top": "52.513101500000",
//     "right": "13.460064600000",
//     "bottom": "52.513101500000",
//     "srid": null,
//     "latlon": "52.513101500000,13.460064600000",
//     "schemaorg_shape": ""
//   },
//   "timezone": "Europe/Berlin",
//   "squat": null,
//   "id": "225",
//   "type": "location",
//   "title": "K19 Café Kreutzigerstr. 19  Berlin Deutschland",
//   "uuid": "8b8b63b2-64e5-4831-a91d-7f77bb510890",
//   "feed_nid": null
//  }
// We can now resolve the address Field (or any of the others) like this
//  https://radar.squat.net/api/1.2/search/events.json?fields=title,offline,offline:address
// In Go
//  radar := NewRadarClient()
//  sb := NewSearch(EVENT)
//  sb.Fields(event.FieldTitle, event.FieldOffline, ResolveField(event.Offline, location.Address))
//  raw, err := radar.Search(sb)
// Result
//  "346823": {
//      "offline": [
//        {
//          "uri": "https://radar.squat.net/api/1.2/location/8b8b63b2-64e5-4831-a91d-7f77bb510890",
//          "id": "8b8b63b2-64e5-4831-a91d-7f77bb510890",
//          "resource": "location",
//          "title": "K19 Café Kreutzigerstr. 19  Berlin Deutschland",
//          "address": {
//            "country": "DE",
//            "name_line": "K19 Café",
//            "first_name": "K19",
//            "last_name": "Café",
//            "organisation_name": null,
//            "administrative_area": null,
//            "sub_administrative_area": null,
//            "locality": "Berlin",
//            "dependent_locality": null,
//            "postal_code": "10247",
//            "thoroughfare": "Kreutzigerstr. 19",
//            "premise": ""
//          }
//        }
//      ],
//      "title": "CryptoParty"
//    }
func ResolveField(reference string, target string, targetN ...string) string {
	var builder strings.Builder
	builder.WriteString(reference)
	builder.WriteString(":")
	builder.WriteString(target)
	for _, f := range targetN {
		builder.WriteString(":")
		builder.WriteString(f)
	}
	return builder.String()
}
