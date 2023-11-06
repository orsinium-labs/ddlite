package ddl

import (
	"github.com/orsinium-labs/sequel/dialects"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tRenameTable struct {
	old Safe
	new Safe
}

func RenameTable(old, new Safe) tRenameTable {
	return tRenameTable{old: old, new: new}
}

func (q tRenameTable) Tokens(dialects.Dialect) tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.old),
		tokens.Keyword("RENAME TO"),
		tokens.ColumnName(q.new),
	)
	return ts
}
