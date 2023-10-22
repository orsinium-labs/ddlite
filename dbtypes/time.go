package dbtypes

import "time"

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
