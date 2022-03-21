package azamat

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type UpdateBuilder struct {
	sq.UpdateBuilder
}

func (b UpdateBuilder) Run(runner Runner) (sql.Result, error) {
	return b.RunWith(runner).Exec()
}

func (b UpdateBuilder) PlaceholderFormat(f sq.PlaceholderFormat) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.PlaceholderFormat(f)
	return b
}

func (b UpdateBuilder) RunWith(runner sq.BaseRunner) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.RunWith(runner)
	return b
}

func (b UpdateBuilder) Prefix(sql string, args ...interface{}) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.Prefix(sql, args...)
	return b
}

func (b UpdateBuilder) PrefixExpr(expr sq.Sqlizer) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.PrefixExpr(expr)
	return b
}

func (b UpdateBuilder) Table(table string) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.Table(table)
	return b
}

func (b UpdateBuilder) Set(column string, value interface{}) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.Set(column, value)
	return b
}

func (b UpdateBuilder) SetMap(clauses map[string]interface{}) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.SetMap(clauses)
	return b
}

func (b UpdateBuilder) Where(pred interface{}, args ...interface{}) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.Where(pred, args...)
	return b
}

func (b UpdateBuilder) OrderBy(orderBys ...string) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.OrderBy(orderBys...)
	return b
}

func (b UpdateBuilder) Limit(limit uint64) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.Limit(limit)
	return b
}

func (b UpdateBuilder) Offset(offset uint64) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.Offset(offset)
	return b
}

func (b UpdateBuilder) Suffix(sql string, args ...interface{}) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.Suffix(sql, args...)
	return b
}

func (b UpdateBuilder) SuffixExpr(expr sq.Sqlizer) UpdateBuilder {
	b.UpdateBuilder = b.UpdateBuilder.SuffixExpr(expr)
	return b
}
