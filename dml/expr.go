package dml

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/priority"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

// Expr is an SQL expression. I can be used as part of SQL queries.
type Expr[T any] interface {
	ExprType() T
	Priority() priority.Priority
	Tokens(dbconf.Config) tokens.Tokens
}
