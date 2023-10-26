package dialects

var CocroachDB Dialect = cocroach{}

type cocroach struct{}

// Placeholder implements [Dialect].
func (cocroach) Placeholder() Placeholder {
	return Dollar
}

// Precedence implements [Dialect].
func (cocroach) Precedence(op string) (uint8, bool) {
	prec, ok := psqlPrec[op]
	return prec, ok
}

func (cocroach) String() string {
	return "CocroachDB"
}

// https://github.com/cockroachdb/cockroach/pull/2305/files
