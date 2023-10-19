package dbtypes

import (
	"fmt"

	"github.com/orsinium-labs/sequel/constraints"
)

func call[I constraints.Integer](prefix string, size I) string {
	return fmt.Sprintf("%s(%d)", prefix, size)
}

func Char[T ~string, I constraints.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     call("CHAR", size),
		oracle:    call("CHAR", size),
		postgres:  call("CHAR", size),
		sqlite:    "TEXT",
		sqlserver: "",
	}
}

func Enum[T ~string, I constraints.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "ENUM",
		mysql:     "ENUM",
		oracle:    call("VARCHAR2", size),
		postgres:  "ENUM",
		sqlite:    "TEXT",
		sqlserver: "",
	}
}

func NChar[T ~string, I constraints.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     call("NCHAR", size),
		oracle:    call("NCHAR", size),
		postgres:  call("CHAR", size),
		sqlite:    "TEXT",
		sqlserver: "",
	}
}

func NVarChar[T ~string, I constraints.Integer](size I) ColumnType[T] {
	return colType0[T]{
		cocroach:  "STRING",
		mysql:     call("NVARCHAR", size),
		oracle:    call("NVARCHAR2", size),
		postgres:  call("VARCHAR", size),
		sqlite:    "TEXT",
		sqlserver: "",
	}
}

func VarChar[T ~string, I constraints.Integer](size I) ColumnType[T] {
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
