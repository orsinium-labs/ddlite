package ddl

import (
	"fmt"
	"strings"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dbtypes"
)

// A private type to represent column definitions and table constraints.
//
// Can be constructed with `dml.ColumnDef` and `dml.Unique` functions.
type iColumnDef interface {
	SQL(dbconf.Config) (string, error)
}

type tColumnDef struct {
	name        string
	colType     dbtypes.ColumnType
	constraints []string
}

func ColumnDef(name string, ctype dbtypes.ColumnType) tColumnDef {
	return tColumnDef{
		name:        name,
		colType:     ctype,
		constraints: make([]string, 0),
	}
}

func (def tColumnDef) Null() tColumnDef {
	def.constraints = append(def.constraints, "NULL")
	return def
}

func (def tColumnDef) NotNull() tColumnDef {
	def.constraints = append(def.constraints, "NOT NULL")
	return def
}

func (def tColumnDef) Unique() tColumnDef {
	def.constraints = append(def.constraints, "UNIQUE")
	return def
}

func (def tColumnDef) PrimaryKey() tColumnDef {
	def.constraints = append(def.constraints, "PRIMARY KEY")
	return def
}

func (def tColumnDef) Collate(collationName string) tColumnDef {
	def.constraints = append(def.constraints, "COLLATE", collationName)
	return def
}

func (def tColumnDef) SQL(conf dbconf.Config) (string, error) {
	constraints := strings.Join(def.constraints, " ")
	colSQL := def.colType.SQL(conf)
	sql := fmt.Sprintf("%s %s %s", def.name, colSQL, constraints)
	sql = strings.TrimRight(sql, " ")
	return sql, nil
}

type tUniqueDef struct {
	names []string
}

func Unique(names ...string) iColumnDef {
	return tUniqueDef{names: names}
}

func (def tUniqueDef) SQL(conf dbconf.Config) (string, error) {
	joined := strings.Join(def.names, ", ")
	sql := fmt.Sprintf("UNIQUE (%s)", joined)
	return sql, nil
}
