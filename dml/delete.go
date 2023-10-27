package dml

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tDelete[T internal.Model] struct {
	whereClause
	model *T
}

func Delete[T internal.Model](model *T) tDelete[T] {
	return tDelete[T]{model: model}
}

func (d tDelete[T]) Where(predicates ...Expr[bool]) tDelete[T] {
	d.predicates = append(d.predicates, predicates...)
	return d
}

// And is an alias for Where.
func (d tDelete[T]) And(conds ...Expr[bool]) tDelete[T] {
	return d.Where(conds...)
}

func (d tDelete[T]) Tokens(conf dbconf.Config) tokens.Tokens {
	conf = conf.WithModel(d.model)
	ts := tokens.New(
		tokens.Keyword("DELETE FROM"),
		internal.GetTableName(conf, d.model),
	)
	ts.Extend(d.buildWhere(conf))
	return ts
}
