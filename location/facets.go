package location

//noinspection GoUnusedConst

const (
	// Uppercase country code, e.g. DE for Germany
	// operator: and
	FacetCountry           = "country"
	// operator: and
	FacetAdminstrativeArea = "adminstrative_area"
	// Always starts with an uppercase letter, e.g. Berlin.
	// In event and group this facet is called city, seems to be an inconsistency.
	// operator: and
	FacetLocality          = "locality"
	// operator: and
	FacetDependentLocality = "depdendent_locality"
)
