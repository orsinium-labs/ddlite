package ddl

import (
	"errors"

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

func (q tRenameColumn) SQL(conf dbconf.Config) (string, error) {
	if q.table == "" {
		return "", errors.New("table name must not be empty")
	}
	if q.old == "" {
		return "", errors.New("old column name must not be empty")
	}
	if q.new == "" {
		return "", errors.New("new column name must not be empty")
	}
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.table),
		tokens.Keyword("RENAME COLUMN"),
		tokens.ColumnName(q.old),
		tokens.Keyword("TO"),
		tokens.ColumnName(q.new),
	)
	return ts.SQL(conf)
}
