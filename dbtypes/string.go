package dbtypes

import (
	"strings"

	c "github.com/orsinium-labs/sequel/constraints"
)

// Char can store an ASCII string of the given size in bytes.
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

// Enum is a string type with a pre-defined list of members.
func Enum[T string, I c.Integer](size I, members ...T) ColumnType[T] {
	ms := make([]string, len(members))
	for _, m := range members {
		ms = append(ms, "'"+string(m)+"'")
	}
	suffix := "(" + strings.Join(ms, ",") + ")"
	return colType0[T]{
		cocroach:  "ENUM" + suffix,
		mysql:     "ENUM" + suffix,
		oracle:    call("VARCHAR2", size),
		postgres:  "ENUM" + suffix,
		sqlite:    "TEXT",
		sqlserver: "TEXT",
	}
}

// Char can store a Unicode string of the given size in byte-pairs.
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

// NVarChar can store a Unicode string of any length up to the given size in byte-pairs.
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

// VarChar can store an ASCII string of any length up to the given size in bytes.
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

// Text can store a string of any length.
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
