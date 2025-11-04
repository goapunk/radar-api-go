package location

import "encoding/json"

type Location struct {
	// Location Fields
	Address    *Address `json:"address"`
	Directions string   `json:"directions"`
	Map        *Map     `json:"map"`
	Timezone   string   `json:"timezone"`
	Squat      string   `json:"squat"`
	Id         string   `json:"id"`
	Type       string   `json:"type"`
	Title      string   `json:"title"`
	UUID       string   `json:"uuid"`
	FeedNid    string   `json:"feed_nid"`
	// Reference fields (when a term is referenced in another entity),
	// see ResolveField().
	ReferenceUri      string `json:"uri,omitempty"`
	ReferenceResource string `json:"resource,omitempty"`
}

type Address struct {
	Country               string `json:"country"`
	NameLine              string `json:"name_line"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	OrganisationName      string `json:"organisation_name"`
	AdministrativeArea    string `json:"administrative_area"`
	SubAdministrativeArea string `json:"sub_administrative_area"`
	Locality              string `json:"locality"`
	DependantLocality     string `json:"dependent_locality"`
	PostalCode            string `json:"postal_code,"`
	Thoroughfare          string `json:"thoroughfare"`
	Premise               string `json:"premise"`
}

type Map struct {
	Geom           string `json:"geom"`
	GeoType        string `json:"geo_type"`
	Lat            string `json:"lat"`
	Lon            string `json:"lon"`
	Left           string `json:"left"`
	Top            string `json:"top"`
	Right          string `json:"right"`
	Bottom         string `json:"bottom"`
	SrId           string `json:"srid"`
	LatLon         string `json:"latlon"`
	SchemaOrgShape string `json:"schemaorg_shape"`
}

func (l *Location) UnmarshalJSON(data []byte) error {
	type tmp *Location
	if err := unmarshalJSON(tmp(l), data); err != nil {
		return err
	}
	// There is an inconsistency in the api: when querying events the `UUID` is
	// missing from the location and the `Id` field is holding the `UUID` value
	// instead.
	if len(l.UUID) == 0 {
		l.UUID = l.Id
	}
	return nil
}

func (a *Address) UnmarshalJSON(data []byte) error {
	type tmp *Address
	return unmarshalJSON(tmp(a), data)
}

func (m *Map) UnmarshalJSON(data []byte) error {
	type tmp *Map
	return unmarshalJSON(tmp(m), data)
}

func unmarshalJSON(e interface{}, data []byte) error {
	if len(data) < 3 {
		return nil
	}
	err := json.Unmarshal(data, e)
	return err
}
