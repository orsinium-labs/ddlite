package dbtypes

import (
	"fmt"

	c "github.com/orsinium-labs/sequel/constraints"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dialects"
)

type ColumnType interface {
	SQL(dbconf.Config) string
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

func (c colType0) SQL(conf dbconf.Config) string {
	switch conf.Dialect {
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

func call[I c.Integer](prefix string, size I) string {
	return fmt.Sprintf("%s(%d)", prefix, size)
}

func call2[I1, I2 c.Integer](prefix string, a I1, b I2) string {
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
