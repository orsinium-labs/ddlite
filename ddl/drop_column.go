package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementDropColumn struct {
	table Safe
	col   Safe
}

var _ Statement = StatementDropColumn{}

func DropColumn(table Safe, col Safe) StatementDropColumn {
	return StatementDropColumn{table: table, col: col}
}

func (q StatementDropColumn) tokens(dialects.Dialect) tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(q.table),
		tokens.Keyword("DROP COLUMN"),
		tokens.ColumnName(q.col),
	)
	return ts
}
