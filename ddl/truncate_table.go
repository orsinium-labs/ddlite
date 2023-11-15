package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementTruncateTable struct {
	table Safe
}

var _ Statement = StatementTruncateTable{}

// TruncateTable builds TRUNCATE TABLE query that removes all data from the table.
func TruncateTable(table Safe) StatementTruncateTable {
	return StatementTruncateTable{
		table: table,
	}
}

func (q StatementTruncateTable) statement() dialects.Feature {
	return "TRUNCATE TABLE"
}

func (q StatementTruncateTable) tokens(dialect dialects.Dialect) tokens.Tokens {
	ts := tokens.New()
	if dialect.Features()["TRUNCATE TABLE"] {
		ts.Add(tokens.Keyword("TRUNCATE TABLE"))
	} else {
		ts.Add(tokens.Keyword("DELETE FROM"))
	}
	ts.Add(tokens.ColumnName(q.table))
	return ts
}
