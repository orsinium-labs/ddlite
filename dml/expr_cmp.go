package dml

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/priority"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

// Expression representing binary a operation returning a bool.
type tCmpOp[T any] struct {
	left  Expr[T]
	op    string
	right Expr[T]
	prio  priority.Priority
}

func (op tCmpOp[T]) Priority() priority.Priority {
	return op.prio
}

func (tCmpOp[T]) ExprType() bool {
	return false
}

func (op tCmpOp[T]) Tokens(c dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	ts.Extend(op.left.Tokens(c))
	ts.Add(tokens.Operator(op.op))
	ts.Extend(op.right.Tokens(c))
	return ts
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
	return tCmpOp[T]{left, ">", right, priority.Comparison}
}

// Gte (>=) checks that the left expression is greater than or equal to the right one.
//
// Example:
//
//	dml.Gte(dml.C(&u.age), dml.V(18)) // SQL: age >= 18
func Gte[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left, ">=", right, priority.Comparison}
}

// Lt (<) checks that the left expression is less than the right one.
//
// Example:
//
//	dml.Lt(dml.C(&u.age), dml.V(18)) // SQL: age < 18
func Lt[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left, "<", right, priority.Comparison}
}

// Lte (<=) checks that the left expression is less than or equal to the right one.
//
// Example:
//
//	dml.Lte(dml.C(&u.age), dml.V(18)) // SQL: age <= 18
func Lte[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left, "<=", right, priority.Comparison}
}

// Eq (=) checks that the left expression is equal to the right one.
//
// Example:
//
//	dml.Eq(&u.created_at, &u.updated_at) // SQL: created_at = updated_at
func Eq[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left, "=", right, priority.Comparison}
}

// Neq (<>) checks that the left expression is not equal to the right one.
//
// Example:
//
//	dml.Neq(dml.C(&u.age), dml.V(18)) // SQL: age != 18
func Neq[T comparable](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left, "<>", right, priority.Comparison}
}

func Like(left, right Expr[string]) Expr[bool] {
	return tCmpOp[string]{left, "LIKE", right, priority.Like}
}

func NotLike(left, right Expr[string]) Expr[bool] {
	return tCmpOp[string]{left, "NOT LIKE", right, priority.Like}
}

func Glob(left, right Expr[string]) Expr[bool] {
	return tCmpOp[string]{left, "GLOB", right, priority.Like}
}

func NotGlob(left, right Expr[string]) Expr[bool] {
	return tCmpOp[string]{left, "NOT GLOB", right, priority.Like}
}

func RegExp(left, right Expr[string]) Expr[bool] {
	return tCmpOp[string]{left, "REGEXP", right, priority.Like}
}

func NotRegExp(left, right Expr[string]) Expr[bool] {
	return tCmpOp[string]{left, "NOT REGEXP", right, priority.Like}
}

func Match(left, right Expr[string]) Expr[bool] {
	return tCmpOp[string]{left, "MATCH", right, priority.Like}
}

func NotMatch(left, right Expr[string]) Expr[bool] {
	return tCmpOp[string]{left, "NOT MATCH", right, priority.Like}
}

func IsDistinctFrom[T any](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left, "IS DISTINCT FROM", right, priority.Is}
}

// Is equal to value or both are nulls (missing data) (`IS NOT DISTINCT FROM`)
func IsNotDistinctFrom[T any](left, right Expr[T]) Expr[bool] {
	return tCmpOp[T]{left, "IS NOT DISTINCT FROM", right, priority.Is}
}
