package ddl

import (
	"errors"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tDropColumn struct {
	table Safe
	col   Safe
}

func DropColumn(table Safe, col Safe) tDropColumn {
	return tDropColumn{table: table, col: col}
}

func (q tDropColumn) Tokens(conf dbconf.Config) (tokens.Tokens, error) {
	if q.table == "" {
		return tokens.New(), errors.New("table name must not be empty")
	}
	if q.col == "" {
		return tokens.New(), errors.New("column name must not be empty")
	}
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.table),
		tokens.Keyword("DROP COLUMN"),
		tokens.ColumnName(q.col),
	)
	return ts, nil
}
