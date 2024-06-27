package ddl

import (
	"github.com/orsinium-labs/ddlite/internal/tokens"
)

type StatementDropColumn struct {
	table Safe
	col   Safe
}

var _ Statement = StatementDropColumn{}

// DropColumn removes a column from a table.
//
// SQL: ALTER TABLE DROP COLUMN
//
// https://www.sqlite.org/lang_altertable.html#alter_table_drop_column
func DropColumn(table Safe, col Safe) StatementDropColumn {
	return StatementDropColumn{table: table, col: col}
}

func (q StatementDropColumn) tokens() tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.table),
		tokens.Keyword("DROP COLUMN"),
		tokens.ColumnName(q.col),
	)
	return ts
}
