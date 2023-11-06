package dialects

var CocroachDB Dialect = cocroach{}

type cocroach struct{}

func (cocroach) Int(bits uint8) string {
	// https://www.cockroachlabs.com/docs/v23.1/int
	if bits <= 16 {
		return "INT2"
	}
	if bits <= 32 {
		return "INT4"
	}
	if bits <= 64 {
		return "INT8"
	}
	return ""
}

func (cocroach) UInt(bits uint8) string {
	return CocroachDB.Int(bits + 1)
}

func (cocroach) Float(precision uint8) string {
	// https://www.cockroachlabs.com/docs/v23.1/float
	if precision <= 24 {
		return "REAL"
	}
	if precision <= 53 {
		return "DOUBLE PRECISION"
	}
	return ""
}

func (cocroach) Interval() string {
	return "INTERVAL"
}

func (cocroach) Date() string {
	return "DATE"
}

func (cocroach) String() string {
	return "CocroachDB"
}

// https://github.com/cockroachdb/cockroach/pull/2305/files
