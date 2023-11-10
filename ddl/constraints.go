package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type Constraint interface {
	isConstraint()
	tokens(dialect dialects.Dialect) tokens.Tokens
}

type tUnique struct {
	names []Safe
}

func Unique(name Safe, names ...Safe) Constraint {
	return tUnique{names: append([]Safe{name}, names...)}
}

func (def tUnique) isConstraint() {}

func (def tUnique) tokens(dialects.Dialect) tokens.Tokens {
	return tokens.New(
		tokens.Keyword("UNIQUE"),
		tokens.LParen(),
		tokens.Raws(def.names...),
		tokens.RParen(),
	)
}

type tPrimaryKey struct {
	names []Safe
}

// Mark multiple columns as a compound primary key.
//
// If you want the table to have a single-column primary key,
// use [ColumnBuilder.PrimaryKey] instead.
func PrimaryKey(name1, name2 Safe, names ...Safe) Constraint {
	return tPrimaryKey{names: append([]Safe{name1, name2}, names...)}
}

func (def tPrimaryKey) isConstraint() {}

func (def tPrimaryKey) tokens(dialects.Dialect) tokens.Tokens {
	return tokens.New(
		tokens.Keyword("PRIMARY KEY"),
		tokens.LParen(),
		tokens.Raws(def.names...),
		tokens.RParen(),
	)
}
