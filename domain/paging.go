package domain

// PageParams are params for querying a page of domain objects.
type PageParams struct {
	Cursor uint64
	Limit  int
}

// NewPageParams creates new pagination parameters.
func NewPageParams(cursor uint64, limit int) PageParams {
	return PageParams{
		Cursor: cursor,
		Limit:  limit,
	}
}
