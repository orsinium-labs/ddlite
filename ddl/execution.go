package ddl

import (
	"database/sql"
	"fmt"

	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type Statement interface {
	tokens(dialects.Dialect) tokens.Tokens
}

type executor interface {
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
func SQL(dialect dialects.Dialect, stmt Statement) (string, error) {
	sql, err := stmt.tokens(dialect).SQL()
	if err != nil {
		return "", fmt.Errorf("convert tokens to SQL: %w", err)
	}
	if err != nil {
		return "", fmt.Errorf("convert placeholders: %w", err)
	}
	return sql, nil
}

func Exec(dialect dialects.Dialect, db executor, stmt Statement) (sql.Result, error) {
	sqlQ, err := SQL(dialect, stmt)
	if err != nil {
		return nil, fmt.Errorf("generate SQL query: %w", err)
	}
	r, err := db.Exec(sqlQ)
	if err != nil {
		return nil, fmt.Errorf("execute SQL query: %w", err)
	}
	return r, nil
}
