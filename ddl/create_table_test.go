package ddl_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dbtypes"
	"github.com/orsinium-labs/sequel/ddl"
)

type sqlized interface {
	SQL(dbconf.Config) (string, error)
}

func TestCreateTable(t *testing.T) {
	is := is.New(t)
	conf := dbconf.New("postgres")
	q := ddl.CreateTable("user",
		ddl.ColumnDef("name", dbtypes.Text()),
		ddl.ColumnDef("age", dbtypes.Int8()),
	)
	sql, err := q.SQL(conf)
	is.NoErr(err)
	is.Equal(sql, "CREATE TABLE user (name TEXT, age SMALLINT)")
}

func TestColumnDef(t *testing.T) {
	conf := dbconf.New("postgres")
	testCases := []struct {
		def sqlized
		sql string
	}{
		{
			def: ddl.ColumnDef("name", dbtypes.Text()),
			sql: "name TEXT",
		},
		{
			def: ddl.ColumnDef("age", dbtypes.Int32()),
			sql: "age INTEGER",
		},
		{
			def: ddl.ColumnDef("age", dbtypes.Int32()).Unique(),
			sql: "age INTEGER UNIQUE",
		},
		{
			def: ddl.ColumnDef("age", dbtypes.Int32()).Null(),
			sql: "age INTEGER NULL",
		},
		{
			def: ddl.ColumnDef("age", dbtypes.Int32()).NotNull(),
			sql: "age INTEGER NOT NULL",
		},
		{
			def: ddl.ColumnDef("age", dbtypes.Int32()).PrimaryKey(),
			sql: "age INTEGER PRIMARY KEY",
		},
		{
			def: ddl.ColumnDef("name", dbtypes.VarChar(20)).Collate("NOCASE"),
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
			sql, err := testCase.def.SQL(conf)
			is.NoErr(err)
			is.Equal(sql, testCase.sql)
		})
	}
}
