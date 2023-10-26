package dialects

var PostgreSQL Dialect = psql{}

type psql struct{}

// Placeholder implements [Dialect].
func (psql) Placeholder() Placeholder {
	return Dollar
}

// Precedence implements [Dialect].
func (psql) Precedence(op string) (uint8, bool) {
	prec, ok := psqlPrec[op]
	return prec, ok
}

func (psql) String() string {
	return "PostgreSQL"
}

// https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-PRECEDENCE
var psqlPrec = map[string]uint8{
	"^": 20,

	"*": 19,
	"/": 19,
	"%": 19,

	"+": 18,
	"-": 18,

	// ...  // all other native and user-defined operators

	"BETWEEN":  16,
	"IN":       16,
	"LIKE":     16,
	"ILIKE":    16,
	"SIMILAR":  16,
	"NOT LIKE": 16,

	// check if these are supported
	"GLOB":       16,
	"NOT GLOB":   16,
	"REGEXP":     16,
	"NOT REGEXP": 16,
	"MATCH":      16,
	"NOT MATCH":  16,

	">":  15,
	">=": 15,
	"<":  15,
	"<=": 15,
	"=":  15,
	"<>": 15,

	"IS":                   14,
	"IS NULL":              14,
	"IS NOT NULL":          14,
	"IS DISTINCT FROM":     14,
	"IS NOT DISTINCT FROM": 14,

	"NOT": 13,
	"AND": 12,
	"OR":  11,

	// "CASE":                 0,
}
