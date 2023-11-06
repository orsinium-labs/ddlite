package ddl

import (
	"errors"
	"strings"

	"github.com/orsinium-labs/ddl/dialects"
	"github.com/orsinium-labs/ddl/internal/tokens"
)

// A private type to represent column definitions and table constraints.
//
// Can be constructed with [Column] and [Unique].
type iColumn interface {
	tokens(dialects.Dialect) tokens.Tokens
}

type tColumn struct {
	name        Safe
	colType     ColumnType
	constraints []string
}

func Column(name Safe, ctype ColumnType) tColumn {
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

func (def tColumn) tokens(dialect dialects.Dialect) tokens.Tokens {
	constraints := strings.Join(def.constraints, " ")
	colSQL := def.colType.SQL(dialect)
	ts := tokens.New(
		tokens.ColumnName(def.name),
		tokens.Raw(colSQL),
	)
	if constraints != "" {
		ts.Add(tokens.Raw(constraints))
	}
	return ts
}

type tUnique struct {
	names []Safe
}

func Unique(names ...Safe) iColumn {
	return tUnique{names: names}
}

func (def tUnique) tokens(dialects.Dialect) tokens.Tokens {
	if len(def.names) == 0 {
		err := errors.New("unique index must have at least one column specified")
		return tokens.New(tokens.Err(err))
	}
	ts := tokens.New(
		tokens.Keyword("UNIQUE"),
		tokens.LParen(),
		tokens.Raws(def.names...),
		tokens.RParen(),
	)
	return ts
}
