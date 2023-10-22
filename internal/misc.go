package internal

import (
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconf"
)

type squirrler interface {
	Squirrel(dbconf.Config) (squirrel.Sqlizer, error)
}

type sqler interface {
	SQL(dbconf.Config) (string, error)
}

func SQL2Squirrel(conf dbconf.Config, queryBuilder sqler) (squirrel.Sqlizer, error) {
	sql, err := queryBuilder.SQL(conf)
	if err != nil {
		return nil, fmt.Errorf("generate SQL: %w", err)
	}
	return squirrel.Expr(sql), nil
}

func Squirrel2SQL(conf dbconf.Config, queryBuilder squirrler) (string, error) {
	qb, err := queryBuilder.Squirrel(conf)
	if err != nil {
		return "", fmt.Errorf("build squirrel query: %w", err)
	}
	sql, args, err := qb.ToSql()
	if err != nil {
		return "", fmt.Errorf("generate SQL: %w", err)
	}
	if len(args) != 0 {
		return "", errors.New("bind args are not supported for the query")
	}
	return sql, nil
}
