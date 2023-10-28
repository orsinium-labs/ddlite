package dml

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

// Direction is the order in which the values will be sorted when using [Ordering].
type Direction bool

const (
	// Asc is ascending order (large goes last).
	Asc Direction = false

	// Desc is descending order (large goes first).
	Desc Direction = true
)

// Nulls is used by [OrderingBuilder] to specify if nulls should go first or last.
type Nulls uint8

const (
	// NullsFirst (NULLS FIRST) puts NULL values before any other values.
	NullsFirst Nulls = 1

	// NullsLast (NULLS LAST) puts NULL values after any other values.
	NullsLast Nulls = 2
)

func Ordering[T any](expr Expr[T], dir Direction) OrderingBuilder {
	return OrderingBuilder{Cast[any](expr), dir, 0}
}

// OrderingBuilder specifies values order for ORDER BY clause.
//
// It must be constructed using the [Ordering] function.
type OrderingBuilder struct {
	expr      Expr[any]
	direction Direction
	nulls     Nulls
}

// Nulls specifies if nulls should go first or last.
//
// If you don't call the method, the default behavior is to consider NULLs
// the largest values. That is, the will go first if [Direction] is [Desc]
// and will go last if the [Direction] is [Asc].
func (ord OrderingBuilder) Nulls(nulls Nulls) OrderingBuilder {
	ord.nulls = nulls
	return ord
}

func (ord OrderingBuilder) Tokens(c dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	ts.Extend(ord.expr.Tokens(c))
	if ord.direction == Asc {
		ts.Add(tokens.Keyword("ASC"))
	} else {
		ts.Add(tokens.Keyword("DESC"))
	}
	switch ord.nulls {
	case NullsFirst:
		ts.Add(tokens.Keyword("NULLS FIRST"))
	case NullsLast:
		ts.Add(tokens.Keyword("NULLS LAST"))
	}
	return ts
}
