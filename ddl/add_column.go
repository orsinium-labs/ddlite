package ddl

import (
	"errors"
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tAddColumn struct {
	table Safe
	col   tColumn
}

// AddColumn builds query that adds a new column to the table.
func AddColumn(table Safe, col tColumn) tAddColumn {
	return tAddColumn{table: table, col: col}
}

func (q tAddColumn) Tokens(conf dbconf.Config) (tokens.Tokens, error) {
	if q.table == "" {
		return tokens.New(), errors.New("table name must not be empty")
	}
	colTokens, err := q.col.Tokens(conf)
	if err != nil {
		return tokens.New(), fmt.Errorf("generate SQL for ColumnDef: %v", err)
	}
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.table),
		tokens.Keyword("ADD COLUMN"),
	)
	ts.Extend(colTokens)
	return ts, nil
}
