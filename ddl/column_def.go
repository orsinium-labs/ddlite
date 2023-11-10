package ddl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
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
	null        Nullable
}

// Nullable is used by [Column] to indicate if the column may be NULL or not.
type Nullable bool

const (
	// Null marks a column that allows NULL values.
	Null Nullable = true
	// NotNull marks a column that cannot store NULL values.
	NotNull Nullable = false
)

func Column(name Safe, ctype ColumnType, null Nullable) tColumn {
	return tColumn{
		name:        name,
		colType:     ctype,
		constraints: make([]string, 0),
	}
}

func (def tColumn) Unique() tColumn {
	def.constraints = append(def.constraints, "UNIQUE")
	return def
}

func (def tColumn) PrimaryKey() tColumn {
	def.constraints = append(def.constraints, "PRIMARY KEY")
	return def
}

func (def tColumn) Collate(collationName Safe) tColumn {
	def.constraints = append(def.constraints, "COLLATE", string(collationName))
	return def
}

func (def tColumn) tokens(dialect dialects.Dialect) tokens.Tokens {
	constraints := strings.Join(def.constraints, " ")
	colSQL := def.colType(dialect)
	if colSQL == "" {
		const msg = "the data type used for the column '%s' is not supported by the dialect"
		err := fmt.Errorf(msg, def.name)
		return tokens.New(tokens.Err(err))
	}
	ts := tokens.New(
		tokens.ColumnName(def.name),
		tokens.Raw(colSQL),
	)
	if !def.null {
		ts.Add(tokens.Keyword("NOT NULL"))
	}
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
