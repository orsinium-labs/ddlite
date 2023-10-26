package dml

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/priority"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

// Expr is an SQL expression. I can be used as part of SQL queries.
type Expr[T any] interface {
	ExprType() T
	Priority() priority.Priority
	Tokens(dbconf.Config) tokens.Tokens
}

type exprOperator[T, R any] struct {
	priority priority.Priority
	prefix   bool
	token    tokens.Token
	left     Expr[T]
	right    Expr[T]
}

func (expr exprOperator[T, R]) ExprType() R {
	return *new(R)
}

func (expr exprOperator[T, R]) Priority() priority.Priority {
	return expr.priority
}

func (expr exprOperator[T, R]) Tokens(c dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	if expr.prefix {
		ts.Add(expr.token)
	}
	ts.Extend(expr.left.Tokens(c))
	if !expr.prefix {
		ts.Add(expr.token)
	}
	if expr.right != nil {
		ts.Extend(expr.right.Tokens(c))
	}
	return ts
}
