package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementAddColumn struct {
	table Safe
	col   ClauseColumn
}

var _ Statement = StatementAddColumn{}

// AddColumn adds a new column at the end of an existing table.
//
// SQL: ALTER TABLE ADD COLUMN
//
// https://www.sqlite.org/lang_altertable.html#alter_table_add_column
func AddColumn(table Safe, col ClauseColumn) StatementAddColumn {
	return StatementAddColumn{table: table, col: col}
}

func (q StatementAddColumn) tokens() tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.table),
		tokens.Keyword("ADD COLUMN"),
	)
	ts.Extend(q.col.tokens())
	return ts
}
