package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
)

// TODO: Serial

// Int is a numeric type that can fit an integer value of the given maximum size in bits.
//
// Bits indicate not the maximum allowed value but the maximum size in bits needed
// to store it. One bit is always used to store the sign.
// That is, Int(8) fits numbers only up to 2^7-1=127.
//
// The Go type int8 is equivalent to the DB type Int(8), int16 to Int(16), etc.
func Int(bits uint8) ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Int(bits)
	}
	return colType{callback}
}

// UInt is a numeric type that can fit an unsigned (non-negative) integer value of the given maximum size in bits.
//
// From all the built-in dialects, only MySQL supports UInt. Well, and SQLite because SQLite
// doesn't care much about types. So, you should always specify as a fallback an [Int].
func UInt(bits uint8) ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.UInt(bits)
	}
	return colType{callback}
}

// Decimal is an arbitrary fixed-precision decimal number type.
func Decimal(precision uint8, scale uint8) ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Decimal(precision, scale)
	}
	return colType{callback}
}

// Float32 is an inexact floating-point variable-precision number type equivalent to float32.
func Float32() ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Float(24)
	}
	return colType{callback}
}

// Float64 is an inexact floating-point variable-precision number type equivalent to float64.
func Float64() ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Float(53)
	}
	return colType{callback}
}

// Float is an inexact floating-point variable-precision number type of arbitrary precision.
//
// Precision 24 is equivalent to [Float32] and precision 53 is equivalent to [Float64].
func Float(precision uint8) ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Float(precision)
	}
	return colType{callback}
}
