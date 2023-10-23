package ddl

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
)

type tCreateTable struct {
	table string
	cols  []iColumnDef
}

func CreateTable(table string, cols ...iColumnDef) tCreateTable {
	return tCreateTable{
		table: table,
		cols:  cols,
	}
}

func (q tCreateTable) Squirrel(conf dbconf.Config) (squirrel.Sqlizer, error) {
	return internal.SQL2Squirrel(conf, q)
}

func (q tCreateTable) SQL(conf dbconf.Config) (string, error) {
	colNames := make([]string, 0, len(q.cols))
	for _, col := range q.cols {
		csql, err := col.SQL(conf)
		if err != nil {
			return "", fmt.Errorf("generate SQL for ColumnDef: %v", err)
		}
		colNames = append(colNames, csql)
	}
	cdefs := strings.Join(colNames, ", ")
	sql := fmt.Sprintf("CREATE TABLE %s (%s)", q.table, cdefs)
	return sql, nil
}
