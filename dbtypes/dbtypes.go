package dbtypes

import (
	"fmt"

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

// colType1 is a parametrized column type with one argument.
type colType1[T any] struct {
	name string
	arg  int
}

func (c colType1[T]) Default() T {
	return *new(T)
}

func (c colType1[T]) SQL(conf dbconfig.Config) string {
	return fmt.Sprintf("%s(%d)", c.name, c.arg)
}

// -- NUMERIC -- //

// SMALLINT, small-range integer
func SmallInt() ColumnType[int] {
	return colType0[int]{
		cocroach:  "SMALLINT",
		mysql:     "SMALLINT",
		oracle:    "SMALLINT", // ?
		postgres:  "SMALLINT",
		sqlite:    "SMALLINT",
		sqlserver: "SMALLINT",
	}
}

// INTEGER, typical choice for integer
func Integer() ColumnType[int] {
	return colType0[int]{
		cocroach:  "INTEGER",
		mysql:     "INT",
		oracle:    "INTEGER",
		postgres:  "INTEGER",
		sqlite:    "INTEGER",
		sqlserver: "INT",
	}
}

// BIGINT, large-range integer
func BigInt() ColumnType[int] {
	return colType0[int]{
		cocroach:  "INTEGER", // ?
		mysql:     "BIGINT",
		oracle:    "SMALLINT",
		postgres:  "BIGINT",
		sqlite:    "INTEGER",
		sqlserver: "INT",
	}
}

// REAL, variable-precision, inexact
func Real() ColumnType[float32] {
	return colType0[float32]{
		cocroach:  "DOUBLE PRECISION",
		mysql:     "FLOAT",
		oracle:    "BINARY_FLOAT",
		postgres:  "REAL",
		sqlite:    "REAL",
		sqlserver: "FLOAT",
	}
}

// Float, variable-precision, inexact
func Float() ColumnType[float64] {
	return colType0[float64]{
		cocroach:  "DOUBLE PRECISION",
		mysql:     "DOUBLE",
		oracle:    "BINARY_DOUBLE",
		postgres:  "DOUBLE PRECISION",
		sqlite:    "REAL",
		sqlserver: "REAL",
	}
}

// // SMALLSERIAL, small autoincrementing integer
// func SmallSerial() ColumnType[int] {
// 	return colType0[int]{
// 		cocroach:  "",
// 		mysql:     "",
// 		oracle:    "",
// 		postgres:  "",
// 		sqlite:    "",
// 		sqlserver: "",
// 	}
// }

// // SERIAL, autoincrementing integer
// func Serial() ColumnType[int] {
// 	return colType0[int]{
// 		cocroach:  "",
// 		mysql:     "",
// 		oracle:    "",
// 		postgres:  "",
// 		sqlite:    "",
// 		sqlserver: "",
// 	}
// }

// // BIGSERIAL, large autoincrementing integer
// func BigSerial() ColumnType[int] {
// 	return colType0[int]{
// 		cocroach:  "",
// 		mysql:     "",
// 		oracle:    "",
// 		postgres:  "",
// 		sqlite:    "",
// 		sqlserver: "",
// 	}
// }

// -- CHARACTER -- //

// TEXT, variable unlimited length
func Text() ColumnType[string] {
	return colType0[string]{
		cocroach:  "STRING",
		mysql:     "TEXT",
		oracle:    "VARCHAR2(4000)",
		postgres:  "TEXT",
		sqlite:    "TEXT",
		sqlserver: "TEXT",
	}
}

// CHARACTER, fixed-length, blank padded
func Character(n int) ColumnType[string] {
	return colType1[string]{"CHARACTER", n}
}

// VARCHAR, variable-length with limit
func VarChar(n int) ColumnType[string] {
	return colType1[string]{"VARCHAR", n}
}

// -- MISC -- //

// BOOLEAN, state of true or false
func Boolean() ColumnType[bool] {
	return colType0[bool]{
		cocroach:  "",
		mysql:     "TINYINT(1)",
		oracle:    "SMALLINT",
		postgres:  "BOOLEAN",
		sqlite:    "INTEGER",
		sqlserver: "TINYINT",
	}
}

// // DATE, date (no time of day)
// func Date() ColumnType[time.Time] {
// 	return colType0[time.Time]{"DATE"}
// }

// // TIMESTAMP, both date and time
// func TimeStamp() ColumnType[time.Time] {
// 	return colType0[time.Time]{"TIMESTAMP"}
// }
