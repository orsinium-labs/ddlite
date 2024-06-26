package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

// ClauseTableConstraint is a contraint applied to the table (a set of columns).
//
// Constructed by [Constraint].
type ClauseTableConstraint struct {
	name       string
	constraint ClauseConstraint
	columns    []Safe
}

// Constraint is a table constraint applied to a set of fields.
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
		ts.Add(tokens.Keyword("CONSTRAINT"))
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

// Unique requires each value in the column to be unique.
//
// SQL: UNIQUE
//
// https://www.sqlite.org/lang_createtable.html#unique_constraints
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

// PrimaryKey marks a column as the primary key for the table.
//
// A table may have only one primary key constraint but that constraint
// may include multiple columns.
//
// SQL: PRIMARY KEY
//
// https://www.sqlite.org/lang_createtable.html#the_primary_key
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
