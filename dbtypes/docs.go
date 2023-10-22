// The package dbtypes defines column data types used for CREATE TABLE queries.
//
// References used to define the data types translation for each dialect:
//
//   - [SQLite]
//   - [Oracle Database]
//   - [Cocroach DB]
//   - [MySQL]
//   - [PostgreSQL]
//   - [Microsoft SQL Server]
//
// Supported column types:
//
//   - Text: [Char], [NChar], [NVarChar], [VarChar], [Text], [Enum].
//   - Binary: [Blob], [UUID].
//   - Signed integer: [Int8], [Int16], [Int32], [Int64].
//   - Unsigned integer: [UInt8], [UInt16], [UInt32], [UInt64].
//   - Other numeric: [Decimal], [Float32], [Float64], [Float].
//   - Time: [Date], [DateTime], [Interval], [Time].
//   - [Bool].
//
// [SQLite]: https://www.sqlite.org/datatype3.html
// [Oracle Database]: https://docs.oracle.com/en/database/oracle/oracle-database/23/sqlrf/Data-Types.html
// [Cocroach DB]: https://www.cockroachlabs.com/docs/stable/data-types
// [MySQL]: https://dev.mysql.com/doc/refman/8.0/en/data-types.html
// [PostgreSQL]: https://www.postgresql.org/docs/current/datatype.html
// [Microsoft SQL Server]: https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16
package dbtypes
