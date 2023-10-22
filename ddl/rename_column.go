package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
)

type tRenameColumn[T any] struct {
	model any
	old   *T
	new   *T
}

func RenameColumn[M, C any](model *M, old, new *C) tRenameColumn[C] {
	return tRenameColumn[C]{model: model, old: old, new: new}
}

func (q tRenameColumn[C]) SQL(conf dbconf.Config) (string, error) {
	tableName := internal.GetTableName(conf, q.model)
	oldName, err := internal.GetColumnName(conf, q.old)
	if err != nil {
		return "", fmt.Errorf("get old column name: %v", err)
	}
	newName, err := internal.GetColumnName(conf, q.new)
	if err != nil {
		return "", fmt.Errorf("get new column name: %v", err)
	}
	sql := fmt.Sprintf("ALTER TABLE %s RENAME COLUMN %s TO %s", tableName, oldName, newName)
	return sql, nil
}
