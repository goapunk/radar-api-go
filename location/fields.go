package location

// Available fields for a location.
const (
	// All will return all available fields.
	FieldAll        = "*"
	FieldAddress    = "address"
	FieldDirections = "directions"
	FieldMap        = "map"
	FieldTimezone   = "timezone"
	FieldSquat      = "squat"
	FieldId         = "id"
	FieldType       = "type"
	FieldTitle      = "title"
	FieldUUID       = "uuid"
	FieldFeedNid    = "feed_nid"
)

// Available fields for an address.
const (
	FieldAddressCountry               = "country"
	FieldAddressNameLine              = "name_line"
	FieldAddressFirstName             = "first_name"
	FieldAddressLastName              = "last_name"
	FieldAddressOrganisationName      = "organisation_name"
	FieldAddressAdministrativeArea    = "administrative_area"
	FieldAddressSubAdministrativeArea = "sub_administrative_area"
	FieldAddressLocality              = "locality"
	FieldAddressDependentLocality     = "dependent_locality"
	FieldAddressPostalCode            = "postal_code"
	FieldAddressThoroughfare          = "thoroughfare"
	FieldAddressPremise               = "premise"
)

// Available fields for a map.
const (
	FieldMapGeom           = "geom"
	FieldMapGeoType        = "geo_type"
	FieldMapLat            = "lat"
	FieldMapLon            = "lon"
	FieldMapLeft           = "left"
	FieldMapTop            = "top"
	FieldMapRight          = "right"
	FieldMapBottom         = "bottom"
	FieldMapSrId           = "srid"
	FieldMapLatLon         = "latlon"
	FieldMapSchemaOrgShape = "schemaorg_shape"
)
