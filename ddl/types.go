package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel-ddl/dialects"
)

type ColumnType interface {
	SQL(dialects.Dialect) string
}

// colType0 is a column type without parametrization.
type colType0 struct {
	cocroach  string
	mysql     string
	oracle    string
	postgres  string
	sqlite    string
	sqlserver string
}

func (c colType0) SQL(dialect dialects.Dialect) string {
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
	callback func(dialects.Dialect) string
}

func (c colType) SQL(dialect dialects.Dialect) string {
	return c.callback(dialect)
}

func call[I uint32 | uint8](prefix string, size I) string {
	return fmt.Sprintf("%s(%d)", prefix, size)
}

func call2[I1, I2 uint32 | uint8](prefix string, a I1, b I2) string {
	return fmt.Sprintf("%s(%d, %d)", prefix, a, b)
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
