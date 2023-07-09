package sequel

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

type Models []any

type Expr[T any] interface {
	Default() T
	Squirrel(Models) squirrel.Sqlizer
}

type tFunc[T any] struct {
	Name string
	Args []any
}

func F[T any](name string, args ...any) Expr[T] {
	return tFunc[T]{Name: name, Args: args}
}

func (tFunc[T]) Default() T {
	return *new(T)
}

func (fn tFunc[T]) Squirrel(ms Models) squirrel.Sqlizer {
	return squirrel.Expr(fmt.Sprintf("%s(?)", fn.Name), fn.Args...)
}

type tCol[T any] struct {
	val *T
}

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

type tVal[T any] struct {
	val T
}

func V[T any](val T) Expr[T] {
	return tVal[T]{val: val}
}

func (tVal[T]) Default() T {
	return *new(T)
}

func (val tVal[T]) Squirrel(ms Models) squirrel.Sqlizer {
	return squirrel.Expr("?", val.val)
}
