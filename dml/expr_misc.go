package dml

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/priority"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type exprChain struct {
	items []Expr[bool]
	infix string
	prio  priority.Priority
}

func (exprChain) ExprType() bool {
	return false
}

func (expr exprChain) Priority() priority.Priority {
	return expr.prio
}

func (expr exprChain) Tokens(c dbconf.Config) tokens.Tokens {
	switch len(expr.items) {
	case 0:
		err := fmt.Errorf("operator %s requires at least one item", expr.infix)
		return tokens.New(tokens.Err(err))
	case 1:
		return expr.items[0].Tokens(c)
	default:
		ts := tokens.New()
		for i, item := range expr.items {
			if i > 0 {
				ts.Add(tokens.Keyword(expr.infix))
			}
			ts.Extend(item.Tokens(c))
		}
		return ts
	}
}

func IsNull[T any](val Expr[T]) Expr[bool] {
	return exprOperator[T, bool]{
		priority: priority.Is,
		left:     val,
		token:    tokens.Keyword("IS NULL"),
	}
}

func IsNotNull[T any](val Expr[T]) Expr[bool] {
	return exprOperator[T, bool]{
		priority: priority.Is,
		left:     val,
		token:    tokens.Keyword("IS NOT NULL"),
	}
}

func Not(val Expr[bool]) Expr[bool] {
	return exprOperator[bool, bool]{
		priority: priority.Not,
		token:    tokens.Keyword("NOT"),
		left:     val,
		prefix:   true,
	}
}

// And checks that all given expressions are true.
//
// Example:
//
//	dml.And(dml.C(&u.is_admin), dml.E(&u.name, "admin"))
//	// SQL: is_admin AND name = "admin"
func And(items ...Expr[bool]) Expr[bool] {
	return exprChain{items, "AND", priority.And}
}

// Or checks any of the expressions is true.
//
// Example:
//
//	dml.Or(dml.C(&u.is_admin), dml.E(&u.name, "admin"))
//	// SQL: is_admin OR name = "admin"
func Or(items ...Expr[bool]) Expr[bool] {
	return exprChain{items, "OR", priority.Or}
}
