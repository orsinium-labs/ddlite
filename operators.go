package sequel

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbfuncs"
)

type condition interface {
	toSQL(model any) string
}

type operator struct {
	left  any
	op    string
	right any
}

func (op operator) toSQL(m any) string {
	lname := stringify(m, op.left)
	rname := stringify(m, op.right)
	return fmt.Sprintf("%s %s %s", lname, op.op, rname)
}

func Gt[T any](field *T, value T) operator {
	return operator{left: field, op: ">", right: value}
}

func GtF[T any](field *T, value dbfuncs.Func[T]) operator {
	return operator{left: field, op: ">", right: value}
}
