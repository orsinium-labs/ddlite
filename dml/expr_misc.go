package dml

import (
	"github.com/orsinium-labs/sequel/internal/tokens"
)

func IsNull[T any](val Expr[T]) Expr[bool] {
	return exprOperator[T, bool]{
		left:     val,
		wrapper:  tokens.Keyword,
		operator: "IS NULL",
	}
}

func IsNotNull[T any](val Expr[T]) Expr[bool] {
	return exprOperator[T, bool]{
		left:     val,
		wrapper:  tokens.Keyword,
		operator: "IS NOT NULL",
	}
}

func Not(val Expr[bool]) Expr[bool] {
	return exprOperator[bool, bool]{
		wrapper:  tokens.Keyword,
		operator: "NOT",
		left:     val,
		prefix:   true,
	}
}

// And checks that both given expressions are true.
//
// Example:
//
//	dml.And(dml.C(&u.is_admin), dml.E(&u.name, "admin"))
//	// SQL: is_admin AND name = "admin"
func And(left, right Expr[bool]) Expr[bool] {
	return exprOperator[bool, bool]{
		wrapper:  tokens.Keyword,
		operator: "AND",
		left:     left,
		right:    right,
	}
}

// Or checks that any of the given expressions is true.
//
// Example:
//
//	dml.Or(dml.C(&u.is_admin), dml.E(&u.name, "admin"))
//	// SQL: is_admin OR name = "admin"
func Or(left, right Expr[bool]) Expr[bool] {
	return exprOperator[bool, bool]{
		wrapper:  tokens.Keyword,
		operator: "OR",
		left:     left,
		right:    right,
	}
}
