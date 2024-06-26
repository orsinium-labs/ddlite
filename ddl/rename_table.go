package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementRenameTable struct {
	old Safe
	new Safe
}

var _ Statement = StatementRenameColumn{}

func RenameTable(old, new Safe) Statement {
	return StatementRenameTable{old: old, new: new}
}

func (q StatementRenameTable) tokens() tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.old),
		tokens.Keyword("RENAME TO"),
		tokens.ColumnName(q.new),
	)
	return ts
}
