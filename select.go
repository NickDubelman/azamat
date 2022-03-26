package azamat

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type SelectBuilder[T any] struct {
	sq.SelectBuilder
}

func Select[T any](columns ...string) SelectBuilder[T] {
	return SelectBuilder[T]{
		SelectBuilder: sq.Select(columns...),
	}
}

func (b SelectBuilder[T]) All(runner Runner) ([]T, error) {
	sql, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	var rows []T
	err = runner.Select(&rows, sql, args...)
	return rows, err
}

func (b SelectBuilder[T]) Only(runner Runner) (T, error) {
	var row T

	sql, args, err := b.ToSql()
	if err != nil {
		return row, err
	}

	var rows []T
	if err := runner.Select(&rows, sql, args...); err != nil {
		return row, err
	}

	if len(rows) == 0 {
		return row, fmt.Errorf("none found")
	}

	if len(rows) != 1 {
		return row, fmt.Errorf("expected to only get row")
	}

	return rows[0], nil
}

func (b SelectBuilder[T]) PlaceholderFormat(f sq.PlaceholderFormat) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.PlaceholderFormat(f)
	return b
}

func (b SelectBuilder[T]) RunWith(runner sq.BaseRunner) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.RunWith(runner)
	return b
}

func (b SelectBuilder[T]) Exec() (sql.Result, error) {
	return b.SelectBuilder.Exec()
}

func (b SelectBuilder[T]) Prefix(sql string, args ...interface{}) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.Prefix(sql, args...)
	return b
}

func (b SelectBuilder[T]) PrefixExpr(expr sq.Sqlizer) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.PrefixExpr(expr)
	return b
}

func (b SelectBuilder[T]) Distinct() SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.Distinct()
	return b
}

func (b SelectBuilder[T]) Options(options ...string) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.Options(options...)
	return b
}

// Downgrades the builder to a regular squirrel builder
func (b SelectBuilder[T]) Columns(columns ...string) sq.SelectBuilder {
	return b.SelectBuilder.Columns(columns...)
}

// Downgrades the builder to a regular squirrel builder
func (b SelectBuilder[T]) Column(
	column interface{}, args ...interface{},
) sq.SelectBuilder {
	return b.SelectBuilder.Column(column, args...)
}

func (b SelectBuilder[T]) From(from string) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.From(from)
	return b
}

func (b SelectBuilder[T]) JoinClause(pred interface{}, args ...interface{}) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.JoinClause(pred, args...)
	return b
}

func (b SelectBuilder[T]) Join(join string, rest ...interface{}) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.Join(join, rest...)
	return b
}

func (b SelectBuilder[T]) LeftJoin(join string, rest ...interface{}) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.LeftJoin(join, rest...)
	return b
}

func (b SelectBuilder[T]) RightJoin(join string, rest ...interface{}) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.RightJoin(join, rest...)
	return b
}

func (b SelectBuilder[T]) InnerJoin(join string, rest ...interface{}) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.InnerJoin(join, rest...)
	return b
}

func (b SelectBuilder[T]) CrossJoin(join string, rest ...interface{}) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.CrossJoin(join, rest...)
	return b
}

func (b SelectBuilder[T]) Where(pred interface{}, args ...interface{}) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.Where(pred, args...)
	return b
}

func (b SelectBuilder[T]) GroupBy(groupBys ...string) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.GroupBy(groupBys...)
	return b
}

func (b SelectBuilder[T]) Having(pred interface{}, rest ...interface{}) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.Having(pred, rest...)
	return b
}

func (b SelectBuilder[T]) OrderByClause(pred interface{}, args ...interface{}) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.OrderByClause(pred, args...)
	return b
}

func (b SelectBuilder[T]) OrderBy(orderBys ...string) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.OrderBy(orderBys...)
	return b
}

func (b SelectBuilder[T]) Limit(limit uint64) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.Limit(limit)
	return b
}

func (b SelectBuilder[T]) RemoveLimit() SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.RemoveLimit()
	return b
}

func (b SelectBuilder[T]) Offset(offset uint64) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.Offset(offset)
	return b
}

func (b SelectBuilder[T]) RemoveOffset() SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.RemoveOffset()
	return b
}

func (b SelectBuilder[T]) Suffix(sql string, args ...interface{}) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.Suffix(sql, args...)
	return b
}

func (b SelectBuilder[T]) SuffixExpr(expr sq.Sqlizer) SelectBuilder[T] {
	b.SelectBuilder = b.SelectBuilder.SuffixExpr(expr)
	return b
}
