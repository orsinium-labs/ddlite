package ddl

import (
	"errors"
	"fmt"

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

func (q tCreateTable) Tokens(conf dbconf.Config) (tokens.Tokens, error) {
	if len(q.cols) == 0 {
		return tokens.New(), errors.New("new table must have columns defined")
	}
	ts := tokens.New(
		tokens.Keyword("CREATE TABLE"),
		tokens.TableName(q.table),
		tokens.LParen(),
	)
	first := true
	for _, col := range q.cols {
		colTokens, err := col.Tokens(conf)
		if err != nil {
			return tokens.New(), fmt.Errorf("generate SQL for ColumnDef: %v", err)
		}
		if !first {
			ts.Add(tokens.Comma())
		}
		first = false
		ts.Add(colTokens)
	}
	ts.Add(tokens.RParen())
	return ts, nil
}
