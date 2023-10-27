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
	precAnd := conf.Dialect.Precedence("AND")
	for i, pred := range c.predicates {
		if i > 0 {
			ts.Add(tokens.Keyword("AND"))
		}
		precPred := pred.Precedence(conf)
		addParens := precAnd == 0 || precPred < precAnd
		if addParens {
			ts.Add(tokens.LParen())
		}
		ts.Extend(pred.Tokens(conf))
		if addParens {
			ts.Add(tokens.RParen())
		}
	}
	return ts
}
