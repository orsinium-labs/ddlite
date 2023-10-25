package dml

import (
	"fmt"

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

func (u tUpdate[T]) Tokens(conf dbconf.Config) (tokens.Tokens, error) {
	conf = conf.WithModel(u.model)
	ts := tokens.New(
		tokens.Keyword("UPDATE"),
		internal.GetTableName(conf, u.model),
		tokens.Keyword("SET"),
	)

	// generate SET clause
	for i, change := range u.changes {
		colName, err := internal.GetColumnName(conf, change.field)
		if err != nil {
			return tokens.New(), fmt.Errorf("get field name: %v", err)
		}
		if i > 0 {
			ts.Add(tokens.Comma())
		}
		ts.Add(
			tokens.ColumnName(colName),
			tokens.Operator("="),
			tokens.Bind(change.value),
		)
	}

	// generate WHERE clause
	if len(u.conds) != 0 {
		ts.Add(tokens.Keyword("WHERE"))
		for i, pred := range u.conds {
			if i > 0 {
				ts.Add(tokens.Keyword("AND"))
			}
			ts.Extend(pred.Tokens(conf))
		}
	}

	return ts, nil
}
