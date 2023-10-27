package dml

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type whereClause struct {
	predicates []Expr[bool]
}

func (c whereClause) buildWhere(conf dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	if len(c.predicates) == 0 {
		return ts
	}
	ts.Add(tokens.Keyword("WHERE"))
	for i, pred := range c.predicates {
		if i > 0 {
			ts.Add(tokens.Keyword("AND"))
		}
		ts.Extend(pred.Tokens(conf))
	}
	return ts
}
