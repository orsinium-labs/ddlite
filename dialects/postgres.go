package dialects

import "strconv"

var PostgreSQL Dialect = psql{}

type psql struct{}

// Placeholder implements [Dialect].
func (psql) Placeholder(pos int) string {
	return "$" + strconv.Itoa(pos+1)
}

// Precedence implements [Dialect].
func (psql) Precedence(op string) uint8 {
	return psqlPrec[op]
}

func (psql) True() string {
	// https://www.postgresql.org/docs/current/datatype-boolean.html
	return "TRUE"
}

func (psql) False() string {
	// https://www.postgresql.org/docs/current/datatype-boolean.html
	return "FALSE"
}

func (psql) Int(bits uint8) string {
	// https://www.postgresql.org/docs/current/datatype-numeric.html
	if bits <= 16 {
		return "SMALLINT"
	}
	if bits <= 32 {
		return "INTEGER"
	}
	// BIGINT fits anything up to 2^8=256 which is +1 from what `bits` (uint8) can fit.
	return "BIGINT"
}

func (psql) UInt(bits uint8) string {
	return PostgreSQL.Int(bits + 1)
}

func (psql) Float(precision uint8) string {
	// https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-FLOAT
	if precision > 53 {
		return ""
	}
	if precision == 24 {
		return "REAL"
	}
	if precision == 53 {
		return "DOUBLE PRECISION"
	}
	return "FLOAT(" + strconv.FormatInt(int64(precision), 10) + ")"
}

func (psql) Interval() string {
	return "INTERVAL"
}

func (psql) Date() string {
	return "DATE"
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
