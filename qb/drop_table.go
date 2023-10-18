package qb

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconfig"
)

type tDropTable struct {
	model    Model
	ifExists bool
}

func DropTable[T Model](model *T) tDropTable {
	return tDropTable{
		model: model,
	}
}

func (q tDropTable) IfExists() tDropTable {
	q.ifExists = true
	return q
}

func (q tDropTable) Squirrel(dbconfig.Config) (squirrel.Sqlizer, error) {
	tableName := getModelName(q.model)
	ifExists := ""
	if q.ifExists {
		ifExists = "IF EXISTS "
	}
	sql := fmt.Sprintf("DROP TABLE %s%s", ifExists, tableName)
	return squirrel.Expr(sql), nil
}
