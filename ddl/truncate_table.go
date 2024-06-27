package ddl

import (
	"github.com/orsinium-labs/ddlite/internal/tokens"
)

type StatementTruncateTable struct {
	table Safe
}

var _ Statement = StatementTruncateTable{}

// TruncateTable deletes all records from the given table.
//
// SQL: DELETE FROM
//
// https://www.sqlite.org/lang_delete.html
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
