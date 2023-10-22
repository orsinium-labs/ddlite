package ddl

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
)

type tAddColumn[C any] struct {
	model any
	col   tColumnDef[C]
}

// AddColumn builds query that adds a new column to the table.
func AddColumn[M, C any](model *M, col tColumnDef[C]) tAddColumn[C] {
	return tAddColumn[C]{model: model, col: col}
}

func (q tAddColumn[C]) Squirrel(conf dbconf.Config) (squirrel.Sqlizer, error) {
	tableName := internal.GetTableName(conf, q.model)
	columnSQL, err := q.col.SQL(conf)
	if err != nil {
		return nil, fmt.Errorf("generate SQL for ColumnDef: %v", err)
	}
	sql := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", tableName, columnSQL)
	return squirrel.Expr(sql), nil
}
