package sequel

import (
	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbfuncs"
)

type Predicate interface {
	Squirrel(model any) squirrel.Sqlizer
}

type tBinOp struct {
	left  any
	op    string
	right any
}

func (op tBinOp) Squirrel(m any) squirrel.Sqlizer {
	lhs := asSquirrel(m, op.left)
	rhs := asSquirrel(m, op.right)
	return squirrel.ConcatExpr(lhs, op.op, rhs)
}

func Gt[T any](field *T, value T) Predicate {
	return tBinOp{left: field, op: " > ", right: value}
}

func GtF[T any](field *T, value dbfuncs.Func[T]) Predicate {
	return tBinOp{left: field, op: " > ", right: value}
}
