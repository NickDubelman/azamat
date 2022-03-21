package azamat

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type InsertBuilder struct {
	sq.InsertBuilder
}

func (b InsertBuilder) Run(db *sqlx.DB) (sql.Result, error) {
	return b.RunWith(db).Exec()
}

func (b InsertBuilder) PlaceholderFormat(f sq.PlaceholderFormat) InsertBuilder {
	b.InsertBuilder = b.InsertBuilder.PlaceholderFormat(f)
	return b
}

func (b InsertBuilder) RunWith(runner sq.BaseRunner) InsertBuilder {
	b.InsertBuilder = b.InsertBuilder.RunWith(runner)
	return b
}

func (b InsertBuilder) Prefix(sql string, args ...interface{}) InsertBuilder {
	b.InsertBuilder = b.InsertBuilder.Prefix(sql, args...)
	return b
}

func (b InsertBuilder) PrefixExpr(expr sq.Sqlizer) InsertBuilder {
	b.InsertBuilder = b.InsertBuilder.PrefixExpr(expr)
	return b
}

func (b InsertBuilder) Options(options ...string) InsertBuilder {
	b.InsertBuilder = b.InsertBuilder.Options(options...)
	return b
}

func (b InsertBuilder) Into(from string) InsertBuilder {
	b.InsertBuilder = b.InsertBuilder.Into(from)
	return b
}

func (b InsertBuilder) Columns(columns ...string) InsertBuilder {
	b.InsertBuilder = b.InsertBuilder.Columns(columns...)
	return b
}

func (b InsertBuilder) Values(values ...interface{}) InsertBuilder {
	b.InsertBuilder = b.InsertBuilder.Values(values...)
	return b
}

func (b InsertBuilder) Suffix(sql string, args ...interface{}) InsertBuilder {
	b.InsertBuilder = b.InsertBuilder.Suffix(sql, args...)
	return b
}

func (b InsertBuilder) SuffixExpr(expr sq.Sqlizer) InsertBuilder {
	b.InsertBuilder = b.InsertBuilder.SuffixExpr(expr)
	return b
}

func (b InsertBuilder) SetMap(clauses map[string]interface{}) InsertBuilder {
	b.InsertBuilder = b.InsertBuilder.SetMap(clauses)
	return b
}

func (b InsertBuilder) Select(sb sq.SelectBuilder) InsertBuilder {
	b.InsertBuilder = b.InsertBuilder.Select(sb)
	return b
}
