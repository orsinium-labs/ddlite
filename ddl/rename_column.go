package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
)

type tRenameColumn struct {
	table string
	old   string
	new   string
}

func RenameColumn(table string, old, new string) tRenameColumn {
	return tRenameColumn{table: table, old: old, new: new}
}

func (q tRenameColumn) SQL(conf dbconf.Config) (string, error) {
	sql := fmt.Sprintf("ALTER TABLE %s RENAME COLUMN %s TO %s", q.table, q.old, q.new)
	return sql, nil
}
