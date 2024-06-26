package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementDropTable struct {
	table    Safe
	ifExists bool
}

var _ Statement = StatementDropTable{}

// DropTable remove a table from the database.
//
// SQL: DROP TABLE
//
// https://www.sqlite.org/lang_droptable.html
func DropTable(table Safe) StatementDropTable {
	return StatementDropTable{
		table:    table,
		ifExists: false,
	}
}

// IfExists makes the statment to not fail if the table does not exist.
//
// SQL: IF EXISTS
func (q StatementDropTable) IfExists() StatementDropTable {
	q.ifExists = true
	return q
}

func (q StatementDropTable) tokens() tokens.Tokens {
	ts := tokens.New(tokens.Keyword("DROP TABLE"))
	if q.ifExists {
		ts.Add(tokens.Keyword("IF EXISTS"))
	}
	ts.Add(tokens.TableName(q.table))
	return ts
}
