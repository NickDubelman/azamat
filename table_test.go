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
	t.Error("epic failure")
}

func TestTableGetByIDs(t *testing.T) {
	t.Error("epic failure")
}

func TestCreate(t *testing.T) {
	t.Error("epic failure")
}

func TestCreateIfNotExists(t *testing.T) {
	t.Error("epic failure")
}

func TestPrefixColumns(t *testing.T) {
	t.Error("epic failure")
}
