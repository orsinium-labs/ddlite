package dml

import (
	"github.com/orsinium-labs/sequel/internal/tokens"
)

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
	return exprOperator[T, bool]{
		wrapper:  tokens.Operator,
		operator: ">",
		left:     left,
		right:    right,
	}
}

// Gte (>=) checks that the left expression is greater than or equal to the right one.
//
// Example:
//
//	dml.Gte(dml.C(&u.age), dml.V(18)) // SQL: age >= 18
func Gte[T comparable](left, right Expr[T]) Expr[bool] {
	return exprOperator[T, bool]{
		wrapper:  tokens.Operator,
		operator: ">=",
		left:     left,
		right:    right,
	}
}

// Lt (<) checks that the left expression is less than the right one.
//
// Example:
//
//	dml.Lt(dml.C(&u.age), dml.V(18)) // SQL: age < 18
func Lt[T comparable](left, right Expr[T]) Expr[bool] {
	return exprOperator[T, bool]{
		wrapper:  tokens.Operator,
		operator: "<",
		left:     left,
		right:    right,
	}
}

// Lte (<=) checks that the left expression is less than or equal to the right one.
//
// Example:
//
//	dml.Lte(dml.C(&u.age), dml.V(18)) // SQL: age <= 18
func Lte[T comparable](left, right Expr[T]) Expr[bool] {
	return exprOperator[T, bool]{
		wrapper:  tokens.Operator,
		operator: "<=",
		left:     left,
		right:    right,
	}
}

// Eq (=) checks that the left expression is equal to the right one.
//
// Example:
//
//	dml.Eq(&u.created_at, &u.updated_at) // SQL: created_at = updated_at
func Eq[T comparable](left, right Expr[T]) Expr[bool] {
	return exprOperator[T, bool]{
		wrapper:  tokens.Operator,
		operator: "=",
		left:     left,
		right:    right,
	}
}

// Neq (<>) checks that the left expression is not equal to the right one.
//
// Example:
//
//	dml.Neq(dml.C(&u.age), dml.V(18)) // SQL: age != 18
func Neq[T comparable](left, right Expr[T]) Expr[bool] {
	return exprOperator[T, bool]{
		wrapper:  tokens.Operator,
		operator: "<>",
		left:     left,
		right:    right,
	}
}

func Like(left, right Expr[string]) Expr[bool] {
	return exprOperator[string, bool]{
		wrapper:  tokens.Keyword,
		operator: "LIKE",
		left:     left,
		right:    right,
		//
	}
}

func NotLike(left, right Expr[string]) Expr[bool] {
	return exprOperator[string, bool]{
		wrapper:  tokens.Keyword,
		operator: "NOT LIKE",
		left:     left,
		right:    right,
	}
}

func Glob(left, right Expr[string]) Expr[bool] {
	return exprOperator[string, bool]{
		wrapper:  tokens.Keyword,
		operator: "GLOB",
		left:     left,
		right:    right,
	}
}

func NotGlob(left, right Expr[string]) Expr[bool] {
	return exprOperator[string, bool]{
		wrapper:  tokens.Keyword,
		operator: "NOT GLOB",
		left:     left,
		right:    right,
	}
}

func RegExp(left, right Expr[string]) Expr[bool] {
	return exprOperator[string, bool]{
		wrapper:  tokens.Keyword,
		operator: "REGEXP",
		left:     left,
		right:    right,
	}
}

func NotRegExp(left, right Expr[string]) Expr[bool] {
	return exprOperator[string, bool]{
		wrapper:  tokens.Keyword,
		operator: "NOT REGEXP",
		left:     left,
		right:    right,
	}
}

func Match(left, right Expr[string]) Expr[bool] {
	return exprOperator[string, bool]{
		wrapper:  tokens.Keyword,
		operator: "MATCH",
		left:     left,
		right:    right,
	}
}

func NotMatch(left, right Expr[string]) Expr[bool] {
	return exprOperator[string, bool]{
		wrapper:  tokens.Keyword,
		operator: "NOT MATCH",
		left:     left,
		right:    right,
	}
}

func IsDistinctFrom[T any](left, right Expr[T]) Expr[bool] {
	return exprOperator[T, bool]{
		wrapper:  tokens.Keyword,
		operator: "IS DISTINCT FROM",
		left:     left,
		right:    right,
	}
}

// Is equal to value or both are nulls (missing data) (`IS NOT DISTINCT FROM`)
func IsNotDistinctFrom[T any](left, right Expr[T]) Expr[bool] {
	return exprOperator[T, bool]{
		wrapper:  tokens.Keyword,
		operator: "IS NOT DISTINCT FROM",
		left:     left,
		right:    right,
	}
}
