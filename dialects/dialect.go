package dialects

import (
	"fmt"
	"strconv"
)

type Dialect interface {
	fmt.Stringer
	Int(bits uint8) string
	UInt(bits uint8) string
	Float(precision uint8) string
	Decimal(precision uint8, scale uint8) string
	Interval() string
	Date() string
	Text() string
}

func call2(prefix string, a, b uint8) string {
	as := strconv.FormatInt(int64(a), 10)
	bs := strconv.FormatInt(int64(b), 10)
	return prefix + "(" + as + ", " + bs + ")"
}
