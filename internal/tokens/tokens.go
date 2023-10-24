package tokens

import (
	"strings"

	"github.com/orsinium-labs/sequel/dbconf"
)

type Token interface {
	sql(dbconf.Config) (string, []any, error)
}

func New(ts ...Token) Tokens {
	return Tokens{ts}
}

type Tokens struct {
	tokens []Token
}

func (tokens *Tokens) Add(ts ...Token) {
	tokens.tokens = append(tokens.tokens, ts...)
}

func (tokens *Tokens) Extend(ts Tokens) {
	tokens.tokens = append(tokens.tokens, ts.tokens...)
}

func (tokens Tokens) SQL(conf dbconf.Config) (string, []any, error) {
	result := strings.Builder{}
	args := make([]any, 0)
	for i, token := range tokens.tokens {
		sql, subArgs, err := token.sql(conf)
		if err != nil {
			return "", nil, err
		}
		result.WriteString(sql)
		if tokens.needsSpace(i) {
			result.WriteString(" ")
		}
		args = append(args, subArgs...)
	}
	return result.String(), args, nil
}

// needsSpace checks if space should be added after the token with the given index.
func (tokens Tokens) needsSpace(i int) bool {
	curr := tokens.tokens[i]
	switch curr.(type) {
	case tFuncName:
		return false
	case tLParen:
		return false
	}

	if i >= len(tokens.tokens)-1 {
		return false
	}
	next := tokens.tokens[i+1]
	switch next.(type) {
	case tComma:
		return false
	case tRParen:
		return false
	}
	return true
}

// Raw SQL string
func Raw[T ~string](s T) Token {
	return tRaw(s)
}

func TableName[T ~string](s T) Token {
	return tRaw(s)
}

func ColumnName[T ~string](s T) Token {
	return tRaw(s)
}

func FuncName[T ~string](s T) Token {
	return tFuncName(s)
}

type tFuncName string

func (token tFuncName) sql(dbconf.Config) (string, []any, error) {
	return string(token), nil, nil
}

func Keyword(s string) Token {
	return tRaw(s)
}

func Operator(s string) Token {
	return tRaw(s)
}

// Left parenthesis
func LParen() Token {
	return tLParen{}
}

type tLParen struct{}

func (token tLParen) sql(dbconf.Config) (string, []any, error) {
	return "(", nil, nil
}

// Right parenthesis
func RParen() Token {
	return tRParen{}
}

type tRParen struct{}

func (token tRParen) sql(dbconf.Config) (string, []any, error) {
	return ")", nil, nil
}

func Comma() Token {
	return tComma{}
}

type tComma struct{}

func (token tComma) sql(dbconf.Config) (string, []any, error) {
	return ",", nil, nil
}

type tRaw string

func (token tRaw) sql(dbconf.Config) (string, []any, error) {
	return string(token), nil, nil
}

// List of raw SQL values
func Raws[T ~string](ss ...T) Token {
	res := make(tList, 0, len(ss))
	for _, s := range ss {
		res = append(res, string(s))
	}
	return res
}

type tList []string

func (token tList) sql(dbconf.Config) (string, []any, error) {
	return strings.Join(token, ", "), nil, nil
}

func Bind(val any) Token {
	return tBind{val}
}

type tBind struct{ val any }

func (token tBind) sql(dbconf.Config) (string, []any, error) {
	return "?", []any{token.val}, nil
}
