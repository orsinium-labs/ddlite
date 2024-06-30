// The package ddl is DDL (Data-Definition Language) SQL queries builder for SQLite.
//
// DDL includes queries like CREATE TABLE, ALTER TABLE, and DROP TABLE.
//
// Statements:
//
//   - [AddColumn] is "ALTER TABLE t ADD COLUMN c"
//   - [CreateTable] is "CREATE TABLE t"
//   - [DropColumn] is "ALTER TABLE t DROP COLUMN c"
//   - [DropTable] is "DROP TABLE t"
//   - [RenameColumn] is "ALTER TABLE t RENAME COLUMN c"
//   - [RenameTable] is "ALTER TABLE t RENAME TO x"
//   - [TruncateTable] is "DELETE FROM t"
package ddl
