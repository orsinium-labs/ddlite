package qb

import (
	"github.com/Masterminds/squirrel"
)

// tCmpOp is a private type to represent binary comparison operations.
type tCmpOp[T any] struct {
	left  Expr[T]
	op    string
	right Expr[T]
}

func (tCmpOp[T]) Default() bool {
	return false
}

func (op tCmpOp[T]) Squirrel(ms ...Model) squirrel.Sqlizer {
	lhs := op.left.Squirrel(ms...)
	rhs := op.right.Squirrel(ms...)
	return squirrel.ConcatExpr(lhs, " ", op.op, " ", rhs)
}

// E checks if the given column value is equal to the given fixed value.
//
// This is an alias for:
//
//	qb.Eq(qb.C(&column), qb.V(value))
func E[T any](column *T, value T) Expr[bool] {
	return Eq(C(column), V(value))
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
func Neq[T any](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "<>", right: right}
}

// Is equal to value or both are nulls (missing data) (`IS NOT DISTINCT FROM`)
func IsNotDistinctFrom[T any](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "IS NOT DISTINCT FROM", right: right}
}

// And checks that both left and right expressions are true.
func And(left, right Expr[bool]) Expr[bool] {
	return tCmpOp[bool]{left: left, op: "AND", right: right}
}

// Or checks that left, right, or both expressions are true.
func Or(left, right Expr[bool]) Expr[bool] {
	return tCmpOp[bool]{left: left, op: "OR", right: right}
}
