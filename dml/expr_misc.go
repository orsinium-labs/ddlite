package dml

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type exprChain struct {
	items []Expr[bool]
	infix string
}

func (exprChain) ExprType() bool {
	return false
}

func (expr exprChain) Precedence(c dbconf.Config) uint8 {
	return c.Dialect.Precedence(expr.infix)
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
		left:     val,
		wrapper:  tokens.Keyword,
		operator: "IS NULL",
	}
}

func IsNotNull[T any](val Expr[T]) Expr[bool] {
	return exprOperator[T, bool]{
		left:     val,
		wrapper:  tokens.Keyword,
		operator: "IS NOT NULL",
	}
}

func Not(val Expr[bool]) Expr[bool] {
	return exprOperator[bool, bool]{
		wrapper:  tokens.Keyword,
		operator: "NOT",
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
	return exprChain{items, "AND"}
}

// Or checks any of the expressions is true.
//
// Example:
//
//	dml.Or(dml.C(&u.is_admin), dml.E(&u.name, "admin"))
//	// SQL: is_admin OR name = "admin"
func Or(items ...Expr[bool]) Expr[bool] {
	return exprChain{items, "OR"}
}
