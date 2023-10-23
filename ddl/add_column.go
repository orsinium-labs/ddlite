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

func (q tAddColumn) SQL(conf dbconf.Config) (string, error) {
	if q.table == "" {
		return "", errors.New("table name must not be empty")
	}
	columnSQL, err := q.col.SQL(conf)
	if err != nil {
		return "", fmt.Errorf("generate SQL for ColumnDef: %v", err)
	}
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.table),
		tokens.Keyword("ADD COLUMN"),
		tokens.Raw(columnSQL),
	)
	return ts.SQL(conf)
}
