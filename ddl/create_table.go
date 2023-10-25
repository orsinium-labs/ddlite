package ddl

import (
	"errors"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tCreateTable struct {
	table Safe
	cols  []iColumn
}

func CreateTable(table Safe, cols ...iColumn) tCreateTable {
	return tCreateTable{
		table: table,
		cols:  cols,
	}
}

func (q tCreateTable) Tokens(conf dbconf.Config) tokens.Tokens {
	if len(q.cols) == 0 {
		err := errors.New("new table must have columns defined")
		return tokens.New(tokens.Err(err))
	}
	ts := tokens.New(
		tokens.Keyword("CREATE TABLE"),
		tokens.TableName(q.table),
		tokens.LParen(),
	)
	for i, col := range q.cols {
		if i > 0 {
			ts.Add(tokens.Comma())
		}
		ts.Extend(col.Tokens(conf))
	}
	ts.Add(tokens.RParen())
	return ts
}
