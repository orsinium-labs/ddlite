package ddl

import (
	"github.com/orsinium-labs/ddl/dialects"
	"github.com/orsinium-labs/ddl/internal/tokens"
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
	if dialect == dialects.SQLite {
		// https://www.sqlite.org/lang_delete.html
		ts.Add(tokens.Keyword("DELETE FROM"))
	} else {
		// https://en.wikipedia.org/wiki/Data_definition_language#TRUNCATE_statement
		ts.Add(tokens.Keyword("TRUNCATE TABLE"))
	}
	ts.Add(tokens.ColumnName(q.table))
	return ts
}
