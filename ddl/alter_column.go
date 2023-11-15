package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type StatementAlterColumn struct {
	table  Safe
	column Safe
}

type stmtAlterColumn struct {
	StatementAlterColumn
	suffix func(dialects.Dialect) tokens.Tokens
}

var _ Statement = stmtAlterColumn{}

func AlterColumn(table, column Safe) StatementAlterColumn {
	return StatementAlterColumn{
		table:  table,
		column: column,
	}
}

func (stmt StatementAlterColumn) SetDefault(expr Safe) Statement {
	return stmtAlterColumn{
		stmt,
		func(d dialects.Dialect) tokens.Tokens {
			return tokens.New(tokens.Keyword("SET DEFAULT"), tokens.Raw(expr))
		},
	}
}

func (stmt StatementAlterColumn) DropDefault() Statement {
	return stmtAlterColumn{
		stmt,
		func(d dialects.Dialect) tokens.Tokens {
			return tokens.New(tokens.Keyword("DROP DEFAULT"))
		},
	}
}

func (stmt StatementAlterColumn) Set(null Nullable) Statement {
	return stmtAlterColumn{
		stmt,
		func(d dialects.Dialect) tokens.Tokens {
			if null == Null {
				return tokens.New(tokens.Keyword("DROP NOT NULL"))
			}
			return tokens.New(tokens.Keyword("SET NOT NULL"))
		},
	}
}

func (stmt StatementAlterColumn) SetDataType(dtype ClauseDataType) Statement {
	return stmtAlterColumn{
		stmt,
		func(d dialects.Dialect) tokens.Tokens {
			return tokens.New(tokens.Keyword("TYPE"), tokens.Raw(dtype(d)))
		},
	}
}

func (q stmtAlterColumn) statement() dialects.Feature {
	return "ALTER TABLE / ALTER COLUMN"
}

func (stmt stmtAlterColumn) tokens(d dialects.Dialect) tokens.Tokens {
	ts := tokens.New(
		tokens.Keyword("ALTER TABLE"),
		tokens.TableName(stmt.table),
		tokens.Keyword("ALTER COLUMN"),
		tokens.TableName(stmt.column),
	)
	ts.Extend(stmt.suffix(d))
	return ts
}
