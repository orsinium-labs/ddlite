package ddl

import (
	"errors"
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
)

type tRenameColumn struct {
	table Safe
	old   Safe
	new   Safe
}

func RenameColumn(table, old, new Safe) tRenameColumn {
	return tRenameColumn{table: table, old: old, new: new}
}

func (q tRenameColumn) SQL(conf dbconf.Config) (string, error) {
	if q.table == "" {
		return "", errors.New("table name must not be empty")
	}
	if q.old == "" {
		return "", errors.New("old column name must not be empty")
	}
	if q.new == "" {
		return "", errors.New("new column name must not be empty")
	}
	sql := fmt.Sprintf("ALTER TABLE %s RENAME COLUMN %s TO %s", q.table, q.old, q.new)
	return sql, nil
}
