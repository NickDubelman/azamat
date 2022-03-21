package azamat

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsertRun(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")

	type Todo struct {
		ID    int
		Title string
	}

	db.MustExec(`CREATE TABLE todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	)`)

	// When creating a single entry...
	todo1 := "buy bear food"
	insert := Insert("todos").Columns("title").Values(todo1)
	todoID, err := insert.Run(db)
	require.NoError(t, err)
	require.NotZero(t, todoID)

	// Make sure entry was actually inserted
	var rows []Todo
	err = db.Select(&rows, "SELECT id, title FROM todos WHERE id = ?", todoID)
	require.NoError(t, err)
	assert.Len(t, rows, 1)

	// When creating multiple entries...
	todo2, todo3 := "buy Korky Buchek album", "fuel van"
	insert = Insert("todos").Columns("title").Values(todo2).Values(todo3)
	todoID, err = insert.Run(db)
	require.NoError(t, err)
	require.NotZero(t, todoID)

	// Make sure entries were actually inserted
	rows = nil
	err = db.Select(&rows, "SELECT id, title FROM todos WHERE id = ?", todoID)
	require.NoError(t, err)
	assert.Len(t, rows, 1)

	rows = nil
	err = db.Select(&rows, "SELECT id, title FROM todos WHERE id = ?", todoID-1)
	require.NoError(t, err)
	assert.Len(t, rows, 1)
}
