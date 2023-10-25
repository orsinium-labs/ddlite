package ddl

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tRenameTable struct {
	old Safe
	new Safe
}

func RenameTable(old, new Safe) tRenameTable {
	return tRenameTable{old: old, new: new}
}

func (q tRenameTable) Tokens(conf dbconf.Config) (tokens.Tokens, error) {
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.old),
		tokens.Keyword("RENAME TO"),
		tokens.ColumnName(q.new),
	)
	return ts, nil
}
