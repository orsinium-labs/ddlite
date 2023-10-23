package ddl

import (
	"errors"
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
)

type tDropTable struct {
	table    Safe
	ifExists bool
}

func DropTable(table Safe) tDropTable {
	return tDropTable{
		table: table,
	}
}

func (q tDropTable) IfExists() tDropTable {
	q.ifExists = true
	return q
}

func (q tDropTable) SQL(conf dbconf.Config) (string, error) {
	if q.table == "" {
		return "", errors.New("table name must not be empty")
	}
	ifExists := ""
	if q.ifExists {
		ifExists = "IF EXISTS "
	}
	sql := fmt.Sprintf("DROP TABLE %s%s", ifExists, q.table)
	return sql, nil
}
