package azamat

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Not sure if this test is actually useful or how to better test this...
// The function CommitTransaction follows a pattern that seems to be pretty common
// for gophers:
//   - https://entgo.io/docs/transactions#best-practices
//   - https://pseudomuto.com/2018/01/clean-sql-transactions-in-golang/
//   - https://stackoverflow.com/questions/16184238/database-sql-tx-detecting-commit-or-rollback
func TestCommitTransaction(t *testing.T) {
	db, _ := sqlx.Open("sqlite3", ":memory:")

	type Todo struct {
		ID    int
		Title string
	}

	db.MustExec(`CREATE TABLE todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	)`)

	// When transaction has an error...
	err := CommitTransaction(db, func(tx *sqlx.Tx) error {
		return fmt.Errorf("999 syntax error")
	})
	require.Error(t, err)

	// Make sure we can still query table

	// When transaction panics...
	assert.Panics(t, func() {
		err = CommitTransaction(db, func(tx *sqlx.Tx) error {
			panic("999 syntax error")
		})
	})
	require.Error(t, err)

	// Make sure we can still query table

	// Simple transaction...
	err = CommitTransaction(db, func(tx *sqlx.Tx) error {
		_, err = tx.Exec("INSERT INTO todos (title) VALUES ('very nice')")
		require.NoError(t, err)
		return err
	})
	require.NoError(t, err)

	// Make sure entry was actually inserted
	var rows []Todo
	db.Select(&rows, "SELECT * FROM todos")
	require.Len(t, rows, 1)
}
