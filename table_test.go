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

	// When the id column is not named id
	type User struct {
		ID   int `db:"user_id"`
		Name string
	}

	UserTable := Table[User]{
		Name:     "users",
		Columns:  []string{"user_id", "name"},
		IDColumn: "user_id",
	}

	db.MustExec(`CREATE TABLE users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	)`)

	// Create some users
	user1, user2 := "Borat", "Pamela"
	db.MustExec(`INSERT INTO users (name) VALUES (?), (?)`, user1, user2)

	// When there are multiple entries in the table...
	user, err := UserTable.GetByID(db, 2)
	require.NoError(t, err)
	assert.Equal(t, 2, user.ID)
	assert.Equal(t, user2, user.Name)
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

	// When the id column is not named id
	type User struct {
		ID   int `db:"user_id"`
		Name string
	}

	UserTable := Table[User]{
		Name:     "users",
		Columns:  []string{"user_id", "name"},
		IDColumn: "user_id",
	}

	db.MustExec(`CREATE TABLE users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	)`)

	// Create some users
	user1, user2 := "Borat", "Pamela"
	db.MustExec(`INSERT INTO users (name) VALUES (?), (?)`, user1, user2)

	// When there are multiple entries in the table...
	users, err := UserTable.GetByIDs(db, 1, 2)
	require.NoError(t, err)
	require.Len(t, users, 2)
	assert.Equal(t, 1, users[0].ID)
	assert.Equal(t, user1, users[0].Name)
	assert.Equal(t, 2, users[1].ID)
	assert.Equal(t, user2, users[1].Name)
}

func TestCreate(t *testing.T) {
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
	err := TodoTable.Create(db)
	require.NoError(t, err)

	// Make sure table was actually created
	db.MustExec("SELECT 1 FROM todos")

	// When table already exists...
	err = TodoTable.Create(db)
	require.Error(t, err)
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
	type testCase struct {
		inputPrefix  string
		inputColumns []string
		expected     []string
	}

	cases := []testCase{
		{
			inputPrefix:  "yekshi",
			inputColumns: []string{"mesh"},
			expected:     []string{"yekshi.mesh"},
		},

		{
			inputPrefix:  "a",
			inputColumns: []string{"b", "C", "d", "E"},
			expected:     []string{"a.b", "a.C", "a.d", "a.E"},
		},

		{
			inputPrefix:  "",
			inputColumns: []string{"b", "C", "d", "E"},
			expected:     []string{".b", ".C", ".d", ".E"},
		},
	}

	for _, c := range cases {
		actual := PrefixColumns(c.inputPrefix, c.inputColumns)
		require.Equal(t, c.expected, actual)
	}
}
