package ddl_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/ddlite/ddl"
)

func TestCreateTable(t *testing.T) {
	is := is.New(t)
	q := ddl.CreateTable("user",
		ddl.Column("name", ddl.Text, ddl.Null),
		ddl.Column("age", ddl.Integer, ddl.Null),
	).Constraints(
		ddl.Constraint("", ddl.Unique(), "name"),
	)
	sql, err := ddl.SQL(q)
	is.NoErr(err)
	is.Equal(sql, "CREATE TABLE user (name TEXT, age INTEGER, UNIQUE (name))")
}

func TestCreateTable_IfNotExists(t *testing.T) {
	is := is.New(t)
	q := ddl.CreateTable("user",
		ddl.Column("name", ddl.Text, ddl.Null).Default(`"greg"`),
	).IfNotExists()
	sql, err := ddl.SQL(q)
	is.NoErr(err)
	is.Equal(sql, `CREATE TABLE IF NOT EXISTS user (name TEXT DEFAULT "greg")`)
}

func TestCreateTable_Temp(t *testing.T) {
	is := is.New(t)
	q := ddl.CreateTable("user",
		ddl.Column("name", ddl.Text, ddl.Null).Collate("BINARY"),
	).Temp()
	sql, err := ddl.SQL(q)
	is.NoErr(err)
	is.Equal(sql, "CREATE TEMP TABLE user (name TEXT COLLATE BINARY)")
}
