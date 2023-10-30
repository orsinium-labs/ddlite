package dialects

import "strconv"

var CocroachDB Dialect = cocroach{}

type cocroach struct{}

// Placeholder implements [Dialect].
func (cocroach) Placeholder(pos int) string {
	return "$" + strconv.Itoa(pos+1)
}

// Precedence implements [Dialect].
func (cocroach) Precedence(op string) uint8 {
	return psqlPrec[op]
}

func (cocroach) True() string {
	// https://www.cockroachlabs.com/docs/v23.1/bool
	return "true"
}

func (cocroach) False() string {
	// https://www.cockroachlabs.com/docs/v23.1/bool
	return "false"
}

func (cocroach) String() string {
	return "CocroachDB"
}

// https://github.com/cockroachdb/cockroach/pull/2305/files
