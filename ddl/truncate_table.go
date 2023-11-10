package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type tTruncateTable struct {
	table Safe
}

// TruncateTable builds TRUNCATE TABLE query that removes all data from the table.
func TruncateTable(table Safe) Statement {
	return tTruncateTable{
		table: table,
	}
}

func (q tTruncateTable) tokens(dialect dialects.Dialect) tokens.Tokens {
	ts := tokens.New()
	if dialect.Features().TruncateTable {
		ts.Add(tokens.Keyword("TRUNCATE TABLE"))
	} else {
		ts.Add(tokens.Keyword("DELETE FROM"))
	}
	ts.Add(tokens.ColumnName(q.table))
	return ts
}
