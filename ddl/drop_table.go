package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
)

type tDropTable struct {
	model    internal.Model
	ifExists bool
}

func DropTable[T internal.Model](model *T) tDropTable {
	return tDropTable{
		model: model,
	}
}

func (q tDropTable) IfExists() tDropTable {
	q.ifExists = true
	return q
}

func (q tDropTable) SQL(conf dbconf.Config) (string, error) {
	tableName := internal.GetTableName(conf, q.model)
	ifExists := ""
	if q.ifExists {
		ifExists = "IF EXISTS "
	}
	sql := fmt.Sprintf("DROP TABLE %s%s", ifExists, tableName)
	return sql, nil
}
