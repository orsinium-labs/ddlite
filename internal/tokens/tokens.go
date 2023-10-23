package tokens

import (
	"strings"

	"github.com/orsinium-labs/sequel/dbconf"
)

type Token interface {
	SQL(dbconf.Config) (string, error)
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

func (tokens Tokens) SQL(conf dbconf.Config) (string, error) {
	parts := make([]string, 0, len(tokens.tokens))
	for _, token := range tokens.tokens {
		sql, err := token.SQL(conf)
		if err != nil {
			return "", err
		}
		parts = append(parts, sql)
	}
	return strings.Join(parts, " "), nil
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

func Keyword(s string) Token {
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

type tRaw string

func (token tRaw) SQL(dbconf.Config) (string, error) {
	return string(token), nil
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

func (token tList) SQL(dbconf.Config) (string, error) {
	return strings.Join(token, ", "), nil
}
