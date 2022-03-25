package azamat

import "github.com/jmoiron/sqlx"

// Connect just wraps sqlx.Connect
func Connect(driverName, dataSourceName string) (*sqlx.DB, error) {
	return sqlx.Connect(driverName, dataSourceName)
}
