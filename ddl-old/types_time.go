package ddl

// Date without time.
func Date() DataType {
	return func(dialect dl) dt { return dialect.Date() }
}

// DateTime is date and time.
//
// The datetime is always stored in the database without the timezone.
// In most of the engines, in UTC. If the timezone is important, store it separately.
func DateTime() DataType {
	return func(dialect dl) dt { return dialect.DateTime() }
}

// Interval is a difference between two datetimes.
func Interval() DataType {
	return func(dialect dl) dt { return dialect.Interval() }
}

// Time of the day, without date.
func Time() DataType {
	return func(dialect dl) dt { return dialect.Time() }
}
