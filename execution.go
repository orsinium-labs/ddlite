package qb

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type query interface {
	Squirrel(...Model) (squirrel.Sqlizer, error)
}

type Scanner[T Model] func(rows *sql.Rows) (T, error)
type scannableQuery[T Model] interface {
	query
	Scanner() (Scanner[T], error)
}
type dbOrTx interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

func SQL(q query) (string, []any, error) {
	builder, err := q.Squirrel()
	if err != nil {
		return "", nil, fmt.Errorf("build query: %v", err)
	}
	sql, arg, err := builder.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("convert query to SQL: %v", err)
	}
	return sql, arg, nil
}

func FetchOne[T Model](db dbOrTx, q scannableQuery[T]) (T, error) {
	var r T
	builder, err := q.Squirrel()
	if err != nil {
		return r, fmt.Errorf("build query: %w", err)
	}
	sqlQ, args, err := builder.ToSql()
	if err != nil {
		return r, fmt.Errorf("generate SQL query: %w", err)
	}
	scanner, err := q.Scanner()
	if err != nil {
		return r, fmt.Errorf("make scanner for the query: %w", err)
	}

	rows, err := db.Query(sqlQ, args...)
	if err != nil {
		return r, fmt.Errorf("run query: %w", err)
	}
	defer rows.Close()

	ok := rows.Next()
	if !ok {
		return r, fmt.Errorf("query returned no results")
	}
	r, err = scanner(rows)
	if err != nil {
		return r, fmt.Errorf("run scanner: %w", err)
	}
	return r, nil
}
