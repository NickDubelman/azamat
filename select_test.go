package azamat

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelectAll(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")

	type Todo struct {
		ID    int
		Title string
	}

	db.MustExec(`CREATE TABLE todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	)`)

	// When there are no entries in the table...
	query := Select[Todo]("id", "title").From("todos")

	todos, err := query.All(db)
	require.NoError(t, err)
	require.Zero(t, todos)

	// Create some todos
	todo1, todo2 := "assist Borat", "find Pamela"
	db.MustExec(`INSERT INTO todos (title) VALUES (?), (?)`, todo1, todo2)

	// When there are multiple entries in the table...
	todos, err = query.All(db)
	require.NoError(t, err)
	require.Len(t, todos, 2)

	// When not just fetching all...
	query = query.Where("id = ?", 2)
	todos, err = query.All(db)
	require.NoError(t, err)
	require.Len(t, todos, 1)
	assert.Equal(t, todo2, todos[0].Title)
}

func TestSelectOnly(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")

	type Todo struct {
		ID    int
		Title string
	}

	db.MustExec(`CREATE TABLE todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	)`)

	// When there are no entries in the table...
	query := Select[Todo]("id", "title").From("todos")

	todo, err := query.Only(db)
	require.Error(t, err)

	// Create some todos
	todo1, todo2 := "assist Borat", "find Pamela"
	db.MustExec(`INSERT INTO todos (title) VALUES (?), (?)`, todo1, todo2)

	// When there are multiple entries in the result...
	todo, err = query.Only(db)
	require.Error(t, err)

	// When there is only a single entry in the result...
	query = query.Where("id = ?", 2)
	todo, err = query.Only(db)
	require.NoError(t, err)
	assert.Equal(t, todo2, todo.Title)
}
