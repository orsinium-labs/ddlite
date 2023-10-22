package dbtypes

import (
	"fmt"

	c "github.com/orsinium-labs/sequel/constraints"
	"github.com/orsinium-labs/sequel/dbconfig"
)

type ColumnType[T any] interface {
	Default() T
	SQL(dbconfig.Config) string
}

// colType0 is a column type without parametrization.
type colType0[T any] struct {
	cocroach  string
	mysql     string
	oracle    string
	postgres  string
	sqlite    string
	sqlserver string
}

func (c colType0[T]) Default() T {
	return *new(T)
}

func (c colType0[T]) SQL(conf dbconfig.Config) string {
	switch conf.Dialect {
	case dbconfig.CockroachDB:
		return c.cocroach
	case dbconfig.MySQL:
		return c.mysql
	case dbconfig.OracleDB:
		return c.oracle
	case dbconfig.PostgreSQL:
		return c.postgres
	case dbconfig.SQLite:
		return c.sqlite
	case dbconfig.SQLServer:
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
func Bool[T ~bool]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "BOOL",
		mysql:     "TINYINT(1)",
		oracle:    "NUMBER(1)",
		postgres:  "BOOLEAN",
		sqlite:    "INTEGER",
		sqlserver: "TINYINT",
	}
}

// Blob is raw binary data.
func Blob[T ~[]byte]() ColumnType[T] {
	return colType0[T]{
		cocroach:  "BYTES",
		mysql:     "BLOB",
		oracle:    "BLOB",
		postgres:  "BYTEA",
		sqlite:    "BLOB",
		sqlserver: "VARBINARY",
	}
}
