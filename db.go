package azamat

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// DBInfo contains the connection info for a SQL database
type DBInfo struct {
	Dialect   string
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

// Defaults for DBInfo.Connect
const (
	ConnectDefaultUsername  = "root"
	ConnectDefaultHostname  = "localhost"
	ConnectDefaultPort      = 3306
	ConnectDefaultParameter = "?charset=utf8&parseTime=True&loc=UTC"
)

// Connect establishes the connection pool to the database
func (info *DBInfo) Connect(
	maxOpenConns int,
	maxIdleConns int,
	connMaxLifetime time.Duration,
) (*sqlx.DB, error) {
	info.applyDefaults() // Apply defaults to fields that are unset

	db, err := sqlx.Connect(info.Dialect, info.DSN())
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)

	return db, nil
}

// DSN returns a "data source name" string that can be used to connect to a database
// via the Go SQL driver.
//
// See: https://github.com/go-sql-driver/mysql#dsn-data-source-name
func (info *DBInfo) DSN() string {
	// Example: root:@tcp(localhost:3306)/test
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s%s",
		info.Username,
		info.Password,
		info.Hostname,
		info.Port,
		info.Name,
		info.Parameter,
	)
}

func (info *DBInfo) applyDefaults() {
	if info.Username == "" {
		info.Username = ConnectDefaultUsername
	}

	if info.Hostname == "" {
		info.Hostname = ConnectDefaultHostname
	}

	if info.Port == 0 {
		info.Port = ConnectDefaultPort
	}

	if info.Parameter == "" {
		info.Parameter = ConnectDefaultParameter
	}
}
