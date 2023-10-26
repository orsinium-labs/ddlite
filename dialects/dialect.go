package dialects

import "fmt"

type Dialect interface {
	fmt.Stringer

	// Precedence of operators and keywords.
	Precedence(string) (uint8, bool)

	// Placeholder style for variable binding.
	Placeholder() Placeholder
}

// const (
// 	CockroachDB Dialect = '🪳'
// 	MySQL       Dialect = '🐬'
// 	OracleDB    Dialect = '👁'
// 	PostgreSQL  Dialect = '🐘'
// 	SQLite      Dialect = '🪶'
// 	SQLServer   Dialect = '🪟'
// )
