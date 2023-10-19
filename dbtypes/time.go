package dbtypes

import "time"

func Date() ColumnType[time.Time] {
	return colType0[time.Time]{
		cocroach:  "DATE",
		mysql:     "",
		oracle:    "DATE",
		postgres:  "DATE",
		sqlite:    "INTEGER",
		sqlserver: "",
	}
}

func DateTime() ColumnType[time.Time] {
	return colType0[time.Time]{
		cocroach:  "TIMESTAMP",
		mysql:     "",
		oracle:    "TIMESTAMP",
		postgres:  "TIMESTAMP",
		sqlite:    "INTEGER",
		sqlserver: "",
	}
}

func Interval() ColumnType[time.Duration] {
	return colType0[time.Duration]{
		cocroach:  "INTERVAL",
		mysql:     "",
		oracle:    "INTERVAL",
		postgres:  "INTERVAL",
		sqlite:    "INTEGER",
		sqlserver: "",
	}
}

func Time() ColumnType[time.Duration] {
	return colType0[time.Duration]{
		cocroach:  "TIME",
		mysql:     "",
		oracle:    "INTERVAL",
		postgres:  "TIME",
		sqlite:    "INTEGER",
		sqlserver: "",
	}
}
