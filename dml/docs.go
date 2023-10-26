// SQL statements builder for Data-Manipulation Language: SELECT, INSERT, DELETE, UPDATE.
//
// # Expressions
//
// Basic expression constructors:
//
//   - [V] is a literal value
//   - [C] is a column
//   - [M] is a column that can be NULL
//   - [F] is a function call (all arguments have the same type)
//   - [F2] is a function call with 2 arguments
//
// Comparison operators:
//
//   - [E] is "equal" (=) for a column ([C]) and a value ([V])
//   - [Eq] is "equal" (=)
//   - [Neq] is "not equal" (<>)
//   - [Gt] is "greater than" (>)
//   - [Gte] is "greater than or equal" (>=)
//   - [Lt] is "less than" (<)
//   - [Lte] is "less than or equal" (<=)
//
// Boolean logic operators:
//
//   - [And] is "AND"
//   - [Or] is "OR"
//   - [IsNull] is "IS NULL"
//   - [IsNotNull] is "IS NOT NULL"
//   - [Not] is "NOT"
//
// Other infix operators:
//
//   - [Like] is "LIKE"
//   - [NotLike] is "NOT LIKE"
//   - [Glob] is "GLOB"
//   - [NotGlob] is "NOT GLOB"
//   - [RegExp] is "REGEXP"
//   - [NotRegExp] is "NOT REGEXP"
//   - [Match] is "MATCH"
//   - [NotMatch] is "NOT MATCH"
//   - [IsDistinctFrom] is "IS DISTINCT FROM"
//   - [IsNotDistinctFrom] is "IS NOT DISTINCT FROM"
//
// And lastly, [Switch] and [SwitchBy] are "CASE", the SQL equivalent
// of the "switch" statement from Go.
//
// # Statements
//
//   - [Select] is "SELECT"
//   - [Update] is "UPDATE"
//   - [Delete] is "DELETE"
//   - [Insert] is "INSERT"
package dml
