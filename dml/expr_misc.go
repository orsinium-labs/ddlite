package dml

import (
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

func IsNull[T any](val Expr[T]) Expr[bool] {
	return tSuffix[T, bool]{val, "IS NULL"}
}

func IsNotNull[T any](val Expr[T]) Expr[bool] {
	return tSuffix[T, bool]{val, "IS NOT NULL"}
}

func Not(val Expr[bool]) Expr[bool] {
	return tPrefix[bool, bool]{"NOT", val}
}
