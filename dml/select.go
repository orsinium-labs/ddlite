package dml

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type Scanner[T internal.Model] func(*sql.Rows) error

type tSelectModel[T internal.Model] struct {
	fields []any
	conds  []Expr[bool]
	model  *T
}

func Select[T internal.Model](model *T, fields ...any) tSelectModel[T] {
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

func (s tSelectModel[T]) Squirrel(conf dbconf.Config) (squirrel.Sqlizer, error) {
	conf = conf.WithModel(s.model)
	fnames := make([]string, 0, len(s.fields))
	for _, f := range s.fields {
		fname, err := internal.GetColumnName(conf, f)
		if err != nil {
			return squirrel.SelectBuilder{}, fmt.Errorf("get column name: %v", err)
		}
		fnames = append(fnames, fname)
	}
	q := squirrel.Select(fnames...)
	q = q.PlaceholderFormat(conf.SquirrelPlaceholder())
	q = q.From(internal.GetTableName(conf, s.model))

	if len(s.conds) != 0 {
		preds := tokens.New()
		first := true
		for _, pred := range s.conds {
			if first {
				first = false
			} else {
				preds.Add(tokens.Keyword("AND"))
			}
			preds.Extend(pred.Tokens(conf))
		}
		sql, args, err := preds.SQL(conf)
		if err != nil {
			return nil, fmt.Errorf("generate SQL for predicates: %w", err)
		}
		q = q.Where(squirrel.Expr(sql, args...))
	}

	return q, nil
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

func (s tSelectModel[T]) String() string {
	conf := dbconf.New("sqlite3").WithModel(s)
	builder, _ := s.Squirrel(dbconf.New("sqlite3").WithModel(conf))
	sql, _, _ := builder.ToSql()
	return sql
}
