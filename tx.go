package azamat

import "github.com/jmoiron/sqlx"

// CommitTransaction should always be called when running a database transaction. It
// takes care of calling Commit/Rollback as necessary. This prevents devs from
// forgetting to call these (which will lead to locked db connections).
func CommitTransaction(db *sqlx.DB, fn func(tx *sqlx.Tx) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
