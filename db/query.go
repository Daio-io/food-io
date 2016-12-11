package db

// QueryOptions - Search options
type QueryOptions struct {
	Amount  int
	filters []map[string]string
}

// NewQuery - Create a new QueryOptions
func NewQuery() QueryOptions {
	qo := QueryOptions{filters: []map[string]string{}}
	return qo
}

// AddFilter - Add new filter to query
func (q *QueryOptions) AddFilter(filterKey string, filterValue string) *QueryOptions {
	if filterValue != "" {
		var filter = map[string]string{}
		filter[filterKey] = filterValue
		q.filters = append(q.filters, filter)
	}
	return q
}

// MultiFilter - Add new filter to query
func (q *QueryOptions) MultiFilter(filterKey string, filterValues []string) *QueryOptions {
	for _, element := range filterValues {
		if element != "" {
			var filter = map[string]string{}
			filter[filterKey] = element
			q.filters = append(q.filters, filter)
		}
	}
	return q
}

// GetFilters - Get map of filters in query
func (q *QueryOptions) GetFilters() []map[string]string {
	return q.filters
}
