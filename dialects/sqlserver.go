package dialects

var SQLServer Dialect = sqlserver{}

type sqlserver struct{}

func (sqlserver) Int(bits uint8) DataType {
	// https://learn.microsoft.com/en-us/sql/t-sql/data-types/int-bigint-smallint-and-tinyint-transact-sql
	if bits <= 8 {
		return "TINYINT"
	}
	if bits <= 16 {
		return "SMALLINT"
	}
	if bits <= 32 {
		return "INT"
	}
	if bits <= 64 {
		return "BIGINT"
	}
	return ""
}

func (sqlserver) UInt(bits uint8) DataType {
	return ""
}

func (sqlserver) Float(precision uint8) DataType {
	if precision > 53 {
		return ""
	}
	if precision == 24 {
		return "REAL"
	}
	if precision == 53 {
		return "FLOAT"
	}
	return call("FLOAT", precision)
}

func (sqlserver) Decimal(precision uint8, scale uint8) DataType {
	return call2("DECIMAL", precision, scale)
}

func (sqlserver) Text() DataType {
	return "TEXT"
}

func (sqlserver) FixedChar(size uint32) DataType {
	return call("CHAR", size)
}

func (sqlserver) VarChar(size uint32) DataType {
	return call("VARCHAR", size)
}

func (sqlserver) Enum(members []string) DataType {
	return ""
}

func (sqlserver) Blob() DataType {
	return "VARBINARY"
}

func (sqlserver) Interval() DataType {
	return "DATETIMEOFFSET"
}

func (sqlserver) Date() DataType {
	return "DATE"
}

func (sqlserver) String() string {
	return "SQLServer"
}
