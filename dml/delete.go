package dml

import (
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

func (d tDelete[T]) Tokens(conf dbconf.Config) (tokens.Tokens, error) {
	conf = conf.WithModel(d.model)
	ts := tokens.New(
		tokens.Keyword("DELETE FROM"),
		internal.GetTableName(conf, d.model),
	)

	if len(d.conds) != 0 {
		ts.Add(tokens.Keyword("WHERE"))
		for i, pred := range d.conds {
			if i > 0 {
				ts.Add(tokens.Keyword("AND"))
			}
			ts.Extend(pred.Tokens(conf))
		}
	}

	return ts, nil
}
