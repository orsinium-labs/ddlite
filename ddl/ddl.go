package ddl

// Safe is a string that is used in SQL queries as-is, without escaping.
//
// String literals and constants are automatically considered safe.
// Variables need to be explicitly converted to Safe.
//
// Never convert to Safe untrusted input, it allows evil people to do SQL injections.
type Safe string
