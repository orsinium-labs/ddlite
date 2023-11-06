package ddl

import (
	"github.com/orsinium-labs/sequel/dialects"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tDropTable struct {
	table    Safe
	ifExists bool
}

func DropTable(table Safe) tDropTable {
	return tDropTable{
		table: table,
	}
}

func (q tDropTable) IfExists() tDropTable {
	q.ifExists = true
	return q
}

func (q tDropTable) Tokens(dialects.Dialect) tokens.Tokens {
	ts := tokens.New(tokens.Keyword("DROP TABLE"))
	if q.ifExists {
		ts.Add(tokens.Keyword("IF EXISTS"))
	}
	ts.Add(tokens.TableName(q.table))
	return ts
}
