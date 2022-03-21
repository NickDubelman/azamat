package azamat

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteRun(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")

	type Todo struct {
		ID    int
		Title string
	}

	db.MustExec(`CREATE TABLE todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	)`)

	// Create some todos
	todo1, todo2, todo3 := "assist Borat", "find Pamela", "buy bear food"
	db.MustExec(`INSERT INTO todos (title) VALUES (?), (?)`, todo1, todo2, todo3)

	// When deleting a single entry...
	delete := Delete("todos").Where("id = ?", 2)
	_, err := delete.Run(db)
	require.NoError(t, err)

	// Make sure delete actually happened
	var rows []Todo
	err = db.Select(&rows, "SELECT id, title FROM todos WHERE id = 2")
	require.NoError(t, err)
	assert.Zero(t, rows)

	// When deleting multiple entries...
	delete = Delete("todos") // no WHERE, so deletes all entries
	_, err = delete.Run(db)
	require.NoError(t, err)

	// Make sure there are no entries left in the table
	err = db.Select(&rows, "SELECT id, title FROM todos")
	require.NoError(t, err)
	assert.Zero(t, rows)
}
