package dbtypes

import (
	c "github.com/orsinium-labs/sequel/constraints"
)

// TODO: Serial

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

func Int8[T ~int8]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "INT",
		mysql:     "TINYINT",
		oracle:    "NUMBER(3,0)",
		postgres:  "SMALLINT",
		sqlite:    "INTEGER",
		sqlserver: "TINYINT",
	}
}

func Int16[T ~int16]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "INT",
		mysql:     "SMALLINT",
		oracle:    "NUMBER(5,0)",
		postgres:  "SMALLINT",
		sqlite:    "INTEGER",
		sqlserver: "SMALLINT",
	}
}

func Int32[T ~int32 | ~int]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "INT",
		mysql:     "INT",
		oracle:    "NUMBER(10,0)",
		postgres:  "INTEGER",
		sqlite:    "INTEGER",
		sqlserver: "INT",
	}
}

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

func UInt8[T ~uint8]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "INT",
		mysql:     "INT UNSIGNED",
		oracle:    "NUMBER(3,0)",
		postgres:  "",
		sqlite:    "INTEGER",
		sqlserver: "",
	}
}

func UInt16[T ~uint16]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "INT",
		mysql:     "SMALLINT UNSIGNED",
		oracle:    "NUMBER(6,0)",
		postgres:  "",
		sqlite:    "INTEGER",
		sqlserver: "",
	}
}

func UInt32[T ~uint32 | ~uintptr]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "OID",
		mysql:     "INT UNSIGNED",
		oracle:    "NUMBER(10,0)",
		postgres:  "",
		sqlite:    "INTEGER",
		sqlserver: "",
	}
}

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
