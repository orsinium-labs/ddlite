package ddl

import (
	"errors"

	"github.com/orsinium-labs/sequel/dbconf"
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

func (q tTruncateTable) Tokens(conf dbconf.Config) (tokens.Tokens, error) {
	if q.table == "" {
		return tokens.New(), errors.New("table name must not be empty")
	}
	ts := tokens.New()
	if conf.Dialect == dbconf.SQLite {
		// https://www.sqlite.org/lang_delete.html
		ts.Add(tokens.Keyword("DELETE FROM"))
	} else {
		// https://en.wikipedia.org/wiki/Data_definition_language#TRUNCATE_statement
		ts.Add(tokens.Keyword("TRUNCATE TABLE"))
	}
	ts.Add(tokens.ColumnName(q.table))
	return ts, nil
}
