package azamat

import (
	sq "github.com/Masterminds/squirrel"
)

// Postgres mode can be enabled globally (if all of your tables are Postgres) or on a
// table by table basis (if you have a mix of Postgres and non-Postgres tables)
var Postgres = false

// Postgres uses $ instead of ?
var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
