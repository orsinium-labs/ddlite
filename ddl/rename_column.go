package ddl

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tRenameColumn struct {
	table Safe
	old   Safe
	new   Safe
}

func RenameColumn(table, old, new Safe) tRenameColumn {
	return tRenameColumn{table: table, old: old, new: new}
}

func (q tRenameColumn) Tokens(conf dbconf.Config) tokens.Tokens {
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
