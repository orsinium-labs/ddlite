package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementTruncateTable struct {
	table Safe
}

var _ Statement = StatementTruncateTable{}

// TruncateTable builds TRUNCATE TABLE query that removes all data from the table.
func TruncateTable(table Safe) StatementTruncateTable {
	return StatementTruncateTable{
		table: table,
	}
}

func (q StatementTruncateTable) tokens() tokens.Tokens {
	return tokens.New(
		tokens.Keyword("DELETE FROM"),
		tokens.ColumnName(q.table),
	)
}
