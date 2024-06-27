package ddl

import (
	"database/sql"
	"fmt"

	"github.com/orsinium-labs/ddlite/internal/tokens"
)

type Statement interface {
	tokens() tokens.Tokens
}

// executor is an interface describing a DB or Tx from sql or sqlx package.
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

// SQL generates SQL string for the given [Statement].
func SQL(stmt Statement) (string, error) {
	sql, err := stmt.tokens().SQL()
	if err != nil {
		return "", fmt.Errorf("convert tokens to SQL: %w", err)
	}
	return sql, nil
}

// Exec generates SQL for the given [Statement] and executes it.
func Exec(db executor, stmt Statement) (sql.Result, error) {
	sqlQ, err := SQL(stmt)
	if err != nil {
		return nil, fmt.Errorf("generate SQL query: %w", err)
	}
	r, err := db.Exec(sqlQ)
	if err != nil {
		return nil, fmt.Errorf("execute SQL query: %w", err)
	}
	return r, nil
}
