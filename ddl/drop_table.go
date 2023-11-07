package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type tDropTable struct {
	table    Safe
	ifExists bool
}

func DropTable(table Safe) Statement {
	return tDropTable{
		table:    table,
		ifExists: false,
	}
}

func DropTableIfExists(table Safe) Statement {
	return tDropTable{
		table:    table,
		ifExists: true,
	}
}

func (q tDropTable) tokens(dialects.Dialect) tokens.Tokens {
	ts := tokens.New(tokens.Keyword("DROP TABLE"))
	if q.ifExists {
		ts.Add(tokens.Keyword("IF EXISTS"))
	}
	ts.Add(tokens.TableName(q.table))
	return ts
}
