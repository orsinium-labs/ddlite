package dialects

import (
	"fmt"
	"strconv"
	"strings"
)

type Dialect interface {
	fmt.Stringer

	// numeric types
	Int(bits uint8) DataType
	UInt(bits uint8) DataType
	Float(precision uint8) DataType
	Decimal(precision uint8, scale uint8) DataType

	// time types
	Interval() DataType
	Date() DataType
	DateTime() DataType
	Time() DataType

	// string types
	Text() DataType
	FixedChar(size uint32) DataType
	VarChar(size uint32) DataType
	Enum(members []string) DataType
	Blob() DataType
}

type DataType string

func call[T uint8 | uint32](prefix string, x T) DataType {
	s := strconv.FormatUint(uint64(x), 10)
	return DataType(prefix + "(" + s + ")")
}

func call2(prefix string, a, b uint8) DataType {
	as := strconv.FormatUint(uint64(a), 10)
	bs := strconv.FormatUint(uint64(b), 10)
	return DataType(prefix + "(" + as + ", " + bs + ")")
}

func callVar(prefix string, args []string) DataType {
	suffix := strings.Join(args, ", ")
	return DataType(prefix + "(" + suffix + ")")
}
