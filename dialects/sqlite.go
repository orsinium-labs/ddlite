package dialects

var SQLite Dialect = sqlite{}

type sqlite struct{}

// Placeholder implements [Dialect].
func (sqlite) Placeholder() Placeholder {
	return Question
}

// Precedence implements [Dialect].
func (sqlite) Precedence(op string) (uint8, bool) {
	prec, ok := sqlitePrec[op]
	return prec, ok
}

func (sqlite) String() string {
	return "SQLite"
}

// https://www.sqlite.org/lang_expr.html
var sqlitePrec = map[string]uint8{
	"COLLATE": 21,

	"||":  20,
	"->":  20,
	"->>": 20,

	"*": 19,
	"/": 19,
	"%": 19,

	"+": 18,
	"-": 18,

	"&":  17,
	"|":  17,
	"<<": 17,
	">>": 17,

	"ESCAPE": 16,

	"<":  15,
	">":  15,
	"<=": 15,
	">=": 15,

	"=":                    14,
	"==":                   14,
	"<>":                   14,
	"!=":                   14,
	"IS":                   14,
	"IS NOT":               14,
	"IS DISTINCT FROM":     14,
	"IS NOT DISTINCT FROM": 14,
	"BETWEEN":              14,
	"NOT BETWEEN":          14,
	"IN":                   14,
	"NOT IN":               14,
	"MATCH":                14,
	"NOT MATCH":            14,
	"LIKE":                 14,
	"NOT LIKE":             14,
	"REGEXP":               14,
	"NOT REGEXP":           14,
	"GLOB":                 14,
	"NOT GLOB":             14,
	"IS NULL":              14,
	"IS NOT NULL":          14,

	"NOT": 13,
	"AND": 12,
	"OR":  11,

	// "CASE":                 0,
}