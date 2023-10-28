package dml

import (
	"database/sql"
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type Scanner[T internal.Model] func(*sql.Rows) error

type tSelectModel[T internal.Model] struct {
	where        // WHERE clause
	order        // ORDER BY clause
	limit        // LIMIT and OFFSET clauses
	fields []any // columns to select
	model  *T    // the table to select from (FROM clause)
}

func Select[T internal.Model](model *T, fields ...any) tSelectModel[T] {
	return tSelectModel[T]{model: model, fields: fields}
}

// Where adds WHERE clause to the SELECT statement.
//
// If multiple expressions are passed or Where called multiple times,
// all the given expressions are concatenated with AND operator.
// In other words, all the given expressions must be true for a row to be selected.
func (s tSelectModel[T]) Where(predicates ...Expr[bool]) tSelectModel[T] {
	s.predicates = append(s.predicates, predicates...)
	return s
}

// OrderBy adds ORDER BY clause to the SELECT statement.
//
// The arguments must be constructed using the [Ordering] function.
func (s tSelectModel[T]) OrderBy(ords ...iOrdering) tSelectModel[T] {
	s.order.ords = ords
	return s
}

// Limit adds LIMIT clause to the SELECT statement.
//
// If you specify Limit, you should also specify OrderBy.
// Otherwise, the result is non-deterministic. There is no default ordering.
func (s tSelectModel[T]) Limit(expr Expr[int]) tSelectModel[T] {
	s.limit.limit = expr
	return s
}

// Offset adds OFFSET clause to the SELECT statement.
//
// If you specify Offset, you should also specify OrderBy.
// Otherwise, the result is non-deterministic. There is no default ordering.
func (s tSelectModel[T]) Offset(expr Expr[int]) tSelectModel[T] {
	s.limit.offset = expr
	return s
}

// And is an alias for Where.
func (s tSelectModel[T]) And(conds ...Expr[bool]) tSelectModel[T] {
	return s.Where(conds...)
}

func (s tSelectModel[T]) Tokens(conf dbconf.Config) tokens.Tokens {
	conf = conf.WithModel(s.model)
	ts := tokens.New(tokens.Keyword("SELECT"))
	for i, f := range s.fields {
		if i > 0 {
			ts.Add(tokens.Comma())
		}
		ts.Add(internal.GetColumnName(conf, f))
	}
	ts.Add(
		tokens.Keyword("FROM"),
		internal.GetTableName(conf, s.model),
	)
	ts.Extend(s.where.build(conf))
	ts.Extend(s.order.build(conf))
	ts.Extend(s.limit.build(conf))
	return ts
}

func (s tSelectModel[T]) Scanner(conf dbconf.Config, target *T) (Scanner[T], error) {
	conf = conf.WithModel(s.model)
	cols := make([]any, 0, len(s.fields))
	for _, field := range s.fields {
		fieldName, err := internal.GetFieldName(conf, field)
		if err != nil {
			return nil, fmt.Errorf("get field name: %v", err)
		}
		col, err := internal.GetField(target, fieldName)
		if err != nil {
			return nil, fmt.Errorf("get struct field by name: %v", err)
		}
		cols = append(cols, col)
	}

	scan := func(rows *sql.Rows) error {
		err := rows.Scan(cols...)
		if err != nil {
			return fmt.Errorf("rows scan: %w", err)
		}
		return nil
	}
	return scan, nil
}
