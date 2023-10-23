package ddl

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
)

type tCreateTable struct {
	table Safe
	cols  []iColumn
}

func CreateTable(table Safe, cols ...iColumn) tCreateTable {
	return tCreateTable{
		table: table,
		cols:  cols,
	}
}

func (q tCreateTable) Squirrel(conf dbconf.Config) (squirrel.Sqlizer, error) {
	return internal.SQL2Squirrel(conf, q)
}

func (q tCreateTable) SQL(conf dbconf.Config) (string, error) {
	colDefs := make([]string, 0, len(q.cols))
	for _, col := range q.cols {
		colSQL, err := col.SQL(conf)
		if err != nil {
			return "", fmt.Errorf("generate SQL for ColumnDef: %v", err)
		}
		colDefs = append(colDefs, colSQL)
	}
	colDefsJoined := strings.Join(colDefs, ", ")
	sql := fmt.Sprintf("CREATE TABLE %s (%s)", q.table, colDefsJoined)
	return sql, nil
}
