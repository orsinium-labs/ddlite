package dialects

var CocroachDB Dialect = cocroach{}

type cocroach struct{}

func (cocroach) Int(bits uint8) DataType {
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

func (cocroach) UInt(bits uint8) DataType {
	return ""
}

func (cocroach) Float(precision uint8) DataType {
	// https://www.cockroachlabs.com/docs/v23.1/float
	if precision <= 24 {
		return "REAL"
	}
	if precision <= 53 {
		return "DOUBLE PRECISION"
	}
	return ""
}

func (cocroach) Decimal(precision uint8, scale uint8) DataType {
	return call2("DECIMAL", precision, scale)
}

func (cocroach) Text() DataType {
	return "STRING"
}

func (cocroach) FixedChar(size uint32) DataType {
	return call("STRING", size)
}

func (cocroach) VarChar(size uint32) DataType {
	return call("STRING", size)
}

func (cocroach) Enum(members []string) DataType {
	return callVar("ENUM", members)
}

func (cocroach) Blob() DataType {
	return "BYTES"
}

func (cocroach) Interval() DataType {
	return "INTERVAL"
}

func (cocroach) Date() DataType {
	return "DATE"
}

func (cocroach) DateTime() DataType {
	return "TIMESTAMP"
}

func (cocroach) Time() DataType {
	return "TIME"
}

func (cocroach) String() string {
	return "CocroachDB"
}

// https://github.com/cockroachdb/cockroach/pull/2305/files
