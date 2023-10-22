package ddl

import (
	"fmt"
	"strings"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dbtypes"
	"github.com/orsinium-labs/sequel/internal"
)

// A private type to represent column definitions and table constraints.
//
// Can be constructed with `dml.ColumnDef` and `dml.Unique` functions.
type iColumnDef interface {
	SQL(dbconf.Config) (string, error)
}

type tColumnDef[T any] struct {
	field       *T
	colType     dbtypes.ColumnType[T]
	constraints []string
}

func ColumnDef[T any](field *T, ctype dbtypes.ColumnType[T]) tColumnDef[T] {
	return tColumnDef[T]{
		field:       field,
		colType:     ctype,
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

func (def tColumnDef[T]) SQL(conf dbconf.Config) (string, error) {
	fieldName, err := internal.GetColumnName(conf, def.field)
	if err != nil {
		return "", fmt.Errorf("get field name: %v", err)
	}
	constraints := strings.Join(def.constraints, " ")
	colSQL := def.colType.SQL(conf)
	sql := fmt.Sprintf("%s %s %s", fieldName, colSQL, constraints)
	sql = strings.TrimRight(sql, " ")
	return sql, nil
}

type tUniqueDef struct {
	fields []any
}

func Unique(fields ...any) iColumnDef {
	return tUniqueDef{fields: fields}
}

func (def tUniqueDef) SQL(conf dbconf.Config) (string, error) {
	columnNames := make([]string, 0, len(def.fields))
	for _, field := range def.fields {
		fieldName, err := internal.GetColumnName(conf, field)
		if err != nil {
			return "", fmt.Errorf("get field name: %v", err)
		}
		columnNames = append(columnNames, fieldName)
	}
	joined := strings.Join(columnNames, ", ")
	sql := fmt.Sprintf("UNIQUE (%s)", joined)
	return sql, nil
}