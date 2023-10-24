package dml

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tChange struct {
	field any
	value any
}

func Set[T any](field *T, value Expr[T]) tChange {
	return tChange{field: field, value: value}
}

type tUpdate[T internal.Model] struct {
	model   *T
	changes []tChange
	conds   []Expr[bool]
}

func Update[T internal.Model](model *T, changes ...tChange) tUpdate[T] {
	return tUpdate[T]{
		model:   model,
		changes: changes,
	}
}

func (u tUpdate[T]) Where(conds ...Expr[bool]) tUpdate[T] {
	u.conds = append(u.conds, conds...)
	return u
}

func (u tUpdate[T]) Squirrel(conf dbconf.Config) (squirrel.Sqlizer, error) {
	conf = conf.WithModel(u.model)
	// make builder, set column names and table name
	q := squirrel.Update(internal.GetTableName(conf, u.model))
	q = q.PlaceholderFormat(conf.SquirrelPlaceholder())

	// generate SET clause
	for _, change := range u.changes {
		fname, err := internal.GetColumnName(conf, change.field)
		if err != nil {
			return nil, fmt.Errorf("get field name: %v", err)
		}
		q = q.Set(fname, change.value)
	}

	// generate WHERE clause
	if len(u.conds) != 0 {
		preds := tokens.New()
		first := true
		for _, pred := range u.conds {
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
