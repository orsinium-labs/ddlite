package dml

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tDelete[T internal.Model] struct {
	model *T
	conds []Expr[bool]
}

func Delete[T internal.Model](model *T) tDelete[T] {
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

func (d tDelete[T]) Squirrel(conf dbconf.Config) (squirrel.Sqlizer, error) {
	conf = conf.WithModel(d.model)
	q := squirrel.Delete(internal.GetTableName(conf, d.model))
	q = q.PlaceholderFormat(conf.SquirrelPlaceholder())

	if len(d.conds) != 0 {
		preds := tokens.New()
		first := true
		for _, pred := range d.conds {
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
