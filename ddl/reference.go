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

type ClauseReferences struct {
	table    Safe
	columns  []Safe
	onDelete Action
	onUpdate Action
	match    Match
}

var _ ClauseConstraint = ClauseReferences{}

func References(table, column Safe, columns ...Safe) ClauseReferences {
	return ClauseReferences{
		table:   table,
		columns: append([]Safe{column}, columns...),
	}
}

func (r ClauseReferences) Match(m Match) ClauseReferences {
	r.match = m
	return r
}

func (r ClauseReferences) OnDelete(action Action) ClauseReferences {
	r.onDelete = action
	return r
}

func (r ClauseReferences) OnUpdate(action Action) ClauseReferences {
	r.onUpdate = action
	return r
}

func (r ClauseReferences) tableTokens(d dialects.Dialect, cols []Safe) tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("FOREIGN KEY"),
		tokens.LParen(),
		tokens.Raws(cols...),
		tokens.RParen(),
	)
	ts.Extend(r.columnTokens(d))
	return ts
}

func (r ClauseReferences) columnTokens(dialects.Dialect) tokens.Tokens {
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
