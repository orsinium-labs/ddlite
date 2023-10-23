package ddl

import (
	"errors"
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
)

type tRenameTable struct {
	old Safe
	new Safe
}

func RenameTable(old, new Safe) tRenameTable {
	return tRenameTable{old: old, new: new}
}

func (q tRenameTable) SQL(conf dbconf.Config) (string, error) {
	if q.old == "" {
		return "", errors.New("old table name must not be empty")
	}
	if q.new == "" {
		return "", errors.New("new table name must not be empty")
	}
	sql := fmt.Sprintf("ALTER TABLE %s RENAME TO %s", q.old, q.new)
	return sql, nil
}
