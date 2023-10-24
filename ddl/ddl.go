package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

// Safe is a string that is used in SQL queries as-is, without escaping.
//
// String literals and constants are automatically considered safe.
// Variables need to be explicitly converted to Safe.
//
// Never convert to Safe untrusted input, it allows evil people to do SQL injections.
type Safe string

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
	return sql, args, nil
}
