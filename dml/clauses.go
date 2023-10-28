package dml

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type where struct {
	predicates []Expr[bool]
}

func (c where) build(conf dbconf.Config) tokens.Tokens {
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

type limit struct {
	limit  Expr[int]
	offset Expr[int]
}

func (c limit) build(conf dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	if c.limit != nil {
		ts.Add(tokens.Keyword("LIMIT"))
		ts.Extend(c.limit.Tokens(conf))
	}
	if c.offset != nil {
		ts.Add(tokens.Keyword("OFFSET"))
		ts.Extend(c.offset.Tokens(conf))
	}
	return ts

}

type order struct {
	ords []iOrdering
}

func (c order) build(conf dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	if len(c.ords) == 0 {
		return ts
	}
	ts.Add(tokens.Keyword("ORDER BY"))
	for i, ord := range c.ords {
		if i > 0 {
			ts.Add(tokens.Comma())
		}
		ts.Extend(ord.Tokens(conf))
	}
	return ts

}
