package sequel

import (
	"github.com/Masterminds/squirrel"
)

type selectQ[T any] struct {
	fields []any
	conds  []Predicate
	model  any
}

func Select[T any](model T, fields ...any) selectQ[T] {
	return selectQ[T]{model: model, fields: fields}
}

func (s selectQ[T]) Where(conds ...Predicate) selectQ[T] {
	s.conds = append(s.conds, conds...)
	return s
}

func (s selectQ[T]) Squirrel() squirrel.SelectBuilder {
	fnames := make([]string, 0, len(s.fields))
	for _, f := range s.fields {
		fname, err := getFieldName(s.model, f)
		if err != nil {
			panic(err)
		}
		fnames = append(fnames, fname)
	}
	q := squirrel.Select(fnames...)
	q = q.From(getModelName(s.model))

	if len(s.conds) != 0 {
		preds := make([]squirrel.Sqlizer, 0, len(s.conds))
		for _, c := range s.conds {
			preds = append(preds, c.Squirrel(s.model))
		}
		q = q.Where(squirrel.And(preds))
	}

	return q
}

func (s selectQ[T]) String() string {
	sq := s.Squirrel()
	sql, _ := sq.MustSql()
	return sql
}
