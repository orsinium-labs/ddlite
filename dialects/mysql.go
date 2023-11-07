package dialects

var MySQL Dialect = mysql{}

type mysql struct{}

func (mysql) Int(bits uint8) DataType {
	// https://dev.mysql.com/doc/refman/8.2/en/integer-types.html
	if bits <= 8 {
		return "TINYINT"
	}
	if bits <= 16 {
		return "SMALLINT"
	}
	if bits <= 24 {
		return "MEDIUMINT"
	}
	if bits <= 32 {
		return "INT"
	}
	if bits <= 64 {
		return "BIGINT"
	}
	return ""
}

func (mysql) UInt(bits uint8) DataType {
	return MySQL.Int(bits) + " UNSIGNED"
}

func (mysql) Float(precision uint8) DataType {
	if precision > 53 {
		return ""
	}
	if precision == 24 {
		return "FLOAT"
	}
	if precision == 53 {
		return "DOUBLE"
	}
	return call("FLOAT", precision)
}

func (mysql) Decimal(precision uint8, scale uint8) DataType {
	return call2("DECIMAL", precision, scale)
}

func (mysql) Text() DataType {
	return "TEXT"
}

func (mysql) Interval() DataType {
	return "INTEGER"
}

func (mysql) Date() DataType {
	return "DATE"
}

func (mysql) String() string {
	return "MySQL"
}
