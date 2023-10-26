package ddl

import (
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dialects"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tTruncateTable struct {
	table Safe
}

// TruncateTable builds TRUNCATE TABLE query that removes all data from the table.
func TruncateTable(table Safe) tTruncateTable {
	return tTruncateTable{
		table: table,
	}
}

func (q tTruncateTable) Tokens(conf dbconf.Config) tokens.Tokens {
	ts := tokens.New()
	if conf.Dialect == dialects.SQLite {
		// https://www.sqlite.org/lang_delete.html
		ts.Add(tokens.Keyword("DELETE FROM"))
	} else {
		// https://en.wikipedia.org/wiki/Data_definition_language#TRUNCATE_statement
		ts.Add(tokens.Keyword("TRUNCATE TABLE"))
	}
	ts.Add(tokens.ColumnName(q.table))
	return ts
}
