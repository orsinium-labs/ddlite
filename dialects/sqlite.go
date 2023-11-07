package dialects

var SQLite Dialect = sqlite{}

type sqlite struct{}

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

func (sqlite) Decimal(precision uint8, scale uint8) string {
	return "NUMERIC"
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
