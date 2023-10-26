package dml

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tSuffix[T, R any] struct {
	left   Expr[T]
	suffix string
}

func (tSuffix[T, R]) ExprType() R {
	return *new(R)
}

func (expr tSuffix[T, R]) Tokens(c dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	ts.Extend(expr.left.Tokens(c))
	ts.Add(tokens.Keyword(expr.suffix))
	return ts
}

type tPrefix[T, R any] struct {
	prefix string
	right  Expr[T]
}

func (tPrefix[T, R]) ExprType() R {
	return *new(R)
}

func (expr tPrefix[T, R]) Tokens(c dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	ts.Add(tokens.Keyword(expr.prefix))
	ts.Extend(expr.right.Tokens(c))
	return ts
}

type tChain struct {
	items []Expr[bool]
	infix string
}

func (tChain) ExprType() bool {
	return false
}

func (expr tChain) Tokens(c dbconf.Config) tokens.Tokens {
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
	return tSuffix[T, bool]{val, "IS NULL"}
}

func IsNotNull[T any](val Expr[T]) Expr[bool] {
	return tSuffix[T, bool]{val, "IS NOT NULL"}
}

func Not(val Expr[bool]) Expr[bool] {
	return tPrefix[bool, bool]{"NOT", val}
}

// And checks that all given expressions are true.
//
// Example:
//
//	dml.And(dml.C(&u.is_admin), dml.E(&u.name, "admin"))
//	// SQL: is_admin AND name = "admin"
func And(items ...Expr[bool]) Expr[bool] {
	return tChain{items, "AND"}
}

// Or checks any of the expressions is true.
//
// Example:
//
//	dml.Or(dml.C(&u.is_admin), dml.E(&u.name, "admin"))
//	// SQL: is_admin OR name = "admin"
func Or(items ...Expr[bool]) Expr[bool] {
	return tChain{items, "OR"}
}
