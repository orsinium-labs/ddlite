package dml

import (
	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconf"
)

// tCmpOp is a private type representing binary comparison operations.
type tCmpOp[T any] struct {
	left  Expr[T]
	op    string
	right Expr[T]
}

func (tCmpOp[T]) ExprType() bool {
	return false
}

func (op tCmpOp[T]) Squirrel(c dbconf.Config) squirrel.Sqlizer {
	lhs := op.left.Squirrel(c)
	rhs := op.right.Squirrel(c)
	return squirrel.ConcatExpr(lhs, " ", op.op, " ", rhs)
}

// E (=) checks if the given column value is equal to the given fixed value.
//
// This is an alias for:
//
//	dml.Eq(dml.C(&column), dml.V(value))
//
// Example:
//
//	dml.E(&u.age, 18) // SQL: age = 18
func E[T comparable](column *T, value T) Expr[bool] {
	return Eq(C(column), V(value))
}

// Gt (>) checks that the left expression is greater than the right one.
//
// Example:
//
//	dml.Gt(dml.C(&u.age), dml.V(18)) // SQL: age > 18
func Gt[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: ">", right: right}
}

// Gte (>=) checks that the left expression is greater than or equal to the right one.
//
// Example:
//
//	dml.Gte(dml.C(&u.age), dml.V(18)) // SQL: age >= 18
func Gte[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: ">=", right: right}
}

// Lt (<) checks that the left expression is less than the right one.
//
// Example:
//
//	dml.Lt(dml.C(&u.age), dml.V(18)) // SQL: age < 18
func Lt[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "<", right: right}
}

// Lte (<=) checks that the left expression is less than or equal to the right one.
//
// Example:
//
//	dml.Lte(dml.C(&u.age), dml.V(18)) // SQL: age <= 18
func Lte[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "<=", right: right}
}

// Eq (=) checks that the left expression is equal to the right one.
//
// Example:
//
//	dml.Eq(&u.created_at, &u.updated_at) // SQL: created_at = updated_at
func Eq[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left: left, op: "=", right: right}
}

// Neq (<>) checks that the left expression is not equal to the right one.
//
// Example:
//
//	dml.Neq(dml.C(&u.age), dml.V(18)) // SQL: age != 18
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
//	dml.And(dml.C(&u.is_admin), dml.E(&u.name, "admin"))
//	// SQL: is_admin AND name = "admin"
func And(left, right Expr[bool]) Expr[bool] {
	return tCmpOp[bool]{left: left, op: "AND", right: right}
}

// Or checks that left, right, or both expressions are true.
//
// Example:
//
//	dml.Or(dml.C(&u.is_admin), dml.E(&u.name, "admin"))
//	// SQL: is_admin OR name = "admin"
func Or(left, right Expr[bool]) Expr[bool] {
	return tCmpOp[bool]{left: left, op: "OR", right: right}
}
