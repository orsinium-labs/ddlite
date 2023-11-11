package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementRenameColumn struct {
	table Safe
	old   Safe
	new   Safe
}

var _ Statement = StatementRenameColumn{}

func RenameColumn(table, old, new Safe) StatementRenameColumn {
	return StatementRenameColumn{table: table, old: old, new: new}
}

func (q StatementRenameColumn) tokens(dialects.Dialect) tokens.Tokens {
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
