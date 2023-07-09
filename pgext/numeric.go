package pgext

import (
	"github.com/orsinium-labs/sequel"
	"github.com/orsinium-labs/sequel/constraints"
)

func Abs[T constraints.Number](val T) sequel.Expr[T] {
	return sequel.F[T]("abs", val)
}
