package pgext

import (
	sq "github.com/orsinium-labs/sequel"
	c "github.com/orsinium-labs/sequel/constraints"
)

func Abs[N c.Number](val sq.Expr[N]) sq.Expr[N] {
	return sq.F[N, N]("abs", val)
}

func Ceil[F c.Float, I c.Integer](val sq.Expr[F]) sq.Expr[I] {
	return sq.F[F, I]("ceil", val)
}

func Ceiling[F c.Float](val sq.Expr[F]) sq.Expr[F] {
	return sq.F[F, F]("ceiling", val)
}

func Div[N c.Number, I c.Integer](a, b sq.Expr[N]) sq.Expr[I] {
	return sq.F[N, I]("div", a, b)
}

func Exp[I c.Integer, F c.Float](val sq.Expr[I]) sq.Expr[F] {
	return sq.F[I, F]("exp", val)
}

func Floor[F c.Float, I c.Integer](val sq.Expr[F]) sq.Expr[I] {
	return sq.F[F, I]("floor", val)
}

func Mod[N c.Number, I c.Integer](a, b sq.Expr[N]) sq.Expr[I] {
	return sq.F[N, I]("mod", a, b)
}

func Power[N c.Number, I c.Integer, F c.Float](a sq.Expr[N], b sq.Expr[I]) sq.Expr[F] {
	return sq.F2[N, I, F]("power", a, b)
}

func Random[F c.Float]() sq.Expr[F] {
	return sq.F[F, F]("random")
}

func Round[F c.Float, I c.Integer](val sq.Expr[F]) sq.Expr[I] {
	return sq.F[F, I]("round", val)
}

func RoundTo[F c.Float, I1, I2 c.Integer](val sq.Expr[F], places sq.Expr[I1]) sq.Expr[I2] {
	return sq.F2[F, I1, I2]("round", val, places)
}

func Sign[F c.Float, I c.Integer](val sq.Expr[F]) sq.Expr[I] {
	return sq.F[F, I]("sign", val)
}

func Sqrt[N c.Number, F c.Float](val sq.Expr[N]) sq.Expr[F] {
	return sq.F[N, F]("sqrt", val)
}

func Trunk[F c.Float, I c.Integer](val sq.Expr[F]) sq.Expr[I] {
	return sq.F[F, I]("trunk", val)
}

func TrunkTo[F c.Float, I1, I2 c.Integer](val sq.Expr[F], places sq.Expr[I1]) sq.Expr[I2] {
	return sq.F2[F, I1, I2]("trunk", val, places)
}
