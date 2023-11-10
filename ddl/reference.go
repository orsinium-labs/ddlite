package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type Action string

const (
	SetNull    Action = "SET NULL"
	SetDefault Action = "SET DEFAULT"
	Cascade    Action = "CASCADE"
	Restrict   Action = "RESTRICT"
	NoAction   Action = "NO ACTION"
)

type Match string

const (
	Full    Match = "FULL"
	Partial Match = "PARTIAL"
	Simple  Match = "Simple"
)

type Reference struct {
	table    Safe
	columns  []Safe
	onDelete Action
	onUpdate Action
	match    Match
}

func References(table, column Safe, columns ...Safe) Reference {
	return Reference{
		table:   table,
		columns: append([]Safe{column}, columns...),
	}
}

func (r Reference) Match(m Match) Reference {
	r.match = m
	return r
}

func (r Reference) OnDelete(action Action) Reference {
	r.onDelete = action
	return r
}

func (r Reference) OnUpdate(action Action) Reference {
	r.onUpdate = action
	return r
}

func (r Reference) tokens(dialects.Dialect) tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("REFERENCES"),
		tokens.TableName(r.table),
		tokens.LParen(),
		tokens.Raws(r.columns...),
		tokens.RParen(),
	)
	if r.match != "" {
		ts.Add(tokens.Keyword("MATCH"))
		ts.Add(tokens.Raw(r.match))
	}
	if r.onDelete != "" {
		ts.Add(tokens.Keyword("ON DELETE"))
		ts.Add(tokens.Raw(r.onDelete))
	}
	if r.onUpdate != "" {
		ts.Add(tokens.Keyword("ON UPDATE"))
		ts.Add(tokens.Raw(r.onUpdate))
	}
	return ts
}
