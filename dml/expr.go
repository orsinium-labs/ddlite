package dml

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

// Expr is an SQL expression. I can be used as part of SQL queries.
type Expr[T any] interface {
	ExprType() T
	Precedence(dbconf.Config) uint8
	Tokens(dbconf.Config) tokens.Tokens
}

type exprOperator[T, R any] struct {
	prefix   bool
	operator string
	wrapper  func(string) tokens.Token
	left     Expr[T]
	right    Expr[T]
}

func (expr exprOperator[T, R]) ExprType() R {
	return *new(R)
}

func (expr exprOperator[T, R]) Precedence(c dbconf.Config) uint8 {
	return c.Dialect.Precedence(expr.operator)
}

func (expr exprOperator[T, R]) Tokens(c dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	precSelf := expr.Precedence(c)

	// write prefix
	if expr.prefix {
		ts.Add(expr.wrapper(expr.operator))
	}

	// write left expression
	precLeft := expr.left.Precedence(c)
	paren := precLeft == 0 || precLeft < precSelf
	if paren {
		ts.Add(tokens.LParen())
	}
	ts.Extend(expr.left.Tokens(c))
	if paren {
		ts.Add(tokens.RParen())
	}

	// write infix (or suffix if there is no right)
	if !expr.prefix {
		ts.Add(expr.wrapper(expr.operator))
	}

	// write right (if provided)
	if expr.right != nil {
		precRight := expr.right.Precedence(c)
		paren := precRight <= precSelf
		if paren {
			ts.Add(tokens.LParen())
		}
		ts.Extend(expr.right.Tokens(c))
		if paren {
			ts.Add(tokens.RParen())
		}
	}
	return ts
}
