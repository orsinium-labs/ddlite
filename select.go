package sequel

import (
	"fmt"
	"strings"
)

type selectQ[T any] struct {
	fields []any
	conds  []condition
	model  any
}

func Select[T any](model T, fields ...any) selectQ[T] {
	return selectQ[T]{model: model, fields: fields}
}

func (s selectQ[T]) Where(conds ...condition) selectQ[T] {
	s.conds = append(s.conds, conds...)
	return s
}

func (s selectQ[T]) String() string {
	fnames := make([]string, 0, len(s.fields))
	for _, f := range s.fields {
		fname, err := getFieldName(s.model, f)
		if err != nil {
			panic(err)
		}
		fnames = append(fnames, fname)
	}
	joined := strings.Join(fnames, ", ")
	table := getModelName(s.model)
	q := fmt.Sprintf("SELECT %s FROM %s", joined, table)

	if len(s.conds) != 0 {
		conds := make([]string, 0, len(s.conds))
		for _, c := range s.conds {
			conds = append(conds, c.toSQL(s.model))
		}
		joined = strings.Join(conds, " AND ")
		q = fmt.Sprintf("%s WHERE %s", q, joined)
	}

	return q
}
