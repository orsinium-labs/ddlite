package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
)

// TODO: NChar, NVarChar

// Char can store an ASCII string of the given size in bytes.
func FixedChar(size uint32) ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.FixedChar(size)
	}
	return colType{callback}
}

// VarChar can store an ASCII string of any length up to the given size in bytes.
func VarChar(size uint32) ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.VarChar(size)
	}
	return colType{callback}
}

// Text can store a string of any length.
func Text() ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Text()
	}
	return colType{callback}
}

// Enum is a string type with a pre-defined list of members.
func Enum(members ...string) ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Enum(members)
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
