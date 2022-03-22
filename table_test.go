package azamat

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableGetAll(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")

	type Todo struct {
		ID    int
		Title string
	}

	TodoTable := Table[Todo]{
		Name:    "todos",
		Columns: []string{"id", "title"},
	}

	db.MustExec(`CREATE TABLE todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	)`)

	// When there are no entries in the table...
	todos, err := TodoTable.GetAll(db)
	require.NoError(t, err)
	require.Zero(t, todos)

	// Create some todos
	todo1, todo2 := "assist Borat", "find Pamela"
	db.MustExec(`INSERT INTO todos (title) VALUES (?), (?)`, todo1, todo2)

	// When there are multiple entries in the table...
	todos, err = TodoTable.GetAll(db)
	require.NoError(t, err)
	require.Len(t, todos, 2)
	assert.Equal(t, 1, todos[0].ID)
	assert.Equal(t, todo1, todos[0].Title)
	assert.Equal(t, 2, todos[1].ID)
	assert.Equal(t, todo2, todos[1].Title)
}

func TestTableGetByID(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")

	type Todo struct {
		ID    int
		Title string
	}

	TodoTable := Table[Todo]{
		Name:    "todos",
		Columns: []string{"id", "title"},
	}

	db.MustExec(`CREATE TABLE todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	)`)

	// When there are no entries in the table...
	todo, err := TodoTable.GetByID(db, 420)
	require.Error(t, err)
	require.Empty(t, todo)

	// Create some todos
	todo1, todo2 := "assist Borat", "find Pamela"
	db.MustExec(`INSERT INTO todos (title) VALUES (?), (?)`, todo1, todo2)

	// When there are multiple entries in the table...
	todo, err = TodoTable.GetByID(db, 2)
	require.NoError(t, err)
	assert.Equal(t, 2, todo.ID)
	assert.Equal(t, todo2, todo.Title)
}

func TestTableGetByIDs(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")

	type Todo struct {
		ID    int
		Title string
	}

	TodoTable := Table[Todo]{
		Name:    "todos",
		Columns: []string{"id", "title"},
	}

	db.MustExec(`CREATE TABLE todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	)`)

	// When there are no entries in the table...
	todos, err := TodoTable.GetByIDs(db, 69, 420)
	require.NoError(t, err)
	require.Len(t, todos, 0)

	// Create some todos
	todo1, todo2, todo3 := "assist Borat", "find Pamela", "buy bear food"
	db.MustExec(
		`INSERT INTO todos (title) VALUES (?), (?), (?)`, todo1, todo2, todo3,
	)

	// When there are multiple entries in the table...
	todos, err = TodoTable.GetByIDs(db, 1, 3)
	require.NoError(t, err)
	require.Len(t, todos, 2)
	assert.Equal(t, 1, todos[0].ID)
	assert.Equal(t, todo1, todos[0].Title)
	assert.Equal(t, 3, todos[1].ID)
	assert.Equal(t, todo3, todos[1].Title)
}

func TestCreate(t *testing.T) {
	t.Error("epic failure")
}

func TestCreateIfNotExists(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")

	type Todo struct {
		ID    int
		Title string
	}

	TodoTable := Table[Todo]{
		Name:    "todos",
		Columns: []string{"id", "title"},
		RawSchema: `
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL
		`,
	}

	// When table doesn't exist...
	err := TodoTable.CreateIfNotExists(db)
	require.NoError(t, err)

	// Make sure table was actually created
	db.MustExec("SELECT 1 FROM todos")

	// When table already exists...
	err = TodoTable.CreateIfNotExists(db)
	require.NoError(t, err)

	// Make sure table still exists
	db.MustExec("SELECT 1 FROM todos")
}

func TestPrefixColumns(t *testing.T) {
	t.Error("epic failure")
}
