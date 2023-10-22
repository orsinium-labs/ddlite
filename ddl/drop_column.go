package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
)

type tDropColumn[T any] struct {
	model any
	col   *T
}

func DropColumn[M, C any](model *M, col *C) tDropColumn[C] {
	return tDropColumn[C]{model: model, col: col}
}

func (q tDropColumn[C]) SQL(conf dbconf.Config) (string, error) {
	tableName := internal.GetTableName(conf, q.model)
	colName, err := internal.GetColumnName(conf, q.col)
	if err != nil {
		return "", fmt.Errorf("get column name: %v", err)
	}
	sql := fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", tableName, colName)
	return sql, nil
}
