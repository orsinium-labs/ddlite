package sequel

import (
	"github.com/Masterminds/squirrel"
)

type selectQ[T Model] struct {
	fields []any
	conds  []Expr[bool]
	model  T
}

func Select[T Model](model T, fields ...any) selectQ[T] {
	return selectQ[T]{model: model, fields: fields}
}

func (s selectQ[T]) Where(conds ...Expr[bool]) selectQ[T] {
	s.conds = append(s.conds, conds...)
	return s
}

// And is an alias for Where.
func (s selectQ[T]) And(conds ...Expr[bool]) selectQ[T] {
	return s.Where(conds...)
}

func (s selectQ[T]) Squirrel(...Model) squirrel.SelectBuilder {
	fnames := make([]string, 0, len(s.fields))
	for _, f := range s.fields {
		fname, err := getFieldName(s.model, f)
		if err != nil {
			panic(err)
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

	return q
}

func (s selectQ[T]) String() string {
	sq := s.Squirrel()
	sql, _ := sq.MustSql()
	return sql
}
