package pgext

import (
	"github.com/orsinium-labs/sequel"
	c "github.com/orsinium-labs/sequel/constraints"
)

func Abs[N c.Number](val sequel.Expr[N]) sequel.Expr[N] {
	return sequel.F[N, N]("abs", val)
}

func Ceil[F c.Float, I c.Integer](val sequel.Expr[F]) sequel.Expr[I] {
	return sequel.F[F, I]("ceil", val)
}

func Ceiling[F c.Float](val sequel.Expr[F]) sequel.Expr[F] {
	return sequel.F[F, F]("ceiling", val)
}

func Div[N c.Number, I c.Integer](a, b sequel.Expr[N]) sequel.Expr[I] {
	return sequel.F[N, I]("div", a, b)
}

func Exp[I c.Integer, F c.Float](val sequel.Expr[I]) sequel.Expr[F] {
	return sequel.F[I, F]("exp", val)
}

func Floor[F c.Float, I c.Integer](val sequel.Expr[F]) sequel.Expr[I] {
	return sequel.F[F, I]("floor", val)
}

func Mod[N c.Number, I c.Integer](a, b sequel.Expr[N]) sequel.Expr[I] {
	return sequel.F[N, I]("mod", a, b)
}

func Power[N c.Number, F c.Float](a, b sequel.Expr[N]) sequel.Expr[F] {
	return sequel.F[N, F]("power", a, b)
}

func Random[F c.Float]() sequel.Expr[F] {
	return sequel.F[F, F]("random")
}

func Round[F c.Float, I c.Integer](val sequel.Expr[F]) sequel.Expr[I] {
	return sequel.F[F, I]("power", val)
}

func Sign[F c.Float, I c.Integer](val sequel.Expr[F]) sequel.Expr[I] {
	return sequel.F[F, I]("sign", val)
}

func Sqrt[N c.Number, F c.Float](val sequel.Expr[N]) sequel.Expr[F] {
	return sequel.F[N, F]("sqrt", val)
}

func Trunk[F c.Float, I c.Integer](val sequel.Expr[F]) sequel.Expr[I] {
	return sequel.F[F, I]("trunk", val)
}
