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
	for _, token := range tokens.tokens {
		sql, subArgs, err := token.sql(conf)
		if err != nil {
			return "", nil, err
		}
		result.WriteString(sql)
		_, isFunc := token.(tFuncName)
		if !isFunc {
			result.WriteString(" ")
		}
		args = append(args, subArgs...)
	}
	return strings.TrimRight(result.String(), " "), args, nil
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
	return tRaw("(")
}

// Right parenthesis
func RParen() Token {
	return tRaw(")")
}

func Comma() Token {
	return tRaw(",")
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
