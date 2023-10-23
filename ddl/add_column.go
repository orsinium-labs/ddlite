package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
)

type tAddColumn struct {
	model any
	col   tColumnDef
}

// AddColumn builds query that adds a new column to the table.
func AddColumn[M, C any](model *M, col tColumnDef) tAddColumn {
	return tAddColumn{model: model, col: col}
}

func (q tAddColumn) SQL(conf dbconf.Config) (string, error) {
	tableName := internal.GetTableName(conf, q.model)
	columnSQL, err := q.col.SQL(conf)
	if err != nil {
		return "", fmt.Errorf("generate SQL for ColumnDef: %v", err)
	}
	sql := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", tableName, columnSQL)
	return sql, nil
}
