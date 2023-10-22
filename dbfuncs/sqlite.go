package dbfuncs

import (
	c "github.com/orsinium-labs/sequel/constraints"
	"github.com/orsinium-labs/sequel/dml"
)

// Abs function returns the absolute value of a number.
func Abs[N c.Number](val dml.Expr[N]) dml.Expr[N] {
	return dml.F[N, N]("abs", val)
}

// Ceil function returns the smallest integer value that is greater than or equal to a number.
func Ceil[I c.Integer, F c.Float](val dml.Expr[F]) dml.Expr[I] {
	return dml.F[F, I]("ceil", val)
}

// Div function is used for integer division where m is divided by n and an integer value is returned.
func Div[I c.Integer, N c.Number](m, n dml.Expr[N]) dml.Expr[I] {
	return dml.F[N, I]("div", m, n)
}

// Exp function returns e raised to the power of number.
func Exp[F c.Float, I c.Integer](val dml.Expr[I]) dml.Expr[F] {
	return dml.F[I, F]("exp", val)
}

// Floor function returns the largest integer value that is equal to or less than a number.
func Floor[I c.Integer, F c.Float](val dml.Expr[F]) dml.Expr[I] {
	return dml.F[F, I]("floor", val)
}

// Mod function returns the remainder of m divided by n.
func Mod[I c.Integer, N c.Number](m, n dml.Expr[N]) dml.Expr[I] {
	return dml.F[N, I]("mod", m, n)
}

// Power function returns m raised to the nth power.
func Power[F c.Float, N c.Number, I c.Integer](m dml.Expr[N], n dml.Expr[I]) dml.Expr[F] {
	return dml.F2[N, I, F]("power", m, n)
}

// Random function can be used to return a random number within 0-1 range.
func Random[F c.Float]() dml.Expr[F] {
	return dml.F[F, F]("random")
}

// Round function returns a rounded number.
func Round[I c.Integer, F c.Float](val dml.Expr[F]) dml.Expr[I] {
	return dml.F[F, I]("round", val)
}

// RoundTo function returns a number rounded to a certain number of decimal places.
func RoundTo[F c.Float, I c.Integer](val dml.Expr[F], dp dml.Expr[I]) dml.Expr[F] {
	return dml.F2[F, I, F]("round", val, dp)
}

// Sign function returns a value indicating the sign of a number.
func Sign[I c.Integer, F c.Float](val dml.Expr[F]) dml.Expr[I] {
	return dml.F[F, I]("sign", val)
}

// Sqrt function returns the square root of a number.
func Sqrt[F c.Float, N c.Number](val dml.Expr[N]) dml.Expr[F] {
	return dml.F[N, F]("sqrt", val)
}

// Trunc function returns a truncated number.
func Trunk[I c.Integer, F c.Float](val dml.Expr[F]) dml.Expr[I] {
	return dml.F[F, I]("trunk", val)
}

// TruncTo function returns a number truncated to a certain number of decimal places.
func TrunkTo[F c.Float, I c.Integer](val dml.Expr[F], dp dml.Expr[I]) dml.Expr[F] {
	return dml.F2[F, I, F]("trunk", val, dp)
}
