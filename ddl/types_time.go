package ddl

import "github.com/orsinium-labs/sequel-ddl/dialects"

// Date without time.
func Date() ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Date()
	}
	return colType{callback}
}

// DateTime is date and time.
//
// The datetime is always stored in the database without the timezone.
// In most of the engines, in UTC. If the timezone is important, store it separately.
func DateTime() ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.DateTime()
	}
	return colType{callback}
}

// Interval is a difference between two datetimes.
func Interval() ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Interval()
	}
	return colType{callback}
}

// Time of the day, without date.
func Time() ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Time()
	}
	return colType{callback}
}
