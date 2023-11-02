package dialects

import "strconv"

var SQLServer Dialect = sqlserver{}

type sqlserver struct{}

// Placeholder implements [Dialect].
func (sqlserver) Placeholder(pos int) string {
	return "@p" + strconv.Itoa(pos+1)
}

// Precedence implements [Dialect].
func (sqlserver) Precedence(op string) uint8 {
	return sqlserverPrec[op]
}

func (sqlserver) True() string {
	return "1"
}

func (sqlserver) False() string {
	return "0"
}

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

func (sqlserver) Interval() string {
	return "DATETIMEOFFSET"
}

func (sqlserver) Date() string {
	return "DATE"
}

func (sqlserver) String() string {
	return "SQLServer"
}

// https://learn.microsoft.com/en-us/sql/t-sql/language-elements/operator-precedence-transact-sql?view=sql-server-ver16
var sqlserverPrec = map[string]uint8{
	"~": 20,

	"*": 19,
	"/": 19,
	"%": 19,

	"+": 18,
	"-": 18,
	"&": 18,
	"^": 18,
	"|": 18,

	"=":  17,
	">":  17,
	"<":  17,
	">=": 17,
	"<=": 17,
	"<>": 17,
	"!=": 17,
	"!>": 17,
	"!<": 17,

	"NOT": 16,

	"AND": 15,

	"ALL":      14,
	"ANY":      14,
	"BETWEEN":  14,
	"IN":       14,
	"NOT IN":   14,
	"LIKE":     14,
	"NOT LIKE": 14,
	"OR":       14,
	"SOME":     14,
}
