package dbtypes

import (
	c "github.com/orsinium-labs/sequel/constraints"
)

// TODO: Serial

// Decimal is an arbitrary fixed-precision decimal number type.
func Decimal[T c.Decimal, I1, I2 c.Integer](precision I1, scale I2) ColumnType[T] {
	return colType0[T]{
		cocroach:  call2("DECIMAL", precision, scale),
		mysql:     call2("DECIMAL", precision, scale),
		oracle:    call2("NUMBER", precision, scale),
		postgres:  call2("NUMERIC", precision, scale),
		sqlite:    "NUMERIC",
		sqlserver: call2("DECIMAL", precision, scale),
	}
}

// Float32 is an inexact floating-point variable-precision number type equivalent to float32.
func Float32[T ~float32]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "REAL",
		mysql:     "FLOAT",
		oracle:    "FLOAT(63)",
		postgres:  "REAL",
		sqlite:    "REAL",
		sqlserver: "REAL",
	}
}

// Float64 is an inexact floating-point variable-precision number type equivalent to float64.
func Float64[T ~float32 | ~float64]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "DOUBLE PRECISION",
		mysql:     "DOUBLE",
		oracle:    "FLOAT",
		postgres:  "DOUBLE PRECISION",
		sqlite:    "REAL",
		sqlserver: "FLOAT",
	}
}

// Float32 is an inexact floating-point variable-precision number type of arbitrary precision.
func Float[T ~float32 | ~float64, I c.Integer](precision I) ColumnType[T] {
	// precision <= 53
	return colType0[T]{
		cocroach:  "FLOAT",
		mysql:     call("FLOAT", precision),
		oracle:    call("FLOAT", precision),
		postgres:  call("FLOAT", precision),
		sqlite:    "FLOAT",
		sqlserver: call("FLOAT", precision),
	}
}

// Int8 is an integer number from -128 to 127.
func Int8[T ~int8]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "INT2",
		mysql:     "TINYINT",
		oracle:    "NUMBER(3,0)",
		postgres:  "SMALLINT",
		sqlite:    "INTEGER",
		sqlserver: "TINYINT",
	}
}

// Int16 is an integer number from -32_768 to 32_767.
func Int16[T ~int16]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "INT2",
		mysql:     "SMALLINT",
		oracle:    "NUMBER(5,0)",
		postgres:  "SMALLINT",
		sqlite:    "INTEGER",
		sqlserver: "SMALLINT",
	}
}

// Int32 is an integer number from -2_147_483_648 to 2_147_483_647.
//
// Typically represented as INT in the database.
func Int32[T ~int32 | ~int]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "INT4",
		mysql:     "INT",
		oracle:    "NUMBER(10,0)",
		postgres:  "INTEGER",
		sqlite:    "INTEGER",
		sqlserver: "INT",
	}
}

// Int64 is an integer number from -2⁶³ to 2⁶³-1.
//
// Typically represented as BIGINT in the database.
func Int64[T ~int64 | ~int]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "INT",
		mysql:     "BIGINT",
		oracle:    "NUMBER(20,0)",
		postgres:  "BIGINT",
		sqlite:    "INTEGER",
		sqlserver: "BIGINT",
	}
}

// UInt8 is an unsigned (non-negative) integer number from 0 to 255.
//
// If the database doesn't support unsigned numbers, the equivalent of [Int16] is used.
func UInt8[T ~uint8]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "INT2",
		mysql:     "INT UNSIGNED",
		oracle:    "NUMBER(3,0)",
		postgres:  "SMALLINT",
		sqlite:    "INTEGER",
		sqlserver: "SMALLINT",
	}
}

// UInt16 is an unsigned (non-negative) integer number from 0 to 65_535.
//
// If the database doesn't support unsigned numbers, the equivalent of [Int32] is used.
func UInt16[T ~uint16]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "INT",
		mysql:     "SMALLINT UNSIGNED",
		oracle:    "NUMBER(6,0)",
		postgres:  "INTEGER",
		sqlite:    "INTEGER",
		sqlserver: "INT",
	}
}

// UInt32 is an unsigned (non-negative) integer number from 0 to 4_294_967_295.
//
// If the database doesn't support unsigned numbers, the equivalent of [Int64] is used.
func UInt32[T ~uint32 | ~uintptr]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "OID",
		mysql:     "INT UNSIGNED",
		oracle:    "NUMBER(10,0)",
		postgres:  "BIGINT",
		sqlite:    "INTEGER",
		sqlserver: "BIGINT",
	}
}

// UInt64 is an unsigned (non-negative) integer number from 0 to 2⁶⁴-1.
//
// Many databses don't support unsigned numbers. So, if that last bit isn't important,
// prefer using [Int64] instead.
func UInt64[T ~uint64 | ~uintptr]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "INT",
		mysql:     "BIGINT UNSIGNED",
		oracle:    "NUMBER(20,0)",
		postgres:  "",
		sqlite:    "INTEGER",
		sqlserver: "",
	}
}
