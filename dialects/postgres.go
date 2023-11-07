package dialects

import "strconv"

var PostgreSQL Dialect = psql{}

type psql struct{}

func (psql) Int(bits uint8) string {
	// https://www.postgresql.org/docs/current/datatype-numeric.html
	if bits <= 16 {
		return "SMALLINT"
	}
	if bits <= 32 {
		return "INTEGER"
	}
	// BIGINT fits anything up to 2^8=256 which is +1 from what `bits` (uint8) can fit.
	return "BIGINT"
}

func (psql) UInt(bits uint8) string {
	return PostgreSQL.Int(bits + 1)
}

func (psql) Float(precision uint8) string {
	// https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-FLOAT
	if precision > 53 {
		return ""
	}
	if precision == 24 {
		return "REAL"
	}
	if precision == 53 {
		return "DOUBLE PRECISION"
	}
	return "FLOAT(" + strconv.FormatInt(int64(precision), 10) + ")"
}

func (psql) Decimal(precision uint8, scale uint8) string {
	return call2("NUMERIC", precision, scale)
}

func (psql) Text() string {
	return "TEXT"
}

func (psql) Interval() string {
	return "INTERVAL"
}

func (psql) Date() string {
	return "DATE"
}

func (psql) String() string {
	return "PostgreSQL"
}
