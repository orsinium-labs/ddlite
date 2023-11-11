package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type ClauseConstraint interface {
	isTableConstraint()
	tokens(dialect dialects.Dialect) tokens.Tokens
}

type tNamed struct {
	name       Safe
	constraint ClauseConstraint
}

func Named(name Safe, constraint ClauseConstraint) ClauseConstraint {
	return tNamed{name, constraint}
}

func (def tNamed) isTableConstraint() {}

func (def tNamed) tokens(d dialects.Dialect) tokens.Tokens {
	ts := tokens.New(tokens.Keyword("CONSTRAINT"))
	ts.Extend(def.constraint.tokens(d))
	return ts
}

type tUnique struct {
	names []Safe
}

func Unique(name Safe, names ...Safe) ClauseConstraint {
	return tUnique{append([]Safe{name}, names...)}
}

func (def tUnique) isTableConstraint() {}

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
func PrimaryKey(name Safe, names ...Safe) ClauseConstraint {
	return tPrimaryKey{append([]Safe{name}, names...)}
}

func (def tPrimaryKey) isTableConstraint() {}

func (def tPrimaryKey) tokens(dialects.Dialect) tokens.Tokens {
	return tokens.New(
		tokens.Keyword("PRIMARY KEY"),
		tokens.LParen(),
		tokens.Raws(def.names...),
		tokens.RParen(),
	)
}

type tCheck struct {
	expr Safe
}

func Check(expr Safe) ClauseConstraint {
	return tCheck{expr}
}

func (def tCheck) isTableConstraint() {}

func (def tCheck) tokens(dialects.Dialect) tokens.Tokens {
	return tokens.New(
		tokens.Keyword("CHECK"),
		tokens.LParen(),
		tokens.Raw(def.expr),
		tokens.RParen(),
	)
}

type tForeignKey struct {
	ref     ClauseReferences
	columns []Safe
}

func ForeignKey(ref ClauseReferences, column Safe, columns ...Safe) ClauseConstraint {
	return tForeignKey{ref, append([]Safe{column}, columns...)}
}

func (def tForeignKey) isTableConstraint() {}

func (def tForeignKey) tokens(d dialects.Dialect) tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("FOREIGN KEY"),
		tokens.LParen(),
		tokens.Raws(def.columns...),
		tokens.RParen(),
	)
	ts.Extend(def.ref.tokens(d))
	return ts
}
