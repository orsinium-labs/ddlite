package dialects

import "fmt"

type Dialect interface {
	fmt.Stringer

	// Precedence of operators and keywords.
	//
	// The precedence is used when generating SQL for expressions
	//  to add parenthesis to sub-expressions to avoid ambiguity.
	//
	// If precedence for the given operator is unknown, zero (the lowest precedence)
	// should be returned. In this case, the operation is almost always wrapped
	// in parenthesis.
	Precedence(string) uint8

	// Placeholder style for variable binding.
	Placeholder(pos int) string
}

// const (
// 	CockroachDB Dialect = '🪳'
// 	MySQL       Dialect = '🐬'
// 	OracleDB    Dialect = '👁'
// 	PostgreSQL  Dialect = '🐘'
// 	SQLite      Dialect = '🪶'
// 	SQLServer   Dialect = '🪟'
// )
