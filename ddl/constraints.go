package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type TableConstraint interface {
	isTableConstraint()
	tokens(dialect dialects.Dialect) tokens.Tokens
}

type tNamedConstraint struct {
	name       Safe
	constraint TableConstraint
}

func NamedConstraint(name Safe, constraint TableConstraint) TableConstraint {
	return tNamedConstraint{name, constraint}
}

func (def tNamedConstraint) isTableConstraint() {}

func (def tNamedConstraint) tokens(d dialects.Dialect) tokens.Tokens {
	ts := tokens.New(tokens.Keyword("CONSTRAINT"))
	ts.Extend(def.constraint.tokens(d))
	return ts
}

type tUnique struct {
	names []Safe
}

func Unique(name Safe, names ...Safe) TableConstraint {
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
func PrimaryKey(name1, name2 Safe, names ...Safe) TableConstraint {
	return tPrimaryKey{append([]Safe{name1, name2}, names...)}
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

func Check(expr Safe) TableConstraint {
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
