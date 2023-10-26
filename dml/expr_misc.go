package dml

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/priority"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tIsNull[T, R any] struct {
	left   Expr[T]
	suffix string
}

func (tIsNull[T, R]) ExprType() R {
	return *new(R)
}

func (tIsNull[T, R]) Priority() priority.Priority {
	return priority.Is
}

func (expr tIsNull[T, R]) Tokens(c dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	ts.Extend(expr.left.Tokens(c))
	ts.Add(tokens.Keyword(expr.suffix))
	return ts
}

type tNot[T, R any] struct {
	prefix string
	right  Expr[T]
}

func (tNot[T, R]) Priority() priority.Priority {
	return priority.Not
}

func (tNot[T, R]) ExprType() R {
	return *new(R)
}

func (expr tNot[T, R]) Tokens(c dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	ts.Add(tokens.Keyword(expr.prefix))
	ts.Extend(expr.right.Tokens(c))
	return ts
}

type tChain struct {
	items []Expr[bool]
	infix string
	prio  priority.Priority
}

func (tChain) ExprType() bool {
	return false
}

func (expr tChain) Priority() priority.Priority {
	return expr.prio
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
	return tIsNull[T, bool]{val, "IS NULL"}
}

func IsNotNull[T any](val Expr[T]) Expr[bool] {
	return tIsNull[T, bool]{val, "IS NOT NULL"}
}

func Not(val Expr[bool]) Expr[bool] {
	return tNot[bool, bool]{"NOT", val}
}

// And checks that all given expressions are true.
//
// Example:
//
//	dml.And(dml.C(&u.is_admin), dml.E(&u.name, "admin"))
//	// SQL: is_admin AND name = "admin"
func And(items ...Expr[bool]) Expr[bool] {
	return tChain{items, "AND", priority.And}
}

// Or checks any of the expressions is true.
//
// Example:
//
//	dml.Or(dml.C(&u.is_admin), dml.E(&u.name, "admin"))
//	// SQL: is_admin OR name = "admin"
func Or(items ...Expr[bool]) Expr[bool] {
	return tChain{items, "OR", priority.Or}
}
