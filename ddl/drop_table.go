package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
)

type tDropTable struct {
	table    string
	ifExists bool
}

func DropTable(table string) tDropTable {
	return tDropTable{
		table: table,
	}
}

func (q tDropTable) IfExists() tDropTable {
	q.ifExists = true
	return q
}

func (q tDropTable) SQL(conf dbconf.Config) (string, error) {
	ifExists := ""
	if q.ifExists {
		ifExists = "IF EXISTS "
	}
	sql := fmt.Sprintf("DROP TABLE %s%s", ifExists, q.table)
	return sql, nil
}
