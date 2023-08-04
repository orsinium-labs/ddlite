package pgext

import (
	c "github.com/orsinium-labs/sequel/constraints"
	"github.com/orsinium-labs/sequel/qb"
)

// Abs function returns the absolute value of a number.
func Abs[N c.Number](val qb.Expr[N]) qb.Expr[N] {
	return qb.F[N, N]("abs", val)
}

// Ceil function returns the smallest integer value that is greater than or equal to a number.
func Ceil[F c.Float, I c.Integer](val qb.Expr[F]) qb.Expr[I] {
	return qb.F[F, I]("ceil", val)
}

// Div function is used for integer division where m is divided by n and an integer value is returned.
func Div[N c.Number, I c.Integer](m, n qb.Expr[N]) qb.Expr[I] {
	return qb.F[N, I]("div", m, n)
}

// Exp function returns e raised to the power of number.
func Exp[I c.Integer, F c.Float](val qb.Expr[I]) qb.Expr[F] {
	return qb.F[I, F]("exp", val)
}

// Floor function returns the largest integer value that is equal to or less than a number.
func Floor[F c.Float, I c.Integer](val qb.Expr[F]) qb.Expr[I] {
	return qb.F[F, I]("floor", val)
}

// Mod function returns the remainder of m divided by n.
func Mod[N c.Number, I c.Integer](m, n qb.Expr[N]) qb.Expr[I] {
	return qb.F[N, I]("mod", m, n)
}

// Power function returns m raised to the nth power.
func Power[N c.Number, I c.Integer, F c.Float](m qb.Expr[N], n qb.Expr[I]) qb.Expr[F] {
	return qb.F2[N, I, F]("power", m, n)
}

// Random function can be used to return a random number within 0-1 range.
func Random[F c.Float]() qb.Expr[F] {
	return qb.F[F, F]("random")
}

// Round function returns a rounded number.
func Round[F c.Float, I c.Integer](val qb.Expr[F]) qb.Expr[I] {
	return qb.F[F, I]("round", val)
}

// RoundTo function returns a number rounded to a certain number of decimal places.
func RoundTo[F c.Float, I c.Integer](val qb.Expr[F], dp qb.Expr[I]) qb.Expr[F] {
	return qb.F2[F, I, F]("round", val, dp)
}

// Sign function returns a value indicating the sign of a number.
func Sign[F c.Float, I c.Integer](val qb.Expr[F]) qb.Expr[I] {
	return qb.F[F, I]("sign", val)
}

// Sqrt function returns the square root of a number.
func Sqrt[N c.Number, F c.Float](val qb.Expr[N]) qb.Expr[F] {
	return qb.F[N, F]("sqrt", val)
}

// Trunc function returns a truncated number.
func Trunk[F c.Float, I c.Integer](val qb.Expr[F]) qb.Expr[I] {
	return qb.F[F, I]("trunk", val)
}

// TruncTo function returns a number truncated to a certain number of decimal places.
func TrunkTo[F c.Float, I c.Integer](val qb.Expr[F], dp qb.Expr[I]) qb.Expr[F] {
	return qb.F2[F, I, F]("trunk", val, dp)
}
