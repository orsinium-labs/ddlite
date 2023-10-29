package dml

import (
	c "github.com/orsinium-labs/sequel/constraints"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

const precedenceAtomic = 255

// tFunc is a private type to represent stored function expression.
// `R` is the type of the function return value.
type tFunc[A, R any] struct {
	Name string
	Args []Expr[A]
}

// F is a stored function.
func F[A, T any](name string, args ...Expr[A]) Expr[T] {
	return tFunc[A, T]{Name: name, Args: args}
}

func (tFunc[A, R]) Precedence(dbconf.Config) uint8 {
	return precedenceAtomic
}

func (tFunc[A, R]) ExprType() R {
	return *new(R)
}

func (fn tFunc[A, R]) Tokens(conf dbconf.Config) tokens.Tokens {
	ts := tokens.New(
		tokens.FuncName(fn.Name),
		tokens.LParen(),
	)
	for i, arg := range fn.Args {
		if i > 0 {
			ts.Add(tokens.Comma())
		}
		ts.Extend(arg.Tokens(conf))
	}
	ts.Add(tokens.RParen())
	return ts
}

// tFunc is a private type to represent 2-argument stored function expression.
// `R` is the type of the function return value.
type tFunc2[A1, A2, R any] struct {
	Name string
	Arg1 Expr[A1]
	Arg2 Expr[A2]
}

// F is a stored function with 2 arguments of different type.
//
// For functions with any number of arguments of the same type
// prefer using `F` instead.
func F2[A1, A2, T any](name string, arg1 Expr[A1], arg2 Expr[A2]) Expr[T] {
	return tFunc2[A1, A2, T]{Name: name, Arg1: arg1, Arg2: arg2}
}

func (tFunc2[A1, A2, R]) Precedence(dbconf.Config) uint8 {
	return precedenceAtomic
}

func (tFunc2[A1, A2, R]) ExprType() R {
	return *new(R)
}

func (fn tFunc2[A1, A2, R]) Tokens(conf dbconf.Config) tokens.Tokens {
	ts := tokens.New(
		tokens.FuncName(fn.Name),
		tokens.LParen(),
	)
	ts.Extend(fn.Arg1.Tokens(conf))
	ts.Add(tokens.Comma())
	ts.Extend(fn.Arg2.Tokens(conf))
	ts.Add(tokens.RParen())
	return ts
}

// tCol is aprivate type to represent a column name expression.
type tCol[T any] struct {
	val any
}

// C is a column.
func C[T any](val *T) Expr[T] {
	return tCol[T]{val: val}
}

// M is a column wrapped into Option/Optional/Maybe monad.
func M[T any](val c.Option[T]) Expr[T] {
	return tCol[T]{val: val}
}

func (tCol[T]) Precedence(dbconf.Config) uint8 {
	return precedenceAtomic
}

func (tCol[T]) ExprType() T {
	return *new(T)
}

func (col tCol[T]) Tokens(conf dbconf.Config) tokens.Tokens {
	return tokens.New(internal.GetColumnName(conf, col.val))
}

// tVal is a private type to represent a literal value expression.
type tVal[T any] struct {
	val T
}

// V is a variable value.
//
// In the generated SQL, it will be represented as a bind parameter.
func V[T any](val T) Expr[T] {
	return tVal[T]{val: val}
}

func (tVal[T]) Precedence(dbconf.Config) uint8 {
	return precedenceAtomic
}

func (tVal[T]) ExprType() T {
	return *new(T)
}

func (val tVal[T]) Tokens(dbconf.Config) tokens.Tokens {
	return tokens.New(tokens.Bind(val.val))
}

// L is a literal (constant) value.
//
// Unlike [V], it will be added right into the SQL expression, without bind placeholders.
func L[T c.Number | bool](val T) Expr[T] {
	return tLiteral[T]{val: val}
}

type tLiteral[T any] struct {
	val any
}

func (tLiteral[T]) Precedence(dbconf.Config) uint8 {
	return precedenceAtomic
}

func (tLiteral[T]) ExprType() T {
	return *new(T)
}

func (expr tLiteral[T]) Tokens(dbconf.Config) tokens.Tokens {
	return tokens.New(tokens.Literal(expr.val))
}

type tCast[From, To any] struct {
	orig Expr[From]
}

func (expr tCast[From, To]) Precedence(c dbconf.Config) uint8 {
	return expr.orig.Precedence(c)
}

func (tCast[From, To]) ExprType() To {
	return *new(To)
}

func (expr tCast[From, To]) Tokens(c dbconf.Config) tokens.Tokens {
	return expr.orig.Tokens(c)
}

// Cast changes the type of the expression.
//
// It doesn't affect the generated SQL expression. It's an escape hatch
// from type safety when enforced types stand on your way.
// For example, when a function expects Expr[int] because it's an INTEGER
// in the database but you use a custom Decimal type to represent that number,
// you can use Cast to convert Expr[Decimal] into Expr[int].
func Cast[To, From any](expr Expr[From]) Expr[To] {
	return tCast[From, To]{expr}
}

type Safe string

func Raw[T any](raw Safe) Expr[T] {
	return tRaw[T]{raw}
}

type tRaw[T any] struct{ val Safe }

func (expr tRaw[T]) Precedence(c dbconf.Config) uint8 {
	return 0
}

func (tRaw[T]) ExprType() T {
	return *new(T)
}

func (expr tRaw[T]) Tokens(c dbconf.Config) tokens.Tokens {
	return tokens.New(tokens.Raw(expr.val))
}
