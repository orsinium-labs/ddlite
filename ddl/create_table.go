package ddl

import (
	"errors"

	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type tCreateTable struct {
	table       Safe
	columns     []ColumnBuilder
	constraints []Constraint
}

func CreateTable(table Safe, columns ...ColumnBuilder) Statement {
	return tCreateTable{
		table:   table,
		columns: columns,
	}
}

func (q tCreateTable) Constraints(constraints ...Constraint) Statement {
	q.constraints = append(q.constraints, constraints...)
	return q
}

func (q tCreateTable) tokens(dialect dialects.Dialect) tokens.Tokens {
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
