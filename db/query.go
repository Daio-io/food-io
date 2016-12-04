package db

// QueryOptions - Search options
type QueryOptions struct {
	Amount  int
	filters map[string]string
}

// NewQuery - Create a new QueryOptions
func NewQuery() QueryOptions {
	qo := QueryOptions{filters: map[string]string{}}
	return qo
}

// AddFilter - Add new filter to query
func (q *QueryOptions) AddFilter(filterKey string, filterValue string) {
	if filterValue != "" {
		q.filters[filterKey] = filterValue
	}
}

// GetFilters - Get map of filters in query
func (q *QueryOptions) GetFilters() map[string]string {
	return q.filters
}
