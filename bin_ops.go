package sequel

import (
	"github.com/Masterminds/squirrel"
)

type tBinOp[T any] struct {
	left  Expr[T]
	op    string
	right Expr[T]
}

func (tBinOp[T]) Default() bool {
	return false
}

func (op tBinOp[T]) Squirrel(ms Models) squirrel.Sqlizer {
	lhs := op.left.Squirrel(ms)
	rhs := op.right.Squirrel(ms)
	return squirrel.ConcatExpr(lhs, " ", op.op, " ", rhs)
}

func Gt[T any](left, right Expr[T]) Expr[bool] {
	return tBinOp[T]{left: left, op: ">", right: right}
}
