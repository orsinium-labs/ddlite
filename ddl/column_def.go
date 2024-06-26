package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type ClauseColumn struct {
	name        Safe
	colType     DataType
	constraints []ClauseConstraint
	suffix      tokens.Tokens
	null        Nullable
}

// Nullable is used by [Column] to indicate if the column may be NULL or not.
type Nullable bool

const (
	// Null marks a column that allows NULL values.
	//
	// SQL: NULL
	Null Nullable = true

	// NotNull marks a column that cannot store NULL values.
	//
	// SQL: NOT NULL
	NotNull Nullable = false
)

// Column is a column definition.
//
// Used in [CreateTable] and [AddColumn].
func Column(
	name Safe,
	ctype DataType,
	null Nullable,
	constraints ...ClauseConstraint,
) ClauseColumn {
	return ClauseColumn{
		name:        name,
		colType:     ctype,
		constraints: constraints,
		null:        null,
		suffix:      tokens.New(),
	}
}

// Collate specifies the name of a collating sequence to use as the default collation sequence for the column.
//
// SQL: COLLATE
func (def ClauseColumn) Collate(collationName Safe) ClauseColumn {
	def.suffix.Add(tokens.Keyword("COLLATE"))
	def.suffix.Add(tokens.Raw(collationName))
	return def
}

func (def ClauseColumn) Default(expr Safe) ClauseColumn {
	def.suffix.Add(tokens.Keyword("DEFAULT"))
	def.suffix.Add(tokens.LParen())
	def.suffix.Add(tokens.Raw(expr))
	def.suffix.Add(tokens.RParen())
	return def
}

func (def ClauseColumn) tokens(dialect dialects.Dialect) tokens.Tokens {
	colSQL := def.colType
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
	ts.Extend(def.suffix)
	for _, con := range def.constraints {
		ts.Extend(con.columnTokens(dialect))
	}
	return ts
}
