package azamat

import "github.com/jmoiron/sqlx"

// Runner can be a *sqlx.DB or a *sqlx.Tx. This allows us to write code that can be
// run as a standalone statement or as part of a transaction
type Runner interface {
	sqlx.Ext
	Select(dest any, query string, args ...any) error
	Get(dest any, query string, args ...any) error
}
