package ddl

import "github.com/orsinium-labs/sequel-ddl/dialects"

// Date without time.
func Date() ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Date()
	}
	return colType{callback}
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
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Interval()
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
