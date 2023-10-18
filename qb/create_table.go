package qb

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconfig"
)

type tCreateTable struct {
	model Model
	cols  []iColumnDef
}

// A private type to represent column definitions and table constraints.
//
// Can be constructed with `qb.ColumnDef` and `qb.Unique` functions.
type iColumnDef interface {
	SQL(m Model) (string, error)
}

func CreateTable[T Model](model *T, cols ...iColumnDef) tCreateTable {
	return tCreateTable{
		model: model,
		cols:  cols,
	}
}

func (q tCreateTable) Squirrel(dbconfig.Config) (squirrel.Sqlizer, error) {
	sql, err := q.SQL()
	if err != nil {
		return nil, err
	}
	return squirrel.Expr(sql), nil
}

func (q tCreateTable) SQL() (string, error) {
	colNames := make([]string, 0, len(q.cols))
	for _, col := range q.cols {
		csql, err := col.SQL(q.model)
		if err != nil {
			return "", fmt.Errorf("generate SQL for ColumnDef: %v", csql)
		}
		colNames = append(colNames, csql)
	}
	tableName := getModelName(q.model)
	cdefs := strings.Join(colNames, ", ")
	sql := fmt.Sprintf("CREATE TABLE %s (%s)", tableName, cdefs)
	return sql, nil
}

type tColumnDef[T any] struct {
	field       *T
	colType     string
	constraints []string
}

func ColumnDef[T any](field *T, ctype ColumnType[T]) tColumnDef[T] {
	return tColumnDef[T]{
		field:       field,
		colType:     ctype.SQL(),
		constraints: make([]string, 0),
	}
}

func (def tColumnDef[T]) Null() tColumnDef[T] {
	def.constraints = append(def.constraints, "NULL")
	return def
}

func (def tColumnDef[T]) NotNull() tColumnDef[T] {
	def.constraints = append(def.constraints, "NOT NULL")
	return def
}

func (def tColumnDef[T]) Unique() tColumnDef[T] {
	def.constraints = append(def.constraints, "UNIQUE")
	return def
}

func (def tColumnDef[T]) PrimaryKey() tColumnDef[T] {
	def.constraints = append(def.constraints, "PRIMARY KEY")
	return def
}

func (def tColumnDef[T]) Collate(collationName string) tColumnDef[T] {
	def.constraints = append(def.constraints, "COLLATE", collationName)
	return def
}

func (def tColumnDef[T]) SQL(model Model) (string, error) {
	fieldName, err := getFieldName(model, def.field)
	if err != nil {
		return "", fmt.Errorf("get field name: %v", err)
	}
	constraints := strings.Join(def.constraints, " ")
	sql := fmt.Sprintf("%s %s %s", fieldName, def.colType, constraints)
	sql = strings.TrimRight(sql, " ")
	return sql, nil
}

type tUniqueDef struct {
	fields []any
}

func Unique(fields ...any) iColumnDef {
	return tUniqueDef{fields: fields}
}

func (def tUniqueDef) SQL(model Model) (string, error) {
	columnNames := make([]string, 0, len(def.fields))
	for _, field := range def.fields {
		fieldName, err := getFieldName(model, field)
		if err != nil {
			return "", fmt.Errorf("get field name: %v", err)
		}
		columnNames = append(columnNames, fieldName)
	}
	joined := strings.Join(columnNames, ", ")
	sql := fmt.Sprintf("UNIQUE (%s)", joined)
	return sql, nil
}
