package dbtypes

import "github.com/orsinium-labs/sequel/dbconf"

// Date without time.
func Date() ColumnType {
	return colType0{
		cocroach:  "DATE",
		mysql:     "DATE",
		oracle:    "DATE",
		postgres:  "DATE",
		sqlite:    "INTEGER",
		sqlserver: "DATE",
	}
}

// DateTime is date and time.
func DateTime() ColumnType {
	return colType0{
		cocroach:  "TIMESTAMP",
		mysql:     "DATETIME",
		oracle:    "TIMESTAMP",
		postgres:  "TIMESTAMP",
		sqlite:    "INTEGER",
		sqlserver: "DATETIME",
	}
}

// Interval is a difference between two datetimes.
func Interval() ColumnType {
	callback := func(c dbconf.Config) string {
		return c.Dialect.Interval()
	}
	return colType{callback}
}

// Time of the day, without date.
func Time() ColumnType {
	return colType0{
		cocroach:  "TIME",
		mysql:     "TIME",
		oracle:    "INTERVAL",
		postgres:  "TIME",
		sqlite:    "INTEGER",
		sqlserver: "TIME",
	}
}
