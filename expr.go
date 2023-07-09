package sequel

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

type Models []any

// Expr is an SQL expression. I can be used as part of SQL queries.
type Expr[T any] interface {
	Default() T
	Squirrel(Models) squirrel.Sqlizer
}

// tFunc is a private type to represent stored function expression.
// `T` is the type of the function return value.
type tFunc[T any] struct {
	Name string
	Args []any
}

// F is a stored function.
func F[T any](name string, args ...any) Expr[T] {
	return tFunc[T]{Name: name, Args: args}
}

func (tFunc[T]) Default() T {
	return *new(T)
}

func (fn tFunc[T]) Squirrel(ms Models) squirrel.Sqlizer {
	return squirrel.Expr(fmt.Sprintf("%s(?)", fn.Name), fn.Args...)
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

func (col tCol[T]) Squirrel(ms Models) squirrel.Sqlizer {
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

func (val tVal[T]) Squirrel(ms Models) squirrel.Sqlizer {
	return squirrel.Expr("?", val.val)
}
