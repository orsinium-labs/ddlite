package qb

import (
	"fmt"

	"github.com/Masterminds/squirrel"
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

func (q tDropTable) Squirrel(dbconf.Config) (squirrel.Sqlizer, error) {
	tableName := internal.GetModelName(q.model)
	ifExists := ""
	if q.ifExists {
		ifExists = "IF EXISTS "
	}
	sql := fmt.Sprintf("DROP TABLE %s%s", ifExists, tableName)
	return squirrel.Expr(sql), nil
}
