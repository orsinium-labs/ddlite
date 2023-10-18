package qb

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconfig"
)

type tChange struct {
	field any
	value any
}

func Set[T any](field *T, value Expr[T]) tChange {
	return tChange{field: field, value: value}
}

type tUpdate[T Model] struct {
	model   *T
	changes []tChange
	conds   []Expr[bool]
}

func Update[T Model](model *T, changes ...tChange) tUpdate[T] {
	return tUpdate[T]{
		model:   model,
		changes: changes,
	}
}

func (u tUpdate[T]) Where(conds ...Expr[bool]) tUpdate[T] {
	u.conds = append(u.conds, conds...)
	return u
}

func (u tUpdate[T]) Squirrel(conf dbconfig.Config) (squirrel.Sqlizer, error) {
	conf = conf.WithModel(u.model)
	// make builder, set column names and table name
	q := squirrel.Update(getModelName(u.model))
	q = q.PlaceholderFormat(squirrel.Dollar)

	// generate SET clause
	for _, change := range u.changes {
		fname, err := getFieldName(u.model, change.field)
		if err != nil {
			return nil, fmt.Errorf("get field name: %v", err)
		}
		q = q.Set(fname, change.value)
	}

	// generate WHERE clause
	if len(u.conds) != 0 {
		preds := make([]squirrel.Sqlizer, 0, len(u.conds))
		for _, cond := range u.conds {
			preds = append(preds, cond.Squirrel(conf))
		}
		q = q.Where(squirrel.And(preds))
	}

	return q, nil
}
