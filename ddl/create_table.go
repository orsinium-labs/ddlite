package ddl

import (
	"github.com/orsinium-labs/ddlite/internal/tokens"
)

type StatementCreateTable struct {
	table       Safe
	columns     []ClauseColumn
	constraints []ClauseTableConstraint
	ifNotExists bool
	temp        bool
}

var _ Statement = StatementCreateTable{}

// CreateTable adds a new table into the database.
//
// SQL: CREATE TABLE
//
// https://www.sqlite.org/lang_createtable.html
func CreateTable(table Safe, column ClauseColumn, columns ...ClauseColumn) StatementCreateTable {
	return StatementCreateTable{
		table:   table,
		columns: append([]ClauseColumn{column}, columns...),
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
func (q StatementCreateTable) IfNotExists() StatementCreateTable {
	q.ifNotExists = true
	return q
}

// Create the table in a temporary database.
//
// SQL: TEMP
func (q StatementCreateTable) Temp() StatementCreateTable {
	q.temp = true
	return q
}

func (q StatementCreateTable) tokens() tokens.Tokens {
	ts := tokens.New(tokens.Keyword("CREATE"))
	if q.temp {
		ts.Add(tokens.Keyword("TEMP"))
	}
	ts.Add(tokens.Keyword("TABLE"))
	if q.ifNotExists {
		ts.Add(tokens.Keyword("IF NOT EXISTS"))
	}
	ts.Add(tokens.TableName(q.table))

	ts.Add(tokens.LParen())
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
