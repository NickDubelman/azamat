package azamat

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// Table is a SQL table
type Table[T any] struct {
	Name    string
	Columns []string
}

func (t Table[T]) GetAll(db *sqlx.DB) ([]T, error) {
	return t.Select().All(db)
}

func (t Table[T]) GetByID(db *sqlx.DB, id int) (T, error) {
	return t.Select().Where("id = ?", id).Only(db)
}

func (t Table[T]) GetByIDs(db *sqlx.DB, ids ...int) ([]T, error) {
	return t.Select().Where(sq.Eq{"id": ids}).All(db)
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
