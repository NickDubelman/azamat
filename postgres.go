package azamat

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// Postgres mode can be enabled globally (if all of your tables are Postgres) or on a
// table by table basis (if you have a mix of Postgres and non-Postgres tables)
var Postgres = false

// Postgres has a different way of getting the last inserted ID. We implement this
// so that MySQL, sqlite, and Postgres behave the same when using azamat. For MySQL
// and sqlite, we can just use result.LastInsertId from the standard sql lib
func runPostgresInsert(b InsertBuilder, runner Runner) (int, error) {
	b = b.Suffix("RETURNING id")

	var lastInsertIDs []int
	rows, err := b.RunWith(runner).Query()
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
		lastInsertIDs = append(lastInsertIDs, id)
	}
	if err := rows.Err(); err != nil {
		return 0, err
	}

	if len(lastInsertIDs) == 0 {
		return 0, fmt.Errorf("could not get last insert ID")
	}

	// We return the /last/ item in lastInsertIDs
	return lastInsertIDs[len(lastInsertIDs)-1], nil
}

// Postgres uses $ instead of ?
var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
