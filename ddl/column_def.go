package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

// ClauseColumn is a column definition. Constructed by [Column].
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
//
// https://www.sqlite.org/datatype3.html#collation
func (def ClauseColumn) Collate(collationName Safe) ClauseColumn {
	def.suffix.Add(tokens.Keyword("COLLATE"))
	def.suffix.Add(tokens.Raw(collationName))
	return def
}

// Default specifies the default expression for the column.
//
// Note that this is the raw SQL. If you want to set a default string literal
// for a [Text] field, make sure it is properly escaped and includes quotes.
//
// SQL: DEFAULT
//
// https://www.sqlite.org/lang_createtable.html#the_default_clause
func (def ClauseColumn) Default(expr Safe) ClauseColumn {
	def.suffix.Add(tokens.Keyword("DEFAULT"))
	def.suffix.Add(tokens.Raw(expr))
	return def
}

func (def ClauseColumn) tokens() tokens.Tokens {
	colSQL := def.colType
	ts := tokens.New(
		tokens.ColumnName(def.name),
		tokens.Raw(colSQL),
	)
	if !def.null {
		ts.Add(tokens.Keyword("NOT NULL"))
	}
	ts.Extend(def.suffix)
	for _, con := range def.constraints {
		ts.Extend(con.columnTokens())
	}
	return ts
}
