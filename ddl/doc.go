// The package ddl is DDL (Data-Definition Language) SQL queries builder.
//
// DDL includes queries like CREATE TABLE, ALTER TABLE, and DROP TABLE.
//
// # Statements
//
//   - [AddColumn] is "ALTER TABLE t ADD COLUMN c"
//   - [CreateTable] is "CREATE TABLE t"
//   - [DropColumn] is "ALTER TABLE t DROP COLUMN c"
//   - [DropTable] is "DROP TABLE t"
//   - [RenameColumn] is "ALTER TABLE t RENAME COLUMN c"
//   - [RenameTable] is "ALTER TABLE t RENAME TO x"
//   - [TruncateTable] is "TRUNCATE TABLE t" (or "DELETE FROM t")
//
// # Data types
//
// Supported column types:
//
//   - Text: [Text], [Enum], [FixedChar], [VarChar].
//   - Numeric: [Int], [UInt], [Decimal], [Float32], [Float64], [Float].
//   - Time: [Date], [DateTime], [Interval], [Time].
//   - Misc: [Blob], [Bool].
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
// # Constraints
//
// Single column constraints:
//
//   - [ColumnBuilder.Check] is "CHECK"
//   - [ColumnBuilder.Collate] is "COLLATE"
//   - [ColumnBuilder.ForeignKey] is "FOREIGN KEY"
//   - [ColumnBuilder.PrimaryKey] is "PRIMARY KEY"
//   - [ColumnBuilder.Unique] is "UNIQUE"
//
// Multi-column (table) constraints:
//
//   - [Check] is "CHECK"
//   - [ForeignKey] is "FOREIGN KEY"
//   - [PrimaryKey] is "PRIMARY KEY"
//   - [Unique] is "UNIQUE"
//   - [Named] is "CONSTRAINT" (sets constraint name)
//
// [SQLite]: https://www.sqlite.org/datatype3.html
// [Oracle Database]: https://docs.oracle.com/en/database/oracle/oracle-database/23/sqlrf/Data-Types.html
// [Cocroach DB]: https://www.cockroachlabs.com/docs/stable/data-types
// [MySQL]: https://dev.mysql.com/doc/refman/8.0/en/data-types.html
// [PostgreSQL]: https://www.postgresql.org/docs/current/datatype.html
// [Microsoft SQL Server]: https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16
package ddl
