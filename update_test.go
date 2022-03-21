package azamat

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateRun(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")

	type Todo struct {
		ID    int
		Title string
	}

	db.MustExec(`CREATE TABLE todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	)`)

	// First, create some entries
	todo1, todo2 := "assist Borat", "find Pamela"
	db.MustExec(`INSERT INTO todos (title) VALUES (?), (?)`, todo1, todo2)

	// When updating a single entry...
	todo1 = "buy bear food"
	todo1ID := 1
	update := Update("todos").Set("title", todo1).Where("id = ?", todo1ID)
	_, err := update.Run(db)
	require.NoError(t, err)

	// Make sure entry was actually updated
	var rows []Todo
	err = db.Select(&rows, "SELECT id, title FROM todos WHERE id = ?", todo1ID)
	require.NoError(t, err)
	require.Len(t, rows, 1)
	assert.Equal(t, todo1, rows[0].Title)

	// When updating multiple entries...
	update = Update("todos").Set("title", todo2)
	_, err = update.Run(db)
	require.NoError(t, err)

	// Make sure entries were actually updated
	rows = nil
	err = db.Select(&rows, "SELECT id, title FROM todos")
	require.NoError(t, err)
	require.Len(t, rows, 2)
	assert.Equal(t, todo2, rows[0].Title)
	assert.Equal(t, todo2, rows[1].Title)
}
