package dialects

import (
	"fmt"
	"strconv"
)

type Dialect interface {
	fmt.Stringer
	Int(bits uint8) DataType
	UInt(bits uint8) DataType
	Float(precision uint8) DataType
	Decimal(precision uint8, scale uint8) DataType
	Interval() DataType
	Date() DataType
	Text() DataType
}

type DataType string

func call(prefix string, x uint8) DataType {
	s := strconv.FormatInt(int64(x), 10)
	return DataType(prefix + "(" + s + ")")
}

func call2(prefix string, a, b uint8) DataType {
	as := strconv.FormatInt(int64(a), 10)
	bs := strconv.FormatInt(int64(b), 10)
	return DataType(prefix + "(" + as + ", " + bs + ")")
}
