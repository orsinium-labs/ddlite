package priority

// Operation precedence.
//
// References:
//
//   - https://www.sqlite.org/lang_expr.html
//   - https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-PRECEDENCE
//   - https://dev.mysql.com/doc/refman/8.2/en/operator-precedence.html
//
// In PostgreSQL, [Like] has higher precedence than [Comparison].
// In SQLite, the other way around. We follow PostgreSQL here,
// but it won't make a difference because type safety will prevent you
// from chaining Like and Comparison.
type Priority uint8

const (
	Atomic     Priority = 20 // Atomic components (literals, names)
	Unary      Priority = 19 // unary plus, unary minus
	Exp        Priority = 18 // exponentiation
	Mul        Priority = 17 // multiplication, division, modulo
	Add        Priority = 16 // addition, subtraction
	Operation  Priority = 15 // all other native and user-defined operators
	Like       Priority = 14 // range containment, set membership, string matching
	Comparison Priority = 13 // comparison operators
	Is         Priority = 12 // IS TRUE, IS FALSE, IS NULL, IS DISTINCT FROM
	Not        Priority = 11 // logical negation
	And        Priority = 10 // logical conjunction
	Or         Priority = 9  // logical disjunction
)
