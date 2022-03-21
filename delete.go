package azamat

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type DeleteBuilder struct {
	sq.DeleteBuilder
}

func Delete[T any](from string) DeleteBuilder {
	return DeleteBuilder{sq.Delete(from)}
}

func (b DeleteBuilder) Run(runner Runner) (sql.Result, error) {
	return b.RunWith(runner).Exec()
}

func (b DeleteBuilder) PlaceholderFormat(f sq.PlaceholderFormat) DeleteBuilder {
	b.DeleteBuilder = b.DeleteBuilder.PlaceholderFormat(f)
	return b
}

func (b DeleteBuilder) RunWith(runner sq.BaseRunner) DeleteBuilder {
	b.DeleteBuilder = b.DeleteBuilder.RunWith(runner)
	return b
}

func (b DeleteBuilder) Prefix(sql string, args ...interface{}) DeleteBuilder {
	b.DeleteBuilder = b.DeleteBuilder.Prefix(sql, args...)
	return b
}

func (b DeleteBuilder) PrefixExpr(expr sq.Sqlizer) DeleteBuilder {
	b.DeleteBuilder = b.DeleteBuilder.PrefixExpr(expr)
	return b
}

func (b DeleteBuilder) From(from string) DeleteBuilder {
	b.DeleteBuilder = b.DeleteBuilder.From(from)
	return b
}

func (b DeleteBuilder) Where(pred interface{}, args ...interface{}) DeleteBuilder {
	b.DeleteBuilder = b.DeleteBuilder.Where(pred, args...)
	return b
}

func (b DeleteBuilder) OrderBy(orderBys ...string) DeleteBuilder {
	b.DeleteBuilder = b.DeleteBuilder.OrderBy(orderBys...)
	return b
}

func (b DeleteBuilder) Limit(limit uint64) DeleteBuilder {
	b.DeleteBuilder = b.DeleteBuilder.Limit(limit)
	return b
}

func (b DeleteBuilder) Offset(offset uint64) DeleteBuilder {
	b.DeleteBuilder = b.DeleteBuilder.Offset(offset)
	return b
}

func (b DeleteBuilder) Suffix(sql string, args ...interface{}) DeleteBuilder {
	b.DeleteBuilder = b.DeleteBuilder.Suffix(sql, args...)
	return b
}

func (b DeleteBuilder) SuffixExpr(expr sq.Sqlizer) DeleteBuilder {
	b.DeleteBuilder = b.DeleteBuilder.SuffixExpr(expr)
	return b
}
