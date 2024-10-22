package tokens

import (
	"errors"
	"strings"
)

type Token interface {
	sql() (string, error)
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

func (tokens Tokens) SQL() (string, error) {
	result := strings.Builder{}
	for i, token := range tokens.tokens {
		sql, err := token.sql()
		if err != nil {
			return "", err
		}
		result.WriteString(sql)
		if tokens.needsSpace(i) {
			result.WriteString(" ")
		}
	}
	return result.String(), nil
}

// needsSpace checks if space should be added after the token with the given index.
func (tokens Tokens) needsSpace(i int) bool {
	curr := tokens.tokens[i]
	switch curr.(type) {
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

func Err(err error) Token {
	return tErr{err}
}

type tErr struct{ err error }

func (token tErr) sql() (string, error) {
	return "", token.err
}

// Raw SQL string
func Raw[T ~string](s T) Token {
	return tRaw(s)
}

func TableName[T ~string](s T) Token {
	if s == "" {
		return tErr{errors.New("table name must not be empty")}
	}
	return tRaw(s)
}

func ColumnName[T ~string](s T) Token {
	if s == "" {
		return tErr{errors.New("column name must not be empty")}
	}
	return tRaw(s)
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

func (token tLParen) sql() (string, error) {
	return "(", nil
}

// Right parenthesis
func RParen() Token {
	return tRParen{}
}

type tRParen struct{}

func (token tRParen) sql() (string, error) {
	return ")", nil
}

func Comma() Token {
	return tComma{}
}

type tComma struct{}

func (token tComma) sql() (string, error) {
	return ",", nil
}

type tRaw string

func (token tRaw) sql() (string, error) {
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

func (token tList) sql() (string, error) {
	return strings.Join(token, ", "), nil
}
