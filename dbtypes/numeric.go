package dbtypes

import (
	c "github.com/orsinium-labs/sequel/constraints"
	"github.com/orsinium-labs/sequel/dbconf"
)

// TODO: Serial

func Int(bits uint8) ColumnType {
	callback := func(c dbconf.Config) string {
		return c.Dialect.Int(bits)
	}
	return colType{callback}
}

func UInt(bits uint8) ColumnType {
	callback := func(c dbconf.Config) string {
		return c.Dialect.UInt(bits)
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
	return colType0{
		cocroach:  "REAL",
		mysql:     "FLOAT",
		oracle:    "FLOAT(63)",
		postgres:  "REAL",
		sqlite:    "REAL",
		sqlserver: "REAL",
	}
}

// Float64 is an inexact floating-point variable-precision number type equivalent to float64.
func Float64() ColumnType {
	return colType0{
		cocroach:  "DOUBLE PRECISION",
		mysql:     "DOUBLE",
		oracle:    "FLOAT",
		postgres:  "DOUBLE PRECISION",
		sqlite:    "REAL",
		sqlserver: "FLOAT",
	}
}

// Float32 is an inexact floating-point variable-precision number type of arbitrary precision.
func Float[I c.Integer](precision I) ColumnType {
	// precision <= 53
	return colType0{
		cocroach:  "FLOAT",
		mysql:     call("FLOAT", precision),
		oracle:    call("FLOAT", precision),
		postgres:  call("FLOAT", precision),
		sqlite:    "FLOAT",
		sqlserver: call("FLOAT", precision),
	}
}
