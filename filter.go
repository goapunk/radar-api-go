package radarapi

// Filters provide a somewhat more complex way to filter results.
//
// Note: filtering only works on indexed fields.
//
// Filters can be passed to the server by using the filter parameter:
//  https://radar.squat.net/api/1.2/search/events.json?filter[~operator][field][~comparator]=value
// Possible operators are
//  * ~and: all filters with ~and must evaluate to true
//  * ~or: at least one of the ~or filters must evaluate to true
// Comparators:
//  * ~eq: equals
//  * ~ne: not equal
//  * ~gt: greater than
//  * ~gte: greater or equal than
//  * ~lt: less than
//  * ~lte: less or equal than
// Note: filtering also works on resolved fields (see ResolveField).
//
// Example
//  Get all events in Berlin
//  https://radar.squat.net/api/1.2/search/events.json?filter[~and][field_offline:field_address:locality][~eq]=Berlin
// Note: most of the equal filters, like in this example, can be be done with Facets in a simpler manner (see Facets).
//
// One of the more interesting things you can do is filtering for a time period
//  Get all events in Berlin with a start date between 2020-11-03 00:00:00 and 2020-11-25 00:00:00
//  https://radar.squat.net/api/1.2/search/events.json?facets[city][]=Berlin&filter[~and][search_api_aggregation_1][~gt]=1604358000&filter[~or][search_api_aggregation_1][~lt]=1606258800
// In Go filters can be created with CreateFilter() and CreateRangeFilter().
//
// There are some predefined constants which can be used as fields when filtering (FilterXXX), see constants above.
//
// Limitations
//
// Filters can't be grouped, i.e. things like
//  (cond1 AND cond2) OR cond3
// are not possible.
//
// A field can't be used in more than one equals filter, i.e. you can't filter for
// events in Berlin OR Amsterdam because that would require using the field locality twice
//  filter[~and][field_offline:field_address:locality][~eq]=Berlin&filter[~or][field_offline:field_address:locality][~eq]=Amsterdam
type Filter interface{}

type filter struct {
	operator   string
	comparator string
	field      string
	value      string
}

type rangeFilter struct {
	field string
	from  string
	to    string
}

//noinspection GoUnusedConst

const (
	OperatorAnd   = "~and"
	OperatorOR    = "~or"
	ComparatorEQ  = "~e"
	ComparatorNEQ = "~ne"
	ComparatorGT  = "~gt"
	ComparatorGTE = "~gte"
	ComparatorLT  = "~lt"
	ComparatorLTE = "~lte"
)

//noinspection GoUnusedConst

const (
	// Start date and time of an event, unix timestamp
	FilterEventStartDateTime = "search_api_aggregation_1"
	// End date and time of an event, unix timestamp
	FilterEventEndDateTime = "search_api_aggregation_3"
	// City an event is taking place at.
	FilterEventLocation = "field_offline:field_address:locality"
)

// Creates a filter which can be passed to SearchBuilder.Filters().
//
// See type Filter for a detailed description.
func CreateFilter(operator string, field string, comparator string, value string) Filter {
	return &filter{operator, field, comparator, value}
}

// Creates a range filter which can be passed to SearchBuilder.Filters().
//
// See type Filter for a detailed description.
func CreaterRangeFilter(field string, from string, to string) Filter {
	return &rangeFilter{field, from, to}
}
