package ddl

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
)

type tCreateTable struct {
	model internal.Model
	cols  []iColumnDef
}

func CreateTable[T internal.Model](model *T, cols ...iColumnDef) tCreateTable {
	return tCreateTable{
		model: model,
		cols:  cols,
	}
}

func (q tCreateTable) Squirrel(conf dbconf.Config) (squirrel.Sqlizer, error) {
	return internal.SQL2Squirrel(conf, q)
}

func (q tCreateTable) SQL(conf dbconf.Config) (string, error) {
	conf = conf.WithModel(q.model)
	colNames := make([]string, 0, len(q.cols))
	for _, col := range q.cols {
		csql, err := col.SQL(conf)
		if err != nil {
			return "", fmt.Errorf("generate SQL for ColumnDef: %v", err)
		}
		colNames = append(colNames, csql)
	}
	tableName := internal.GetTableName(conf, q.model)
	cdefs := strings.Join(colNames, ", ")
	sql := fmt.Sprintf("CREATE TABLE %s (%s)", tableName, cdefs)
	return sql, nil
}
