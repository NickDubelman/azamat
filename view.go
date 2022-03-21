package azamat

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// View can be used when an entity doesn't map precisely to a specific table. This is
// not to be confused with "SQL views," but it is similar in concept. A View is
// associated with a custom query. This can be useful when you have an entity where
// some of the fields come from joining multiple tables. Often, you may be joining
// two tables that both have an ID column-- in this case you need to specify IDFrom
// so that GetByID knows which table's ID to use
type View[T any] struct {
	IDFrom fmt.Stringer
	Query  func() sq.SelectBuilder
}

func (v View[T]) GetAll(db *sqlx.DB) ([]T, error) {
	sql, args, err := v.Query().ToSql()
	if err != nil {
		return nil, err
	}

	var rows []T
	err = db.Select(&rows, sql, args...)
	return rows, err
}

func (v View[T]) GetByID(db *sqlx.DB, id int) (T, error) {
	var row T

	idColumn := "id"
	idFrom := v.IDFrom.String()
	if idFrom != "" {
		idColumn = fmt.Sprintf("%s.id", idFrom)
	}

	sql, args, err := v.Query().Where(sq.Eq{idColumn: id}).ToSql()
	if err != nil {
		return row, err
	}

	var rows []T
	if err := db.Select(&rows, sql, args...); err != nil {
		return row, err
	}

	if len(rows) == 0 {
		return row, fmt.Errorf("none found")
	}

	if len(rows) != 1 {
		return row, fmt.Errorf("expected to only get row")
	}

	return rows[0], nil
}

func (v View[T]) GetByIDs(db *sqlx.DB, ids ...int) ([]T, error) {
	idColumn := "id"
	idFrom := v.IDFrom.String()
	if idFrom != "" {
		idColumn = fmt.Sprintf("%s.id", idFrom)
	}

	sql, args, err := v.Query().Where(sq.Eq{idColumn: ids}).ToSql()
	if err != nil {
		return nil, err
	}

	var rows []T
	err = db.Select(&rows, sql, args...)
	return rows, err
}
