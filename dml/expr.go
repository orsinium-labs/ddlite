package dml

import (
	"github.com/orsinium-labs/sequel/constraints"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

// Expr is an SQL expression. I can be used as part of SQL queries.
type Expr[T any] interface {
	ExprType() T
	Tokens(dbconf.Config) tokens.Tokens
}

// tFunc is a private type to represent stored function expression.
// `R` is the type of the function return value.
type tFunc[A, R any] struct {
	Name string
	Args []Expr[A]
}

// F is a stored function.
func F[A, T any](name string, args ...Expr[A]) Expr[T] {
	return tFunc[A, T]{Name: name, Args: args}
}

func (tFunc[A, R]) ExprType() R {
	return *new(R)
}

func (fn tFunc[A, R]) Tokens(conf dbconf.Config) tokens.Tokens {
	ts := tokens.New(
		tokens.Raw(fn.Name),
		tokens.LParen(),
	)
	first := true
	for _, arg := range fn.Args {
		if first {
			first = false
		} else {
			ts.Add(tokens.Comma())
		}
		ts.Extend(arg.Tokens(conf))
	}
	ts.Add(tokens.RParen())
	return ts
}

// tFunc is a private type to represent 2-argument stored function expression.
// `R` is the type of the function return value.
type tFunc2[A1, A2, R any] struct {
	Name string
	Arg1 Expr[A1]
	Arg2 Expr[A2]
}

// F is a stored function with 2 arguments of different type.
//
// For functions with any number of arguments of the same type
// prefer using `F` instead.
func F2[A1, A2, T any](name string, arg1 Expr[A1], arg2 Expr[A2]) Expr[T] {
	return tFunc2[A1, A2, T]{Name: name, Arg1: arg1, Arg2: arg2}
}

func (tFunc2[A1, A2, R]) ExprType() R {
	return *new(R)
}

func (fn tFunc2[A1, A2, R]) Tokens(conf dbconf.Config) tokens.Tokens {
	ts := tokens.New(
		tokens.Raw(fn.Name),
		tokens.LParen(),
	)
	ts.Extend(fn.Arg1.Tokens(conf))
	ts.Add(tokens.Comma())
	ts.Extend(fn.Arg2.Tokens(conf))
	ts.Add(tokens.RParen())
	return ts
}

// tCol is aprivate type to represent a column name expression.
type tCol[T any] struct {
	val any
}

// C is a column.
func C[T any](val *T) Expr[T] {
	return tCol[T]{val: val}
}

// M is a column wrapped into Option/Optional/Maybe monad.
func M[T any](val *constraints.Option[T]) Expr[T] {
	return tCol[T]{val: val}
}

func (tCol[T]) ExprType() T {
	return *new(T)
}

func (col tCol[T]) Tokens(conf dbconf.Config) tokens.Tokens {
	colName, err := internal.GetColumnName(conf, col.val)
	if err != nil {
		panic("uknown column")
	}
	return tokens.New(tokens.ColumnName(colName))
}

// tVal is a private type to represent a literal value expression.
type tVal[T any] struct {
	val T
}

// V is a literal value.
func V[T any](val T) Expr[T] {
	return tVal[T]{val: val}
}

func (tVal[T]) ExprType() T {
	return *new(T)
}

func (val tVal[T]) Tokens(dbconf.Config) tokens.Tokens {
	return tokens.New(tokens.Bind(val.val))
}
