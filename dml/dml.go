package dml

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tokener interface {
	Tokens(dbconf.Config) (tokens.Tokens, error)
}

// SQL generates SQL string for the given sequel query.
func SQL(conf dbconf.Config, query tokener) (string, []any, error) {
	ts, err := query.Tokens(conf)
	if err != nil {
		return "", nil, fmt.Errorf("generate tokens: %w", err)
	}
	sql, args, err := ts.SQL(conf)
	if err != nil {
		return "", nil, fmt.Errorf("convert tokens to SQL: %w", err)
	}
	sql, err = conf.SquirrelPlaceholder().ReplacePlaceholders(sql)
	if err != nil {
		return "", nil, fmt.Errorf("convert placeholders: %w", err)
	}
	return sql, args, nil
}
