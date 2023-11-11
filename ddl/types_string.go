package ddl

// TODO: NChar, NVarChar

// FixedChar can store an string of the fixed size.
//
// Use it when all possible values have the same fixed length. For example:
//
//   - country codes ("UK", 2 chars)
//   - language codes ("en-UK", 4 chars if you remove "-")
//   - IATA airport codes ("AMS", 3 chars)
//
// The size is usually in UTF-8 code points but can also mean bytes,
// especially in older database engines. If compatibility is important,
// use this column type only for ASCII values.
//
// If the list of possible values is known in advance and doesn't change too often,
// prefer using [Enum] instead (assuming that your database engine supports it).
func FixedChar(size uint32) ClauseDataType {
	return func(dialect dl) dt { return dialect.FixedChar(size) }
}

// VarChar can store a string of any length up to the given size.
//
// The size is usually in UTF-8 code points but can also mean bytes,
// especially in older database engines. If compatibility is important,
// use this column type only for ASCII values.
//
// If the list of possible values is known in advance and doesn't change too often,
// prefer using [Enum] instead (assuming that your database engine supports it).
//
// If all possible values have the same length, prefer using [FixedChar] instead.
//
// If the maximum length is not known in advance or too big, use [Text] instead.
func VarChar(size uint32) ClauseDataType {
	return func(dialect dl) dt { return dialect.VarChar(size) }
}

// Text can store a string of any length.
//
// In some database engines, it's better to use [VarChar] instead whenever possible
// to prevent [write amplification].
//
// [write amplification]: https://en.wikipedia.org/wiki/Write_amplification
func Text() ClauseDataType {
	return func(dialect dl) dt { return dialect.Text() }
}

// Enum is a string type with a pre-defined list of members.
//
// Only some database engines support it. If compatibility is important,
// use [FixedChar] or [VarChar] as a fallback.
func Enum(members ...string) ClauseDataType {
	return func(dialect dl) dt { return dialect.Enum(members) }
}

// Blob is a raw binary data type or variable length.
func Blob() ClauseDataType {
	return func(dialect dl) dt { return dialect.Blob() }
}
