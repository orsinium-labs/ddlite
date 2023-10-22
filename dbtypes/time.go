package dbtypes

import "time"

// Date without time.
func Date() ColumnType[time.Time] {
	return colType0[time.Time]{
		cocroach:  "DATE",
		mysql:     "DATE",
		oracle:    "DATE",
		postgres:  "DATE",
		sqlite:    "INTEGER",
		sqlserver: "DATE",
	}
}

// DateTime is date and time.
func DateTime() ColumnType[time.Time] {
	return colType0[time.Time]{
		cocroach:  "TIMESTAMP",
		mysql:     "DATETIME",
		oracle:    "TIMESTAMP",
		postgres:  "TIMESTAMP",
		sqlite:    "INTEGER",
		sqlserver: "DATETIME",
	}
}

// Interval is a difference between two datetimes.
func Interval() ColumnType[time.Duration] {
	return colType0[time.Duration]{
		cocroach:  "INTERVAL",
		mysql:     "",
		oracle:    "INTERVAL",
		postgres:  "INTERVAL",
		sqlite:    "INTEGER",
		sqlserver: "DATETIMEOFFSET",
	}
}

// Time of the day, without date.
func Time() ColumnType[time.Duration] {
	return colType0[time.Duration]{
		cocroach:  "TIME",
		mysql:     "TIME",
		oracle:    "INTERVAL",
		postgres:  "TIME",
		sqlite:    "INTEGER",
		sqlserver: "TIME",
	}
}
