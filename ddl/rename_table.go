package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
)

type tRenameTable struct {
	old internal.Model
	new internal.Model
}

func RenameTable[T internal.Model](old, new *T) tRenameTable {
	return tRenameTable{old: old, new: new}
}

func (q tRenameTable) SQL(conf dbconf.Config) (string, error) {
	oldName := internal.GetTableName(conf, q.old)
	newName := internal.GetTableName(conf, q.new)
	sql := fmt.Sprintf("ALTER TABLE %s RENAME TO %s", oldName, newName)
	return sql, nil
}
