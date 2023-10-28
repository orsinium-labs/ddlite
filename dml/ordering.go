package dml

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type Direction bool

const (
	Asc  Direction = false
	Desc Direction = true
)

type nulls uint8

const (
	nullsFirst nulls = 1
	nullsLast  nulls = 2
)

func Ordering[T any](expr Expr[T], dir Direction) tOrdering {
	return tOrdering{Cast[any](expr), dir, 0}
}

type iOrdering interface {
	Ordering()
	Tokens(dbconf.Config) tokens.Tokens
}

type tOrdering struct {
	expr      Expr[any]
	direction Direction
	nulls     nulls
}

func (ord tOrdering) NullsFirst() iOrdering {
	ord.nulls = nullsFirst
	return ord
}

func (ord tOrdering) NullsLast() iOrdering {
	ord.nulls = nullsLast
	return ord
}

// Ordering implements the [iOrdering] interface.
func (ord tOrdering) Ordering() {}

// Tokens implements the [iOrdering] interface.
func (ord tOrdering) Tokens(c dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	ts.Extend(ord.expr.Tokens(c))
	if ord.direction == Asc {
		ts.Add(tokens.Keyword("ASC"))
	} else {
		ts.Add(tokens.Keyword("DESC"))
	}
	switch ord.nulls {
	case nullsFirst:
		ts.Add(tokens.Keyword("NULLS FIRST"))
	case nullsLast:
		ts.Add(tokens.Keyword("NULLS LAST"))
	}
	return ts
}
