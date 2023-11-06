package dialects

var SQLite Dialect = sqlite{}

type sqlite struct{}

// Placeholder implements [Dialect].
func (sqlite) Placeholder(int) string {
	return "?"
}

// Precedence implements [Dialect].
func (sqlite) Precedence(op string) uint8 {
	return sqlitePrec[op]
}

func (sqlite) True() string {
	// https://www.sqlite.org/datatype3.html#boolean_datatype
	return "1"
}

func (sqlite) False() string {
	// https://www.sqlite.org/datatype3.html#boolean_datatype
	return "0"
}

func (sqlite) Int(bits uint8) string {
	// https://www.sqlite.org/datatype3.html#boolean_datatype
	// INTEGER fits up to 8 bytes.
	return "INTEGER"
}

func (sqlite) UInt(bits uint8) string {
	return "INTEGER"
}

func (sqlite) Float(precision uint8) string {
	if precision <= 53 {
		return "REAL"
	}
	return ""
}

func (sqlite) Interval() string {
	return "INTEGER"
}

func (sqlite) Date() string {
	return "INTEGER"
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
