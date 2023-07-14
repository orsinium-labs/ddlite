package qb

import (
	"fmt"
	"strings"
)

type tNull bool

const (
	Null    tNull = true
	NotNull tNull = false
)

type tCreateTable struct {
	model Model
	cols  []tColumnDef
}

func CreateTable[T Model](model *T, cols ...tColumnDef) tCreateTable {
	return tCreateTable{
		model: model,
		cols:  cols,
	}
}

func (q tCreateTable) SQL() (string, error) {
	cols := make([]string, 0, len(q.cols))
	for _, col := range q.cols {
		csql, err := col.SQL(q.model)
		if err != nil {
			return "", fmt.Errorf("generate SQL for ColumnDef: %v", csql)
		}
		cols = append(cols, csql)
	}
	tname := getModelName(q.model)
	cdefs := strings.Join(cols, ", ")
	sql := fmt.Sprintf("CREATE TABLE %s (%s)", tname, cdefs)
	return sql, nil
}

type tColumnDef struct {
	field any
	ctype string
	null  tNull
}

func ColumnDef[T any](field *T, ctype ColumnType[T], null tNull) tColumnDef {
	return tColumnDef{
		field: field,
		ctype: ctype.SQL(),
		null:  null,
	}
}

func (def tColumnDef) SQL(m Model) (string, error) {
	fname, err := getFieldName(m, def.field)
	if err != nil {
		return "", fmt.Errorf("get field name: %v", err)
	}
	null := "NOT NULL"
	if def.null {
		null = "NULL"
	}
	sql := fmt.Sprintf("%s %s %s", fname, def.ctype, null)
	return sql, nil
}
