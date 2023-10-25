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

func IsNull[T any](val Expr[T]) Expr[bool] {
	return tSuffix[T, bool]{val, "ISNULL"}
}

func IsNotNull[T any](val Expr[T]) Expr[bool] {
	return tSuffix[T, bool]{val, "ISNOTNULL"}
}
