package ddl

import (
	"errors"

	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementCreateTable struct {
	table       Safe
	columns     []ClauseColumn
	constraints []ClauseTableConstraint
}

var _ Statement = StatementCreateTable{}

func CreateTable(table Safe, columns ...ClauseColumn) StatementCreateTable {
	return StatementCreateTable{
		table:   table,
		columns: columns,
	}
}

func (q StatementCreateTable) Constraints(cs ...ClauseTableConstraint) StatementCreateTable {
	q.constraints = append(q.constraints, cs...)
	return q
}

func (q StatementCreateTable) statement() dialects.Feature {
	return "CREATE TABLE"
}

func (q StatementCreateTable) tokens(dialect dialects.Dialect) tokens.Tokens {
	if len(q.columns) == 0 {
		err := errors.New("new table must have columns defined")
		return tokens.New(tokens.Err(err))
	}
	ts := tokens.New(
		tokens.Keyword("CREATE TABLE"),
		tokens.TableName(q.table),
		tokens.LParen(),
	)
	for i, col := range q.columns {
		if i > 0 {
			ts.Add(tokens.Comma())
		}
		ts.Extend(col.tokens(dialect))
	}
	for _, con := range q.constraints {
		ts.Add(tokens.Comma())
		ts.Extend(con.tokens(dialect))
	}
	ts.Add(tokens.RParen())
	return ts
}
