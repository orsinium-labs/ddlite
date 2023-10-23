package ddl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dbtypes"
)

// Safe is a string that is used in SQL queries as-is, without escaping.
//
// String literals and constants are automatically considered safe.
// Variables need to be explicitly converted to Safe.
//
// Never convert to Safe untrusted input, it allows evil people to do SQL injections.
type Safe string

// A private type to represent column definitions and table constraints.
//
// Can be constructed with [Column] and [Unique].
type iColumn interface {
	SQL(dbconf.Config) (string, error)
}

type tColumn struct {
	name        Safe
	colType     dbtypes.ColumnType
	constraints []string
}

func Column(name Safe, ctype dbtypes.ColumnType) tColumn {
	return tColumn{
		name:        name,
		colType:     ctype,
		constraints: make([]string, 0),
	}
}

func (def tColumn) Null() tColumn {
	def.constraints = append(def.constraints, "NULL")
	return def
}

func (def tColumn) NotNull() tColumn {
	def.constraints = append(def.constraints, "NOT NULL")
	return def
}

func (def tColumn) Unique() tColumn {
	def.constraints = append(def.constraints, "UNIQUE")
	return def
}

func (def tColumn) PrimaryKey() tColumn {
	def.constraints = append(def.constraints, "PRIMARY KEY")
	return def
}

func (def tColumn) Collate(collationName string) tColumn {
	def.constraints = append(def.constraints, "COLLATE", collationName)
	return def
}

func (def tColumn) SQL(conf dbconf.Config) (string, error) {
	if def.name == "" {
		return "", errors.New("column name must not be empty")
	}
	constraints := strings.Join(def.constraints, " ")
	colSQL := def.colType.SQL(conf)
	sql := fmt.Sprintf("%s %s %s", def.name, colSQL, constraints)
	sql = strings.TrimRight(sql, " ")
	return sql, nil
}

type tUnique struct {
	names []Safe
}

func Unique(names ...Safe) iColumn {
	return tUnique{names: names}
}

func (def tUnique) SQL(conf dbconf.Config) (string, error) {
	if len(def.names) == 0 {
		return "", errors.New("unique index must have at least one column specified")
	}
	names := make([]string, 0, len(def.names))
	for _, name := range def.names {
		names = append(names, string(name))
	}
	joined := strings.Join(names, ", ")
	sql := fmt.Sprintf("UNIQUE (%s)", joined)
	return sql, nil
}
