package dialects

import "strconv"

var MySQL Dialect = mysql{}

type mysql struct{}

func (mysql) Int(bits uint8) string {
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

func (mysql) UInt(bits uint8) string {
	return MySQL.Int(bits) + " UNSIGNED"
}

func (mysql) Float(precision uint8) string {
	if precision > 53 {
		return ""
	}
	if precision == 24 {
		return "FLOAT"
	}
	if precision == 53 {
		return "DOUBLE"
	}
	return "FLOAT(" + strconv.FormatInt(int64(precision), 10) + ")"
}

func (mysql) Decimal(precision uint8, scale uint8) string {
	return call2("DECIMAL", precision, scale)
}

func (mysql) Text() string {
	return "TEXT"
}

func (mysql) Interval() string {
	return "INTEGER"
}

func (mysql) Date() string {
	return "DATE"
}

func (mysql) String() string {
	return "MySQL"
}
