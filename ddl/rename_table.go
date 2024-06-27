package ddl

import (
	"github.com/orsinium-labs/ddlite/internal/tokens"
)

type StatementRenameTable struct {
	old Safe
	new Safe
}

var _ Statement = StatementRenameColumn{}

// RenameTable changes the name of the table.
//
// SQL: ALTER TABLE RENAME TO
//
// https://www.sqlite.org/lang_altertable.html#alter_table_rename
func RenameTable(old, new Safe) Statement {
	return StatementRenameTable{old: old, new: new}
}

func (q StatementRenameTable) tokens() tokens.Tokens {
	return tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.old),
		tokens.Keyword("RENAME TO"),
		tokens.ColumnName(q.new),
	)
}
