package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type tAddColumn struct {
	table Safe
	col   ColumnBuilder
}

// AddColumn builds query that adds a new column to the table.
//
// SQL: ALTER TABLE ADD COLUMN
func AddColumn(table Safe, col ColumnBuilder) Statement {
	return tAddColumn{table: table, col: col}
}

func (q tAddColumn) tokens(dialect dialects.Dialect) tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.table),
		tokens.Keyword("ADD COLUMN"),
	)
	ts.Extend(q.col.tokens(dialect))
	return ts
}
