package dbtypes

import (
	c "github.com/orsinium-labs/ddl/constraints"
	"github.com/orsinium-labs/ddl/dialects"
)

// TODO: Serial

func Int(bits uint8) ColumnType {
	callback := func(dialect dialects.Dialect) string {
		return dialect.Int(bits)
	}
	return colType{callback}
}

func UInt(bits uint8) ColumnType {
	callback := func(dialect dialects.Dialect) string {
		return dialect.UInt(bits)
	}
	return colType{callback}
}

// Decimal is an arbitrary fixed-precision decimal number type.
func Decimal[I1, I2 c.Integer](precision I1, scale I2) ColumnType {
	return colType0{
		cocroach:  call2("DECIMAL", precision, scale),
		mysql:     call2("DECIMAL", precision, scale),
		oracle:    call2("NUMBER", precision, scale),
		postgres:  call2("NUMERIC", precision, scale),
		sqlite:    "NUMERIC",
		sqlserver: call2("DECIMAL", precision, scale),
	}
}

// Float32 is an inexact floating-point variable-precision number type equivalent to float32.
func Float32() ColumnType {
	callback := func(dialect dialects.Dialect) string {
		return dialect.Float(24)
	}
	return colType{callback}
}

// Float64 is an inexact floating-point variable-precision number type equivalent to float64.
func Float64() ColumnType {
	callback := func(dialect dialects.Dialect) string {
		return dialect.Float(53)
	}
	return colType{callback}
}

// Float32 is an inexact floating-point variable-precision number type of arbitrary precision.
func Float(precision uint8) ColumnType {
	callback := func(dialect dialects.Dialect) string {
		return dialect.Float(precision)
	}
	return colType{callback}
}
