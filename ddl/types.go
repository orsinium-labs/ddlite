package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel-ddl/dialects"
)

type ColumnType interface {
	SQL(dialects.Dialect) dialects.DataType
}

// colType0 is a column type without parametrization.
type colType0 struct {
	cocroach  dialects.DataType
	mysql     dialects.DataType
	oracle    dialects.DataType
	postgres  dialects.DataType
	sqlite    dialects.DataType
	sqlserver dialects.DataType
}

func (c colType0) SQL(dialect dialects.Dialect) dialects.DataType {
	switch dialect {
	case dialects.CocroachDB:
		return c.cocroach
	case dialects.MySQL:
		return c.mysql
	case dialects.Oracle:
		return c.oracle
	case dialects.PostgreSQL:
		return c.postgres
	case dialects.SQLite:
		return c.sqlite
	case dialects.SQLServer:
		return c.sqlserver
	default:
		return c.sqlite
	}
}

type colType struct {
	callback func(dialects.Dialect) dialects.DataType
}

func (c colType) SQL(dialect dialects.Dialect) dialects.DataType {
	return c.callback(dialect)
}

func call[I uint32 | uint8](prefix string, size I) dialects.DataType {
	return dialects.DataType(fmt.Sprintf("%s(%d)", prefix, size))
}

// Bool is a boolean type.
//
// If the database doesn't support BOOL natively,
// the smallest integer type is used.
func Bool() ColumnType {
	return colType0{
		cocroach:  "BOOL",
		mysql:     "TINYINT(1)",
		oracle:    "NUMBER(1)",
		postgres:  "BOOLEAN",
		sqlite:    "INTEGER",
		sqlserver: "TINYINT",
	}
}

// Blob is raw binary data.
func Blob() ColumnType {
	return colType0{
		cocroach:  "BYTES",
		mysql:     "BLOB",
		oracle:    "BLOB",
		postgres:  "BYTEA",
		sqlite:    "BLOB",
		sqlserver: "VARBINARY",
	}
}
