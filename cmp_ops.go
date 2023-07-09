package sequel

import (
	"github.com/Masterminds/squirrel"
)

// tCmpOp is aprivate type to represent binary comparison operations.
type tCmpOp[T any] struct {
	left  Expr[T]
	op    string
	right Expr[T]
}

func (tCmpOp[T]) Default() bool {
	return false
}

func (op tCmpOp[T]) Squirrel(ms Models) squirrel.Sqlizer {
	lhs := op.left.Squirrel(ms)
	rhs := op.right.Squirrel(ms)
	return squirrel.ConcatExpr(lhs, " ", op.op, " ", rhs)
}

// Greater than (`>`)
func Gt[T any](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: ">", right: right}
}

// Greater than or equal (`>=`)
func Gte[T any](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: ">=", right: right}
}

// Less than (`<`)
func Lt[T any](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "<", right: right}
}

// Less than or equal (`<=`)
func Lte[T any](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "<=", right: right}
}

// Equal to (`=`)
func Eq[T any](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "=", right: right}
}

// Not equal to (`<>`)
func Ne[T any](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "<>", right: right}
}

// Is equal to value or both are nulls (missing data) (`IS NOT DISTINCT FROM`)
func IsNotDistinctFrom[T any](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "IS NOT DISTINCT FROM", right: right}
}
