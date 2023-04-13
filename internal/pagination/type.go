package pagination

import "github.com/sekalahita/epirus/internal/ent/gen"

type order string

const (
	OrderAscending  order = "ASC"
	OrderDescending order = "DESC"
)

type Cursor struct {
	*gen.Cursor
}

type CursorPagination struct {
	Order order

	First  *int
	Last   *int
	Before Cursor
	After  Cursor
}
