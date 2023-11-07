package dialects

import "strconv"

var SQLServer Dialect = sqlserver{}

type sqlserver struct{}

func (sqlserver) Int(bits uint8) string {
	// https://learn.microsoft.com/en-us/sql/t-sql/data-types/int-bigint-smallint-and-tinyint-transact-sql
	if bits <= 8 {
		return "TINYINT"
	}
	if bits <= 16 {
		return "SMALLINT"
	}
	if bits <= 32 {
		return "INT"
	}
	if bits <= 64 {
		return "BIGINT"
	}
	return ""
}

func (sqlserver) UInt(bits uint8) string {
	return SQLServer.Int(bits + 1)
}

func (sqlserver) Float(precision uint8) string {
	if precision > 53 {
		return ""
	}
	if precision == 24 {
		return "REAL"
	}
	if precision == 53 {
		return "FLOAT"
	}
	return "FLOAT(" + strconv.FormatInt(int64(precision), 10) + ")"
}

func (sqlserver) Decimal(precision uint8, scale uint8) string {
	return call2("DECIMAL", precision, scale)
}

func (sqlserver) Text() string {
	return "TEXT"
}

func (sqlserver) Interval() string {
	return "DATETIMEOFFSET"
}

func (sqlserver) Date() string {
	return "DATE"
}

func (sqlserver) String() string {
	return "SQLServer"
}
