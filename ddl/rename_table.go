package ddl

import (
	"errors"

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

func (q tRenameTable) SQL(conf dbconf.Config) (string, error) {
	if q.old == "" {
		return "", errors.New("old table name must not be empty")
	}
	if q.new == "" {
		return "", errors.New("new table name must not be empty")
	}
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.old),
		tokens.Keyword("RENAME TO"),
		tokens.ColumnName(q.new),
	)
	return ts.SQL(conf)
}
