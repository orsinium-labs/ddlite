package dml

import (
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
	whereClause
	model   *T
	changes []tChange
}

func Update[T internal.Model](model *T, changes ...tChange) tUpdate[T] {
	return tUpdate[T]{
		model:   model,
		changes: changes,
	}
}

func (u tUpdate[T]) Where(predicates ...Expr[bool]) tUpdate[T] {
	u.predicates = append(u.predicates, predicates...)
	return u
}

func (u tUpdate[T]) Tokens(conf dbconf.Config) tokens.Tokens {
	conf = conf.WithModel(u.model)
	ts := tokens.New(
		tokens.Keyword("UPDATE"),
		internal.GetTableName(conf, u.model),
		tokens.Keyword("SET"),
	)

	// generate SET clause
	for i, change := range u.changes {
		if i > 0 {
			ts.Add(tokens.Comma())
		}
		ts.Add(
			internal.GetColumnName(conf, change.field),
			tokens.Operator("="),
			tokens.Bind(change.value),
		)
	}

	// generate WHERE clause
	ts.Extend(u.buildWhere(conf))
	return ts
}
