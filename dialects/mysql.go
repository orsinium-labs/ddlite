package dialects

var MySQL Dialect = mysql{}

type mysql struct{}

// Placeholder implements [Dialect].
func (mysql) Placeholder() Placeholder {
	return Question
}

// Precedence implements [Dialect].
func (mysql) Precedence(op string) uint8 {
	return mysqlPrec[op]
}

func (mysql) String() string {
	return "MySQL"
}

// https://dev.mysql.com/doc/refman/8.2/en/operator-precedence.html
var mysqlPrec = map[string]uint8{
	"INTERVAL": 22,

	"BINARY":  21,
	"COLLATE": 21,

	"!": 20,

	"^": 19,

	"*":   18,
	"/":   18,
	"%":   18,
	"DIV": 18,
	"MOD": 18,

	"+": 17,
	"-": 17,

	"<<": 16,
	">>": 16,

	"&": 15,

	"|": 14,

	"=":             13,
	"<=>":           13,
	">=":            13,
	">":             13,
	"<=":            13,
	"<":             13,
	"<>":            13,
	"!=":            13,
	"IS":            13,
	"IS NULL":       13,
	"IS NOT NULL":   13,
	"IS NOT":        13,
	"LIKE":          13,
	"NOT LIKE":      13,
	"REGEXP":        13,
	"NOT REGEXP":    13,
	"IN":            13,
	"NOT IN":        13,
	"MEMBER OF":     13,
	"NOT MEMBER OF": 13,

	"BETWEEN":     12,
	"NOT BETWEEN": 12,
	"CASE":        12,

	"NOT": 11,

	"AND": 10,
	"&&":  10,

	"XOR": 9,

	"OR": 8,
	"||": 8,

	":=": 7,

	// "MATCH":                14,
	// "NOT MATCH":            14,
	// "GLOB":                 14,
	// "NOT GLOB":             14,
}
