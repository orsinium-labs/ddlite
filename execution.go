package sequel

import (
	"database/sql"
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type query interface {
	Tokens(dbconf.Config) tokens.Tokens
}

type dbOrTx interface {
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
}

// Must wraps a function call returning a value and an error and panics if the error is not nil.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// SQL generates SQL string for the given sequel query.
func SQL(conf dbconf.Config, query query) (string, error) {
	sql, err := query.Tokens(conf).SQL()
	if err != nil {
		return "", fmt.Errorf("convert tokens to SQL: %w", err)
	}
	if err != nil {
		return "", fmt.Errorf("convert placeholders: %w", err)
	}
	return sql, nil
}

func Exec(
	conf dbconf.Config,
	db dbOrTx,
	q query,
) (sql.Result, error) {
	sqlQ, err := SQL(conf, q)
	if err != nil {
		return nil, fmt.Errorf("generate SQL query: %w", err)
	}
	r, err := db.Exec(sqlQ)
	if err != nil {
		return nil, fmt.Errorf("execute SQL query: %w", err)
	}
	return r, nil
}
