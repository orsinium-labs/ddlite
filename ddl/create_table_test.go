package ddl_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dbtypes"
	"github.com/orsinium-labs/sequel/ddl"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tokener interface {
	Tokens(dbconf.Config) (tokens.Tokens, error)
}

func TestCreateTable(t *testing.T) {
	is := is.New(t)
	conf := dbconf.New("postgres")
	q := ddl.CreateTable("user",
		ddl.Column("name", dbtypes.Text()),
		ddl.Column("age", dbtypes.Int8()),
	)
	sql, _, err := ddl.SQL(conf, q)
	is.NoErr(err)
	is.Equal(sql, "CREATE TABLE user ( name TEXT, age SMALLINT )")
}

func TestColumnDef(t *testing.T) {
	conf := dbconf.New("postgres")
	testCases := []struct {
		def tokener
		sql string
	}{
		{
			def: ddl.Column("name", dbtypes.Text()),
			sql: "name TEXT",
		},
		{
			def: ddl.Column("age", dbtypes.Int32()),
			sql: "age INTEGER",
		},
		{
			def: ddl.Column("age", dbtypes.Int32()).Unique(),
			sql: "age INTEGER UNIQUE",
		},
		{
			def: ddl.Column("age", dbtypes.Int32()).Null(),
			sql: "age INTEGER NULL",
		},
		{
			def: ddl.Column("age", dbtypes.Int32()).NotNull(),
			sql: "age INTEGER NOT NULL",
		},
		{
			def: ddl.Column("age", dbtypes.Int32()).PrimaryKey(),
			sql: "age INTEGER PRIMARY KEY",
		},
		{
			def: ddl.Column("name", dbtypes.VarChar(20)).Collate("NOCASE"),
			sql: "name VARCHAR(20) COLLATE NOCASE",
		},
		{
			def: ddl.Unique("age"),
			sql: "UNIQUE ( age )",
		},
		{
			def: ddl.Unique("age", "name"),
			sql: "UNIQUE ( age, name )",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.sql, func(t *testing.T) {
			is := is.New(t)
			sql, _, err := ddl.SQL(conf, testCase.def)
			is.NoErr(err)
			is.Equal(sql, testCase.sql)
		})
	}
}
