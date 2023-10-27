package dialects

var SQLServer Dialect = sqlserver{}

type sqlserver struct{}

// Placeholder implements [Dialect].
func (sqlserver) Placeholder() Placeholder {
	return AtP
}

// Precedence implements [Dialect].
func (sqlserver) Precedence(op string) uint8 {
	return sqlserverPrec[op]
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
