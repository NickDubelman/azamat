package azamat

import (
	sq "github.com/Masterminds/squirrel"
)

// Table is a SQL table
type Table[T any] struct {
	Name    string
	Columns []string
}

// Builders

// Select returns a buildable Select query that is bound to a specific table name. If
// no columns are provided, it gets all columns specified by the table
func (t Table[T]) Select(columns ...string) SelectBuilder[T] {
	actualColumns := columns
	if len(columns) == 0 {
		actualColumns = t.Columns
	}

	return SelectBuilder[T]{sq.Select(actualColumns...).From(t.Name)}
}

// Insert returns a buildable Insert statement that is bound to a specific table name
func (t Table[T]) Insert() InsertBuilder {
	return InsertBuilder{sq.Insert(t.Name)}
}

// Update returns a buildable Update statement that is bound to a specific table name
func (t Table[T]) Update() UpdateBuilder {
	return UpdateBuilder{sq.Update(t.Name)}
}

// Delete returns a buildable Delete statement that is bound to a specific table name
func (t Table[T]) Delete() DeleteBuilder {
	return DeleteBuilder{sq.Delete(t.Name)}
}
