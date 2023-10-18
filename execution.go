package sequel

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconfig"
	"github.com/orsinium-labs/sequel/qb"
)

type query interface {
	Squirrel(dbconfig.Config) (squirrel.Sqlizer, error)
}

type scannableQuery[T qb.Model] interface {
	query
	Scanner(dbconfig.Config, *T) (qb.Scanner[T], error)
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

// SQL converts the given expression to SQL.
func SQL(conf dbconfig.Config, q query) (string, []any, error) {
	builder, err := q.Squirrel(conf)
	if err != nil {
		return "", nil, fmt.Errorf("build query: %v", err)
	}
	sql, arg, err := builder.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("convert query to SQL: %v", err)
	}
	return sql, arg, nil
}

func Exec(
	conf dbconfig.Config,
	db dbOrTx,
	q query,
) (sql.Result, error) {
	builder, err := q.Squirrel(conf)
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}
	// driver.DriverName()
	sqlQ, args, err := builder.ToSql()
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
func FetchOne[T qb.Model](
	conf dbconfig.Config,
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
func FetchOneInto[T qb.Model](
	conf dbconfig.Config,
	db dbOrTx,
	q scannableQuery[T],
	target *T,
) error {
	builder, err := q.Squirrel(conf)
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}
	sqlQ, args, err := builder.ToSql()
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
