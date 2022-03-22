package azamat

import (
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestViewGetAll(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")

	type User struct {
		ID   int
		Name string
	}

	type Todo struct {
		ID     int
		Title  string
		Author string
	}

	type TodoRow struct {
		ID       int
		Title    string
		AuthorID int `db:"authorID"`
	}

	UserTable := Table[User]{
		Name:    "users",
		Columns: []string{"id", "name"},
	}

	TodoTable := Table[TodoRow]{
		Name:    "todos",
		Columns: []string{"id", "title", "authorID"},
	}

	TodoView := View[Todo]{
		IDFrom: TodoTable,
		Query: func() sq.SelectBuilder {
			join := fmt.Sprintf(
				"%s ON %s.id = %s.authorID", UserTable, UserTable, TodoTable,
			)

			return TodoTable.
				SelectColumns("id", "title").
				Columns("name AS author").
				Join(join)
		},
	}

	db.MustExec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	)`)

	db.MustExec(`CREATE TABLE todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		authorID INTEGER NOT NULL,
		
		FOREIGN KEY (authorID) REFERENCES users(id)
	)`)

	// When there are no entries in the table...
	todos, err := TodoView.GetAll(db)
	require.NoError(t, err)
	require.Zero(t, todos)

	// Create some users
	user1, user2 := "Azamat", "Borat"
	db.MustExec(`INSERT INTO users (name) VALUES (?), (?)`, user1, user2)

	// Create some todos
	todo1, todo2 := "assist Borat", "find Pamela"
	db.MustExec(
		`INSERT INTO todos (title, authorID) VALUES (?, 2), (?, 1)`, todo1, todo2,
	)

	// When there are multiple entries in the table...
	todos, err = TodoView.GetAll(db)
	require.NoError(t, err)
	require.Len(t, todos, 2)
	assert.Equal(t, 1, todos[0].ID)
	assert.Equal(t, todo1, todos[0].Title)
	assert.Equal(t, user2, todos[0].Author)
	assert.Equal(t, 2, todos[1].ID)
	assert.Equal(t, todo2, todos[1].Title)
	assert.Equal(t, user1, todos[1].Author)
}

func TestViewGetByID(t *testing.T) {
	t.Error("epic failure")
}

func TestViewGetByIDs(t *testing.T) {
	t.Error("epic failure")
}
