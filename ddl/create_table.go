package ddl

import (
	"errors"

	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementCreateTable struct {
	table       Safe
	columns     []ClauseColumn
	constraints []ClauseTableConstraint
	ifNotExists bool
}

var _ Statement = StatementCreateTable{}

// CreateTable adds a new table into the database.
//
// SQL: CREATE TABLE
//
// https://www.sqlite.org/lang_createtable.html
func CreateTable(table Safe, columns ...ClauseColumn) StatementCreateTable {
	return StatementCreateTable{
		table:   table,
		columns: columns,
	}
}

// Constraints adds to the table constraints acting on multiple fields.
//
// https://www.sqlite.org/lang_createtable.html#constraint_enforcement
func (q StatementCreateTable) Constraints(cs ...ClauseTableConstraint) StatementCreateTable {
	q.constraints = append(q.constraints, cs...)
	return q
}

// IfNotExists makes the statement to not fail if the table already exists.
//
// SQL: IF NOT EXISTS
func (q StatementCreateTable) IfNotExists(cs ...ClauseTableConstraint) StatementCreateTable {
	q.ifNotExists = true
	return q
}

func (q StatementCreateTable) tokens() tokens.Tokens {
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
		ts.Extend(col.tokens())
	}
	for _, con := range q.constraints {
		ts.Add(tokens.Comma())
		ts.Extend(con.tokens())
	}
	ts.Add(tokens.RParen())
	return ts
}
