package dialects

var Oracle Dialect = oracle{}

type oracle struct{}

func (oracle) Int(bits uint8) DataType {
	// https://docs.oracle.com/en/database/oracle/oracle-database/23/sqlrf/Data-Types.html
	return call("NUMBER", bits)
}

func (oracle) UInt(bits uint8) DataType {
	return Oracle.Int(bits + 1)
}

func (oracle) Float(precision uint8) DataType {
	if precision > 53 {
		return ""
	}
	if precision == 24 {
		return "FLOAT(63)"
	}
	if precision == 53 {
		return "FLOAT"
	}
	return call("FLOAT", precision)
}

func (oracle) Decimal(precision uint8, scale uint8) DataType {
	return call2("NUMBER", precision, scale)
}

func (oracle) Text() DataType {
	return ""
}

func (oracle) FixedChar(size uint32) DataType {
	return call("CHAR", size)
}

func (oracle) VarChar(size uint32) DataType {
	return call("VARCHAR2", size)
}

func (oracle) Enum(members []string) DataType {
	return ""
}

func (oracle) Interval() DataType {
	return "INTERVAL"
}

func (oracle) Date() DataType {
	return "DATE"
}

func (oracle) String() string {
	return "Oracle"
}
