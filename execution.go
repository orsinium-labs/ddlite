package sequel

import (
	"database/sql"
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dml"
	"github.com/orsinium-labs/sequel/internal"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type query interface {
	Tokens(dbconf.Config) (tokens.Tokens, error)
}

type scannableQuery[T internal.Model] interface {
	query
	Scanner(dbconf.Config, *T) (dml.Scanner[T], error)
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
func SQL(conf dbconf.Config, query query) (string, []any, error) {
	ts, err := query.Tokens(conf)
	if err != nil {
		return "", nil, fmt.Errorf("generate tokens: %w", err)
	}
	sql, args, err := ts.SQL(conf)
	if err != nil {
		return "", nil, fmt.Errorf("convert tokens to SQL: %w", err)
	}
	if err != nil {
		return "", nil, fmt.Errorf("convert placeholders: %w", err)
	}
	return sql, args, nil
}

func Exec(
	conf dbconf.Config,
	db dbOrTx,
	q query,
) (sql.Result, error) {
	sqlQ, args, err := SQL(conf, q)
	if err != nil {
		return nil, fmt.Errorf("generate SQL query: %w", err)
	}
	r, err := db.Exec(sqlQ, args...)
	if err != nil {
		return nil, fmt.Errorf("execute SQL query: %w", err)
	}
	return r, nil
}

// FetchOne runs the query and the result as a struct.
//
// The query is expected to return exactly one record.
func FetchOne[T internal.Model](
	conf dbconf.Config,
	db dbOrTx,
	q scannableQuery[T],
) (T, error) {
	var result T
	err := FetchOneInto(conf, db, q, &result)
	return result, err
}

// FetchOneInto runs the query and places the result into the given struct.
//
// The query is expected to return exactly one record.
func FetchOneInto[T internal.Model](
	conf dbconf.Config,
	db dbOrTx,
	q scannableQuery[T],
	target *T,
) error {
	sqlQ, args, err := SQL(conf, q)
	if err != nil {
		return fmt.Errorf("generate SQL query: %w", err)
	}
	scanner, err := q.Scanner(conf, target)
	if err != nil {
		return fmt.Errorf("make scanner for the query: %w", err)
	}

	rows, err := db.Query(sqlQ, args...)
	if err != nil {
		return fmt.Errorf("run query: %w", err)
	}
	defer rows.Close()

	ok := rows.Next()
	if !ok {
		return fmt.Errorf("query returned no results")
	}
	err = scanner(rows)
	if err != nil {
		return fmt.Errorf("run scanner: %w", err)
	}
	return nil
}
