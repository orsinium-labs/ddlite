package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
)

type tRenameTable struct {
	old string
	new string
}

func RenameTable(old, new string) tRenameTable {
	return tRenameTable{old: old, new: new}
}

func (q tRenameTable) SQL(conf dbconf.Config) (string, error) {
	sql := fmt.Sprintf("ALTER TABLE %s RENAME TO %s", q.old, q.new)
	return sql, nil
}
