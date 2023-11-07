package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type tDropColumn struct {
	table Safe
	col   Safe
}

func DropColumn(table Safe, col Safe) Statement {
	return tDropColumn{table: table, col: col}
}

func (q tDropColumn) tokens(dialects.Dialect) tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.table),
		tokens.Keyword("DROP COLUMN"),
		tokens.ColumnName(q.col),
	)
	return ts
}
