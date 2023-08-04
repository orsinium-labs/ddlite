package qb

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type tSelectModel[T Model] struct {
	fields []any
	conds  []Expr[bool]
	model  *T
}

func Select[T Model](model *T, fields ...any) tSelectModel[T] {
	return tSelectModel[T]{model: model, fields: fields}
}

func (s tSelectModel[T]) Where(conds ...Expr[bool]) tSelectModel[T] {
	s.conds = append(s.conds, conds...)
	return s
}

// And is an alias for Where.
func (s tSelectModel[T]) And(conds ...Expr[bool]) tSelectModel[T] {
	return s.Where(conds...)
}

func (s tSelectModel[T]) Squirrel(...Model) (squirrel.Sqlizer, error) {
	fnames := make([]string, 0, len(s.fields))
	for _, f := range s.fields {
		fname, err := getFieldName(s.model, f)
		if err != nil {
			return squirrel.SelectBuilder{}, fmt.Errorf("get column name: %v", err)
		}
		fnames = append(fnames, fname)
	}
	q := squirrel.Select(fnames...)
	q = q.PlaceholderFormat(squirrel.Dollar)
	q = q.From(getModelName(s.model))

	if len(s.conds) != 0 {
		preds := make([]squirrel.Sqlizer, 0, len(s.conds))
		for _, c := range s.conds {
			preds = append(preds, c.Squirrel(s.model))
		}
		q = q.Where(squirrel.And(preds))
	}

	return q, nil
}

func (s tSelectModel[T]) Scanner() (Scanner[T], error) {
	var r T
	cols := make([]any, 0, len(s.fields))
	for _, field := range s.fields {
		fieldName, err := getFieldName(s.model, field)
		if err != nil {
			return nil, fmt.Errorf("get field name: %v", err)
		}
		col, err := getField(&r, fieldName)
		if err != nil {
			return nil, fmt.Errorf("get struct field by name: %v", err)
		}
		cols = append(cols, col)
	}

	scan := func(rows *sql.Rows) (T, error) {
		err := rows.Scan(cols...)
		if err != nil {
			return r, fmt.Errorf("rows scan: %w", err)
		}
		return r, nil
	}
	return scan, nil
}

func (s tSelectModel[T]) String() string {
	builder, _ := s.Squirrel()
	sql, _, _ := builder.ToSql()
	return sql
}
