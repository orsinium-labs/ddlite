package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
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

func (q tTruncateTable) SQL(conf dbconf.Config) (string, error) {
	// https://en.wikipedia.org/wiki/Data_definition_language#TRUNCATE_statement
	// https://www.sqlite.org/lang_delete.html
	prefix := "TRUNCATE TABLE"
	if conf.Dialect == dbconf.SQLite {
		prefix = "DELETE FROM"
	}
	sql := fmt.Sprintf("%s %s", prefix, q.table)
	return sql, nil
}
