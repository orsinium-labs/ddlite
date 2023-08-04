package qb

import (
	"github.com/Masterminds/squirrel"
)

// tCmpOp is a private type representing binary comparison operations.
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

// E (=) checks if the given column value is equal to the given fixed value.
//
// This is an alias for:
//
//	qb.Eq(qb.C(&column), qb.V(value))
//
// Example:
//
//	qb.E(&u.age, 18) // SQL: age = 18
func E[T comparable](column *T, value T) Expr[bool] {
	return Eq(C(column), V(value))
}

// Gt (>) checks that the left expression is greater than the right one.
//
// Example:
//
//	qb.Gt(qb.C(&u.age), qb.V(18)) // SQL: age > 18
func Gt[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: ">", right: right}
}

// Gte (>=) checks that the left expression is greater than or equal to the right one.
//
// Example:
//
//	qb.Gte(qb.C(&u.age), qb.V(18)) // SQL: age >= 18
func Gte[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: ">=", right: right}
}

// Lt (<) checks that the left expression is less than the right one.
//
// Example:
//
//	qb.Lt(qb.C(&u.age), qb.V(18)) // SQL: age < 18
func Lt[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "<", right: right}
}

// Lte (<=) checks that the left expression is less than or equal to the right one.
//
// Example:
//
//	qb.Lte(qb.C(&u.age), qb.V(18)) // SQL: age <= 18
func Lte[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "<=", right: right}
}

// Eq (=) checks that the left expression is equal to the right one.
//
// Example:
//
//	qb.Eq(&u.created_at, &u.updated_at) // SQL: created_at = updated_at
func Eq[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "=", right: right}
}

// Neq (<>) checks that the left expression is not equal to the right one.
//
// Example:
//
//	qb.Neq(qb.C(&u.age), qb.V(18)) // SQL: age != 18
func Neq[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "<>", right: right}
}

// Is equal to value or both are nulls (missing data) (`IS NOT DISTINCT FROM`)
func IsNotDistinctFrom[T any](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "IS NOT DISTINCT FROM", right: right}
}

// And checks that both left and right expressions are true.
//
// Example:
//
//	qb.And(qb.C(&u.is_admin), qb.E(&u.name, "admin"))
//	// SQL: is_admin AND name = "admin"
func And(left, right Expr[bool]) Expr[bool] {
	return tCmpOp[bool]{left: left, op: "AND", right: right}
}

// Or checks that left, right, or both expressions are true.
//
// Example:
//
//	qb.Or(qb.C(&u.is_admin), qb.E(&u.name, "admin"))
//	// SQL: is_admin OR name = "admin"
func Or(left, right Expr[bool]) Expr[bool] {
	return tCmpOp[bool]{left: left, op: "OR", right: right}
}
