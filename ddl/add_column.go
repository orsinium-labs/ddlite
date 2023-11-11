package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementAddColumn struct {
	table Safe
	col   ClauseColumn
}

var _ Statement = StatementAddColumn{}

// AddColumn builds query that adds a new column to the table.
//
// SQL: ALTER TABLE ADD COLUMN
func AddColumn(table Safe, col ClauseColumn) StatementAddColumn {
	return StatementAddColumn{table: table, col: col}
}

func (q StatementAddColumn) tokens(dialect dialects.Dialect) tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.table),
		tokens.Keyword("ADD COLUMN"),
	)
	ts.Extend(q.col.tokens(dialect))
	return ts
}
