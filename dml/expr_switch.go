package dml

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/priority"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type SwitchBuilder[T, R any] struct {
	val     Expr[T]
	conds   []Expr[T]
	results []Expr[R]
	def     Expr[R]
}

func Switch[R any]() SwitchBuilder[bool, R] {
	return SwitchBuilder[bool, R]{}
}

func SwitchBy[R, T any](val Expr[T]) SwitchBuilder[T, R] {
	return SwitchBuilder[T, R]{val: val}
}

func (s SwitchBuilder[T, R]) Case(cond Expr[T], res Expr[R]) SwitchBuilder[T, R] {
	s.conds = append(s.conds, cond)
	s.results = append(s.results, res)
	return s
}

func (s SwitchBuilder[T, R]) Else(res Expr[R]) SwitchBuilder[T, R] {
	s.def = res
	return s
}

func (SwitchBuilder[T, R]) Priority() priority.Priority {
	return priority.Atomic
}

func (SwitchBuilder[T, R]) ExprType() R {
	return *new(R)
}

func (s SwitchBuilder[T, R]) Tokens(c dbconf.Config) tokens.Tokens {
	ts := tokens.New(tokens.Keyword("CASE"))
	if s.val != nil {
		ts.Extend(s.val.Tokens(c))
	}
	for i, cond := range s.conds {
		res := s.results[i]
		ts.Add(tokens.Keyword("WHEN"))
		ts.Extend(cond.Tokens(c))
		ts.Add(tokens.Keyword("THEN"))
		ts.Extend(res.Tokens(c))
	}
	if s.def != nil {
		ts.Add(tokens.Keyword("ELSE"))
		ts.Extend(s.def.Tokens(c))
	}
	ts.Add(tokens.Keyword("END"))
	return ts
}
