package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
)

type tDropColumn struct {
	table Safe
	col   Safe
}

func DropColumn(table Safe, col Safe) tDropColumn {
	return tDropColumn{table: table, col: col}
}

func (q tDropColumn) SQL(conf dbconf.Config) (string, error) {
	sql := fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", q.table, q.col)
	return sql, nil
}
