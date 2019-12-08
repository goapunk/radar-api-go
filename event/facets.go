package event

//noinspection GoUnusedConst

const (
	// operator: and
	FacetPostalCode = "field_offline:field_address:postal_code"
	// Always starts with an uppercase letter, e.g. Berlin
	// operator: and
	FacetCity = "city"
	// Uppercase country code, e.g. DE for Germany
	// operator: and
	FacetCountry = "country"
	// The date must be in the format YY YY-MM or YYYY-MM-DD
	// operator: or
	FacetDate = "date"
	// price category, e.g. "free-121"
	FacetPrice = "price"
	// Topic tag, e.g. "sport", "vegan"
	// operator: and
	FacetTag = "tag"
	// operator: or
	FacetGroup = "group"
	// Spaces and slashes should be replaced by a dash, e.g. "work space/diy" =>
	// "work-space-diy"
	// operator: and
	FacetCategory = "category"
)
