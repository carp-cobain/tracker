package domain

// Page holds a slice of domain objects and a cursor that points to the next page.
type Page[T any] struct {
	Cursor uint64 `json:"cursor"`
	Limit  int    `json:"limit"`
	Data   []T    `json:"data"`
	Size   int    `json:"size"`
}

// NewPage creates a new page of data
func NewPage[T any](cursor uint64, limit int, data []T) Page[T] {
	return Page[T]{
		Cursor: cursor,
		Limit:  limit,
		Data:   data,
		Size:   len(data),
	}
}

// IsEmpty returns true if a page has no data (empty slice).
func (self Page[T]) IsEmpty() bool {
	return len(self.Data) == 0
}

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
