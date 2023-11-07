package dialects

var SQLite Dialect = sqlite{}

type sqlite struct{}

func (sqlite) Int(bits uint8) DataType {
	// https://www.sqlite.org/datatype3.html#boolean_datatype
	// INTEGER fits up to 8 bytes.
	return "INTEGER"
}

func (sqlite) UInt(bits uint8) DataType {
	return "INTEGER"
}

func (sqlite) Float(precision uint8) DataType {
	if precision <= 53 {
		return "REAL"
	}
	return ""
}

func (sqlite) Decimal(precision uint8, scale uint8) DataType {
	return "NUMERIC"
}

func (sqlite) Text() DataType {
	return "TEXT"
}

func (sqlite) Interval() DataType {
	return "INTEGER"
}

func (sqlite) Date() DataType {
	return "INTEGER"
}

func (sqlite) String() string {
	return "SQLite"
}
