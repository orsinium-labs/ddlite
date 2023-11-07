package ddl_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel-ddl/ddl"
	"github.com/orsinium-labs/sequel-ddl/dialects"
)

func TestCreateTable(t *testing.T) {
	is := is.New(t)
	q := ddl.CreateTable("user",
		ddl.Column("name", ddl.Text()),
		ddl.Column("age", ddl.Int(8)),
	)
	sql, err := ddl.SQL(dialects.PostgreSQL, q)
	is.NoErr(err)
	is.Equal(sql, "CREATE TABLE user (name TEXT, age SMALLINT)")
}

func TestColumnDef(t *testing.T) {
	testCases := []struct {
		def ddl.Statement
		sql string
	}{
		{
			def: ddl.Column("name", ddl.Text()),
			sql: "name TEXT",
		},
		{
			def: ddl.Column("age", ddl.Int(32)),
			sql: "age INTEGER",
		},
		{
			def: ddl.Column("age", ddl.Int(32)).Unique(),
			sql: "age INTEGER UNIQUE",
		},
		{
			def: ddl.Column("age", ddl.Int(32)).Null(),
			sql: "age INTEGER NULL",
		},
		{
			def: ddl.Column("age", ddl.Int(32)).NotNull(),
			sql: "age INTEGER NOT NULL",
		},
		{
			def: ddl.Column("age", ddl.Int(32)).PrimaryKey(),
			sql: "age INTEGER PRIMARY KEY",
		},
		{
			def: ddl.Column("name", ddl.VarChar(20)).Collate("NOCASE"),
			sql: "name VARCHAR(20) COLLATE NOCASE",
		},
		{
			def: ddl.Unique("age"),
			sql: "UNIQUE (age)",
		},
		{
			def: ddl.Unique("age", "name"),
			sql: "UNIQUE (age, name)",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.sql, func(t *testing.T) {
			is := is.New(t)
			sql, err := ddl.SQL(dialects.PostgreSQL, testCase.def)
			is.NoErr(err)
			is.Equal(sql, testCase.sql)
		})
	}
}
