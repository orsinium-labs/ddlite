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
	Scanner() (qb.Scanner[T], error)
}
type dbOrTx interface {
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
}

func SQL(c dbconfig.Config, q query) (string, []any, error) {
	builder, err := q.Squirrel(c)
	if err != nil {
		return "", nil, fmt.Errorf("build query: %v", err)
	}
	sql, arg, err := builder.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("convert query to SQL: %v", err)
	}
	return sql, arg, nil
}

func Exec(c dbconfig.Config, db dbOrTx, q query) (sql.Result, error) {
	builder, err := q.Squirrel(c)
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

func FetchOne[T qb.Model](c dbconfig.Config, db dbOrTx, q scannableQuery[T]) (T, error) {
	var r T
	builder, err := q.Squirrel(c)
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
