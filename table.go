package azamat

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// Table is a SQL table
type Table[T any] struct {
	Name      string
	Columns   []string
	RawSchema string
}

func (t Table[T]) String() string {
	return t.Name
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

func (t Table[T]) Create(db *sqlx.DB) error {
	createTable := fmt.Sprintf("CREATE TABLE %s (%s)", t.Name, t.RawSchema)
	_, err := db.Exec(createTable)
	return err
}

func (t Table[T]) CreateIfNotExists(db *sqlx.DB) error {
	createTable := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (%s)", t.Name, t.RawSchema,
	)
	_, err := db.Exec(createTable)
	return err
}

func PrefixColumns(prefix string, columns []string) (cols []string) {
	for _, c := range columns {
		cols = append(cols, fmt.Sprintf("%s.%s", prefix, c))
	}
	return
}

// Builders

// Select returns a buildable Select query that is bound to a specific table name. If
// no columns are provided, it gets all columns specified by the table
func (t Table[T]) Select() SelectBuilder[T] {
	return SelectBuilder[T]{
		sq.Select(PrefixColumns(t.Name, t.Columns)...).From(t.Name),
	}
}

// BasicSelect is like Select but downgrades the builder to be non-generic. This is
// useful when you don't want to select all columns of the table.
func (t Table[T]) BasicSelect(columns ...string) sq.SelectBuilder {
	actualColumns := PrefixColumns(t.Name, columns)
	if len(columns) == 0 {
		actualColumns = PrefixColumns(t.Name, t.Columns)
	}

	return sq.Select(actualColumns...).From(t.Name)
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
