package ddl

import (
	"strings"

	"github.com/orsinium-labs/sequel-ddl/dialects"
)

// Char can store an ASCII string of the given size in bytes.
func Char(size uint32) ColumnType {
	return colType0{
		cocroach:  "STRING",
		mysql:     call("CHAR", size),
		oracle:    call("CHAR", size),
		postgres:  call("CHAR", size),
		sqlite:    "TEXT",
		sqlserver: call("CHAR", size),
	}
}

// Enum is a string type with a pre-defined list of members.
func Enum(size uint32, members ...string) ColumnType {
	ms := make([]string, len(members))
	for _, m := range members {
		ms = append(ms, "'"+string(m)+"'")
	}
	suffix := "(" + strings.Join(ms, ",") + ")"
	return colType0{
		cocroach:  dialects.DataType("ENUM" + suffix),
		mysql:     dialects.DataType("ENUM" + suffix),
		oracle:    call("VARCHAR2", size),
		postgres:  dialects.DataType("ENUM" + suffix),
		sqlite:    "TEXT",
		sqlserver: "TEXT",
	}
}

// Char can store a Unicode string of the given size in byte-pairs.
func NChar(size uint32) ColumnType {
	return colType0{
		cocroach:  "STRING",
		mysql:     call("NCHAR", size),
		oracle:    call("NCHAR", size),
		postgres:  call("CHAR", size),
		sqlite:    "TEXT",
		sqlserver: call("NCHAR", size),
	}
}

// NVarChar can store a Unicode string of any length up to the given size in byte-pairs.
func NVarChar(size uint32) ColumnType {
	return colType0{
		cocroach:  "STRING",
		mysql:     call("NVARCHAR", size),
		oracle:    call("NVARCHAR2", size),
		postgres:  call("VARCHAR", size),
		sqlite:    "TEXT",
		sqlserver: call("NVARCHAR", size),
	}
}

// VarChar can store an ASCII string of any length up to the given size in bytes.
func VarChar(size uint32) ColumnType {
	return colType0{
		cocroach:  "STRING",
		mysql:     call("VARCHAR", size),
		oracle:    call("VARCHAR2", size),
		postgres:  call("VARCHAR", size),
		sqlite:    "TEXT",
		sqlserver: call("VARCHAR", size),
	}
}

// Text can store a string of any length.
func Text() ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Text()
	}
	return colType{callback}
}

// UUID is a random and unique 16-bytes identifier (RFC 4122).
func UUID() ColumnType {
	return colType0{
		cocroach:  "UUID",
		mysql:     "BINARY(16)",
		oracle:    "RAW(16)",
		postgres:  "UUID",
		sqlite:    "BLOB",
		sqlserver: "BINARY(16)",
	}
}
