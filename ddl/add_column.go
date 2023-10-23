package ddl

import (
	"errors"
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
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
	sql := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", q.table, columnSQL)
	return sql, nil
}
