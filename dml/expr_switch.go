package dml

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tSwitch[T, R any] struct {
	val     Expr[T]
	conds   []Expr[T]
	results []Expr[R]
	def     Expr[R]
}

func Switch[R any]() tSwitch[bool, R] {
	return tSwitch[bool, R]{}
}

func SwitchBy[R, T any](val Expr[T]) tSwitch[T, R] {
	return tSwitch[T, R]{val: val}
}

func (s tSwitch[T, R]) Case(cond Expr[T], res Expr[R]) tSwitch[T, R] {
	s.conds = append(s.conds, cond)
	s.results = append(s.results, res)
	return s
}

func (s tSwitch[T, R]) Else(res Expr[R]) tSwitch[T, R] {
	s.def = res
	return s
}

func (tSwitch[T, R]) ExprType() R {
	return *new(R)
}

func (s tSwitch[T, R]) Tokens(c dbconf.Config) tokens.Tokens {
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
