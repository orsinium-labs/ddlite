package ddl

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
)

type tTruncateTable struct {
	model internal.Model
}

// TruncateTable builds TRUNCATE TABLE query that removes all data from the table.
func TruncateTable[T internal.Model](model *T) tTruncateTable {
	return tTruncateTable{
		model: model,
	}
}

func (q tTruncateTable) SQL(conf dbconf.Config) (string, error) {
	// https://en.wikipedia.org/wiki/Data_definition_language#TRUNCATE_statement
	// https://www.sqlite.org/lang_delete.html
	tableName := internal.GetTableName(conf, q.model)
	prefix := "TRUNCATE TABLE"
	if conf.Dialect == dbconf.SQLite {
		prefix = "DELETE FROM"
	}
	sql := fmt.Sprintf("%s %s", prefix, tableName)
	return sql, nil
}
