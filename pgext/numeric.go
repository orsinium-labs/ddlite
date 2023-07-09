package pgext

import "github.com/orsinium-labs/sequel"

type Args []any

func Abs(val int) sequel.Expr[int] {
	return sequel.F[int]("abs", val)
}
