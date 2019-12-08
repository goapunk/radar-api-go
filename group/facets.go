package group

//noinspection GoUnusedConst

const (
	//operator: none
	FacetOfflineMap = "offline:map"
	// For some reason values need to start with a lowercase here, e.g. berlin
	// operator: and
	FacetCity = "city"
	// Uppercase country code, e.g. DE for Germany
	// operator: and
	FacetCountry = "country"
	// operator: and
	FacetSquated = "squated"
	// Group is listed by this group.
	// operator: and
	FacetRadarGroupListedBy = "radar_group_listed_by"
	// If the group is active
	// operator: or
	FacetActive = "field_active"
	// topic / tag
	// operator: and
	FacetTopic = "topic"
	// Spaces and slashes should be replaced by a dash, e.g. "work space/diy" =>
	// "work-space-diy"
	// operator: and
	FacetCategory = "category"
	// operator: and
	FacetTitle = "title"
)
