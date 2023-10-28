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
	whereClause
	limitClause
	fields []any
	model  *T
}

func Select[T internal.Model](model *T, fields ...any) tSelectModel[T] {
	return tSelectModel[T]{model: model, fields: fields}
}

func (s tSelectModel[T]) Where(predicates ...Expr[bool]) tSelectModel[T] {
	s.predicates = append(s.predicates, predicates...)
	return s
}

func (s tSelectModel[T]) Limit(expr Expr[int]) tSelectModel[T] {
	s.limit = expr
	return s
}

func (s tSelectModel[T]) Offset(expr Expr[int]) tSelectModel[T] {
	s.offset = expr
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
	ts.Extend(s.buildWhere(conf))
	ts.Extend(s.buildLimit(conf))
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
