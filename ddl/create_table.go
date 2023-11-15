package ddl

import (
	"errors"

	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementCreateTable struct {
	table       Safe
	columns     []ClauseColumn
	constraints []tableConstraint
}

var _ Statement = StatementCreateTable{}

func CreateTable(table Safe, columns ...ClauseColumn) StatementCreateTable {
	return StatementCreateTable{
		table:   table,
		columns: columns,
	}
}

func (q StatementCreateTable) Constraint(
	name string,
	constraint ClauseConstraint,
	column Safe,
	columns ...Safe,
) StatementCreateTable {
	tc := tableConstraint{
		name:       name,
		constraint: constraint,
		columns:    append([]Safe{column}, columns...),
	}
	q.constraints = append(q.constraints, tc)
	return q
}

type tableConstraint struct {
	name       string
	constraint ClauseConstraint
	columns    []Safe
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
		if con.name != "" {
			ts.Add(tokens.Keyword("CHECK"))
			ts.Add(tokens.Raw(con.name))
		}
		ts.Extend(con.constraint.tableTokens(dialect, con.columns))
	}
	ts.Add(tokens.RParen())
	return ts
}
