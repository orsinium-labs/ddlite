package qb

import (
	"github.com/Masterminds/squirrel"
)

type tDelete[T Model] struct {
	model *T
	conds []Expr[bool]
}

func Delete[T Model](model *T) tDelete[T] {
	return tDelete[T]{model: model}
}

func (d tDelete[T]) Where(conds ...Expr[bool]) tDelete[T] {
	d.conds = append(d.conds, conds...)
	return d
}

// And is an alias for Where.
func (d tDelete[T]) And(conds ...Expr[bool]) tDelete[T] {
	return d.Where(conds...)
}

func (s tDelete[T]) Squirrel(...Model) (squirrel.Sqlizer, error) {
	q := squirrel.Delete(getModelName(s.model))
	q = q.PlaceholderFormat(squirrel.Dollar)

	if len(s.conds) != 0 {
		preds := make([]squirrel.Sqlizer, 0, len(s.conds))
		for _, c := range s.conds {
			preds = append(preds, c.Squirrel(s.model))
		}
		q = q.Where(squirrel.And(preds))
	}

	return q, nil
}
