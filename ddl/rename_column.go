package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementRenameColumn struct {
	table Safe
	old   Safe
	new   Safe
}

var _ Statement = StatementRenameColumn{}

// RenameColumn changes the column name.
//
// SQL: ALTER TABLE RENAME COLUMN
//
// https://www.sqlite.org/lang_altertable.html#alter_table_rename_column
func RenameColumn(table, old, new Safe) StatementRenameColumn {
	return StatementRenameColumn{table: table, old: old, new: new}
}

func (q StatementRenameColumn) tokens() tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.table),
		tokens.Keyword("RENAME COLUMN"),
		tokens.ColumnName(q.old),
		tokens.Keyword("TO"),
		tokens.ColumnName(q.new),
	)
	return ts
}
