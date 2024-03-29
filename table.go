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
	IDColumn  string
	Postgres  bool
}

func (t Table[T]) String() string {
	return t.Name
}

func (t Table[T]) IsPostgres() bool {
	return t.Postgres || Postgres
}

func (t Table[T]) GetAll(runner Runner) ([]T, error) {
	return t.Select().All(runner)
}

func (t Table[T]) GetByID(runner Runner, id int) (T, error) {
	idCol := "id"
	if t.IDColumn != "" {
		idCol = t.IDColumn
	}

	return t.Select().Where(sq.Eq{idCol: id}).Only(runner)
}

func (t Table[T]) GetByIDs(runner Runner, ids ...int) ([]T, error) {
	idCol := "id"
	if t.IDColumn != "" {
		idCol = t.IDColumn
	}

	return t.Select().Where(sq.Eq{idCol: ids}).All(runner)
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
	if t.IsPostgres() {
		return SelectBuilder[T]{
			psql.Select(PrefixColumns(t.Name, t.Columns)...).From(t.Name),
		}
	}

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

	if t.IsPostgres() {
		return psql.Select(actualColumns...).From(t.Name)
	}

	return sq.Select(actualColumns...).From(t.Name)
}

// Insert returns a buildable Insert statement that is bound to a specific table name
func (t Table[T]) Insert() InsertBuilder {
	if t.IsPostgres() {
		return InsertBuilder{psql.Insert(t.Name)}
	}

	return InsertBuilder{sq.Insert(t.Name)}
}

// Update returns a buildable Update statement that is bound to a specific table name
func (t Table[T]) Update() UpdateBuilder {
	if t.IsPostgres() {
		return UpdateBuilder{psql.Update(t.Name)}
	}

	return UpdateBuilder{sq.Update(t.Name)}
}

// Delete returns a buildable Delete statement that is bound to a specific table name
func (t Table[T]) Delete() DeleteBuilder {
	if t.IsPostgres() {
		return DeleteBuilder{psql.Delete(t.Name)}
	}

	return DeleteBuilder{sq.Delete(t.Name)}
}
