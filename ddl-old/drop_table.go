package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementDropTable struct {
	table    Safe
	ifExists bool
}

var _ Statement = StatementDropTable{}

func DropTable(table Safe) StatementDropTable {
	return StatementDropTable{
		table:    table,
		ifExists: false,
	}
}

func (q StatementDropTable) IfExists() StatementDropTable {
	q.ifExists = true
	return q
}

func (q StatementDropTable) statement() dialects.Feature {
	return "DROP TABLE"
}

func (q StatementDropTable) tokens(dialects.Dialect) tokens.Tokens {
	ts := tokens.New(tokens.Keyword("DROP TABLE"))
	if q.ifExists {
		ts.Add(tokens.Keyword("IF EXISTS"))
	}
	ts.Add(tokens.TableName(q.table))
	return ts
}
