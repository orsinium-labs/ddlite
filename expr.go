package qb

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
)

// Expr is an SQL expression. I can be used as part of SQL queries.
type Expr[T any] interface {
	Default() T
	Squirrel(...Model) squirrel.Sqlizer
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

func (tFunc[A, R]) Default() R {
	return *new(R)
}

func (fn tFunc[A, R]) Squirrel(ms ...Model) squirrel.Sqlizer {
	args := make([]any, 0, len(fn.Args))
	for _, arg := range fn.Args {
		args = append(args, arg.Squirrel(ms...))
	}
	phs := strings.Repeat("?, ", len(args))
	phs = phs[:len(phs)-2]
	return squirrel.Expr(fmt.Sprintf("%s(%s)", fn.Name, phs), args...)
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

func (tFunc2[A1, A2, R]) Default() R {
	return *new(R)
}

func (fn tFunc2[A1, A2, R]) Squirrel(ms ...Model) squirrel.Sqlizer {
	arg1 := fn.Arg1.Squirrel(ms...)
	arg2 := fn.Arg2.Squirrel(ms...)
	return squirrel.Expr(fmt.Sprintf("%s(?, ?)", fn.Name), arg1, arg2)
}

// tCol is aprivate type to represent a column name expression.
type tCol[T any] struct {
	val *T
}

// C is a column.
func C[T any](val *T) Expr[T] {
	return tCol[T]{val: val}
}

func (tCol[T]) Default() T {
	return *new(T)
}

func (col tCol[T]) Squirrel(ms ...Model) squirrel.Sqlizer {
	for _, model := range ms {
		fname, err := getFieldName(model, col.val)
		if err == nil {
			return squirrel.Expr(fname)
		}
	}
	panic("uknown column")
}

// tVal is a private type to represent a literal value expression.
type tVal[T any] struct {
	val T
}

// V is a literal value.
func V[T any](val T) Expr[T] {
	return tVal[T]{val: val}
}

func (tVal[T]) Default() T {
	return *new(T)
}

func (val tVal[T]) Squirrel(ms ...Model) squirrel.Sqlizer {
	return squirrel.Expr("?", val.val)
}
