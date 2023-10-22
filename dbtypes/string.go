package dbtypes

import (
	"fmt"

	c "github.com/orsinium-labs/sequel/constraints"
)

func call[I c.Integer](prefix string, size I) string {
	return fmt.Sprintf("%s(%d)", prefix, size)
}

func call2[I1, I2 c.Integer](prefix string, a I1, b I2) string {
	return fmt.Sprintf("%s(%d, %d)", prefix, a, b)
}

func Char[T ~string, I c.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     call("CHAR", size),
		oracle:    call("CHAR", size),
		postgres:  call("CHAR", size),
		sqlite:    "TEXT",
		sqlserver: "",
	}
}

func Enum[T ~string, I c.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "ENUM",
		mysql:     "ENUM",
		oracle:    call("VARCHAR2", size),
		postgres:  "ENUM",
		sqlite:    "TEXT",
		sqlserver: "",
	}
}

func NChar[T ~string, I c.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     call("NCHAR", size),
		oracle:    call("NCHAR", size),
		postgres:  call("CHAR", size),
		sqlite:    "TEXT",
		sqlserver: "",
	}
}

func NVarChar[T ~string, I c.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     call("NVARCHAR", size),
		oracle:    call("NVARCHAR2", size),
		postgres:  call("VARCHAR", size),
		sqlite:    "TEXT",
		sqlserver: "",
	}
}

func VarChar[T ~string, I c.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     call("VARCHAR", size),
		oracle:    call("VARCHAR2", size),
		postgres:  call("VARCHAR", size),
		sqlite:    "TEXT",
		sqlserver: "",
	}
}

func Text[T ~string]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     "TEXT",
		oracle:    "",
		postgres:  "TEXT",
		sqlite:    "TEXT",
		sqlserver: "",
	}
}

func UUID[T ~string]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "UUID",
		mysql:     "",
		oracle:    "BLOB",
		postgres:  "",
		sqlite:    "TEXT",
		sqlserver: "",
	}
}
