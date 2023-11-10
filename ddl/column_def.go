package ddl

import (
	"fmt"
	"strings"

	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type ColumnBuilder struct {
	name        Safe
	colType     ColumnType
	constraints []string
	null        Nullable
	reference   *Reference
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
func Column(name Safe, ctype ColumnType, null Nullable) ColumnBuilder {
	return ColumnBuilder{
		name:        name,
		colType:     ctype,
		constraints: make([]string, 0),
		null:        null,
	}
}

// Unique makes sure that each value in the column is unique.
//
// SQL: UNIQUE
func (def ColumnBuilder) Unique() ColumnBuilder {
	def.constraints = append(def.constraints, "UNIQUE")
	return def
}

// PrimaryKey makes the column the primary key.
//
// Only one column can be marked as primary key that way. If you want the primary key
// to consist of multiple columns, use the [PrimaryKey] constraint instead.
//
// SQL: PRIMARY KEY
func (def ColumnBuilder) PrimaryKey() ColumnBuilder {
	def.constraints = append(def.constraints, "PRIMARY KEY")
	return def
}

func (def ColumnBuilder) ForeignKey(ref Reference) ColumnBuilder {
	def.reference = &ref
	return def
}

// Collate specifies the name of a collating sequence to use as the default collation sequence for the column.
//
// SQL: COLLATE
func (def ColumnBuilder) Collate(collationName Safe) ColumnBuilder {
	def.constraints = append(def.constraints, "COLLATE", string(collationName))
	return def
}

func (def ColumnBuilder) Check(expr Safe) ColumnBuilder {
	def.constraints = append(def.constraints, "CHECK", "(", string(expr), ")")
	return def
}

func (def ColumnBuilder) Default(expr Safe) ColumnBuilder {
	def.constraints = append(def.constraints, "DEFAULT", "(", string(expr), ")")
	return def
}

func (def ColumnBuilder) tokens(dialect dialects.Dialect) tokens.Tokens {
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
	if def.reference != nil {
		ts.Extend(def.reference.tokens(dialect))
	}
	return ts
}
