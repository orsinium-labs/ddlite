package ddl

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tDropColumn struct {
	table Safe
	col   Safe
}

func DropColumn(table Safe, col Safe) tDropColumn {
	return tDropColumn{table: table, col: col}
}

func (q tDropColumn) Tokens(conf dbconf.Config) tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.table),
		tokens.Keyword("DROP COLUMN"),
		tokens.ColumnName(q.col),
	)
	return ts
}
