package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type ClauseTableConstraint struct {
	name       string
	constraint ClauseConstraint
	columns    []Safe
}

func Constraint(
	name string,
	constraint ClauseConstraint,
	column Safe,
	columns ...Safe,
) ClauseTableConstraint {
	return ClauseTableConstraint{
		name:       name,
		constraint: constraint,
		columns:    append([]Safe{column}, columns...),
	}
}

func (con ClauseTableConstraint) tokens() tokens.Tokens {
	ts := tokens.New()
	if con.name != "" {
		ts.Add(tokens.Keyword("CHECK"))
		ts.Add(tokens.Raw(con.name))
	}
	ts.Extend(con.constraint.tableTokens(con.columns))
	return ts
}

type ClauseConstraint interface {
	columnTokens() tokens.Tokens
	tableTokens(cols []Safe) tokens.Tokens
}

type tUnique struct{}

func Unique() ClauseConstraint {
	return tUnique{}
}

func (def tUnique) columnTokens() tokens.Tokens {
	return tokens.New(
		tokens.Keyword("UNIQUE"),
	)
}
func (def tUnique) tableTokens(cols []Safe) tokens.Tokens {
	return tokens.New(
		tokens.Keyword("UNIQUE"),
		tokens.LParen(),
		tokens.Raws(cols...),
		tokens.RParen(),
	)
}

type tPrimaryKey struct{}

// Mark multiple columns as a compound primary key.
//
// If you want the table to have a single-column primary key,
// use [ColumnBuilder.PrimaryKey] instead.
func PrimaryKey() ClauseConstraint {
	return tPrimaryKey{}
}

func (def tPrimaryKey) columnTokens() tokens.Tokens {
	return tokens.New(
		tokens.Keyword("PRIMARY KEY"),
	)
}

func (def tPrimaryKey) tableTokens(cols []Safe) tokens.Tokens {
	return tokens.New(
		tokens.Keyword("PRIMARY KEY"),
		tokens.LParen(),
		tokens.Raws(cols...),
		tokens.RParen(),
	)
}

type tCheck struct {
	expr Safe
}

func Check(expr Safe) ClauseConstraint {
	return tCheck{expr}
}

func (def tCheck) columnTokens() tokens.Tokens {
	return tokens.New(
		tokens.Keyword("CHECK"),
		tokens.LParen(),
		tokens.Raw(def.expr),
		tokens.RParen(),
	)
}

func (def tCheck) tableTokens(cols []Safe) tokens.Tokens {
	return def.columnTokens()
}
