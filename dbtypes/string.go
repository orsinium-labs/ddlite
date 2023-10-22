package dbtypes

import (
	c "github.com/orsinium-labs/sequel/constraints"
)

func Char[T ~string, I c.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     call("CHAR", size),
		oracle:    call("CHAR", size),
		postgres:  call("CHAR", size),
		sqlite:    "TEXT",
		sqlserver: call("CHAR", size),
	}
}

func Enum[T ~string, I c.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "ENUM",
		mysql:     "ENUM",
		oracle:    call("VARCHAR2", size),
		postgres:  "ENUM",
		sqlite:    "TEXT",
		sqlserver: "TEXT",
	}
}

func NChar[T ~string, I c.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     call("NCHAR", size),
		oracle:    call("NCHAR", size),
		postgres:  call("CHAR", size),
		sqlite:    "TEXT",
		sqlserver: call("NCHAR", size),
	}
}

func NVarChar[T ~string, I c.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     call("NVARCHAR", size),
		oracle:    call("NVARCHAR2", size),
		postgres:  call("VARCHAR", size),
		sqlite:    "TEXT",
		sqlserver: call("NVARCHAR", size),
	}
}

func VarChar[T ~string, I c.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     call("VARCHAR", size),
		oracle:    call("VARCHAR2", size),
		postgres:  call("VARCHAR", size),
		sqlite:    "TEXT",
		sqlserver: call("VARCHAR", size),
	}
}

func Text[T ~string]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     "TEXT",
		oracle:    "",
		postgres:  "TEXT",
		sqlite:    "TEXT",
		sqlserver: "TEXT", // Should be NTEXT?
	}
}

// UUID is a random and unique 16-bytes identifier (RFC 4122).
func UUID[T ~string]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "UUID",
		mysql:     "BINARY(16)",
		oracle:    "RAW(16)",
		postgres:  "UUID",
		sqlite:    "BLOB",
		sqlserver: "BINARY(16)",
	}
}
