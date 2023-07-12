package pgext

import (
	sq "github.com/orsinium-labs/sequel"
	c "github.com/orsinium-labs/sequel/constraints"
)

// Abs function returns the absolute value of a number.
func Abs[N c.Number](val sq.Expr[N]) sq.Expr[N] {
	return sq.F[N, N]("abs", val)
}

// Ceil function returns the smallest integer value that is greater than or equal to a number.
func Ceil[F c.Float, I c.Integer](val sq.Expr[F]) sq.Expr[I] {
	return sq.F[F, I]("ceil", val)
}

// Ceiling function returns the smallest integer value that is greater than or equal to a number.
func Ceiling[F c.Float](val sq.Expr[F]) sq.Expr[F] {
	return sq.F[F, F]("ceiling", val)
}

// Div function is used for integer division where m is divided by n and an integer value is returned.
func Div[N c.Number, I c.Integer](m, n sq.Expr[N]) sq.Expr[I] {
	return sq.F[N, I]("div", m, n)
}

// Exp function returns e raised to the power of number.
func Exp[I c.Integer, F c.Float](val sq.Expr[I]) sq.Expr[F] {
	return sq.F[I, F]("exp", val)
}

// Floor function returns the largest integer value that is equal to or less than a number.
func Floor[F c.Float, I c.Integer](val sq.Expr[F]) sq.Expr[I] {
	return sq.F[F, I]("floor", val)
}

// Mod function returns the remainder of m divided by n.
func Mod[N c.Number, I c.Integer](m, n sq.Expr[N]) sq.Expr[I] {
	return sq.F[N, I]("mod", m, n)
}

// Power function returns m raised to the nth power.
func Power[N c.Number, I c.Integer, F c.Float](m sq.Expr[N], n sq.Expr[I]) sq.Expr[F] {
	return sq.F2[N, I, F]("power", m, n)
}

// Random function can be used to return a random number within 0-1 range.
func Random[F c.Float]() sq.Expr[F] {
	return sq.F[F, F]("random")
}

// Round function returns a rounded number.
func Round[F c.Float, I c.Integer](val sq.Expr[F]) sq.Expr[I] {
	return sq.F[F, I]("round", val)
}

// RoundTo function returns a number rounded to a certain number of decimal places.
func RoundTo[F c.Float, I c.Integer](val sq.Expr[F], dp sq.Expr[I]) sq.Expr[F] {
	return sq.F2[F, I, F]("round", val, dp)
}

// Sign function returns a value indicating the sign of a number.
func Sign[F c.Float, I c.Integer](val sq.Expr[F]) sq.Expr[I] {
	return sq.F[F, I]("sign", val)
}

// Sqrt function returns the square root of a number.
func Sqrt[N c.Number, F c.Float](val sq.Expr[N]) sq.Expr[F] {
	return sq.F[N, F]("sqrt", val)
}

// Trunc function returns a truncated number.
func Trunk[F c.Float, I c.Integer](val sq.Expr[F]) sq.Expr[I] {
	return sq.F[F, I]("trunk", val)
}

// TruncTo function returns a number truncated to a certain number of decimal places.
func TrunkTo[F c.Float, I c.Integer](val sq.Expr[F], dp sq.Expr[I]) sq.Expr[F] {
	return sq.F2[F, I, F]("trunk", val, dp)
}
