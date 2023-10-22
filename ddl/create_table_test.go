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
	type User struct {
		name string
		age  int8
	}
	u := User{}
	conf := dbconf.New("postgres")
	q := ddl.CreateTable(&u,
		ddl.ColumnDef(&u.name, dbtypes.Text[string]()),
		ddl.ColumnDef(&u.age, dbtypes.Int8[int8]()),
	)
	sql, err := q.SQL(conf)
	is.NoErr(err)
	is.Equal(sql, "CREATE TABLE user (name TEXT, age SMALLINT)")
}

func TestColumnDef(t *testing.T) {
	type User struct {
		name string
		age  int
	}
	u := User{}
	conf := dbconf.New("postgres").WithModel(&u)
	testCases := []struct {
		def sqlized
		sql string
	}{
		{
			def: ddl.ColumnDef(&u.name, dbtypes.Text[string]()),
			sql: "name TEXT",
		},
		{
			def: ddl.ColumnDef(&u.age, dbtypes.Int32[int]()),
			sql: "age INTEGER",
		},
		{
			def: ddl.ColumnDef(&u.age, dbtypes.Int32[int]()).Unique(),
			sql: "age INTEGER UNIQUE",
		},
		{
			def: ddl.ColumnDef(&u.age, dbtypes.Int32[int]()).Null(),
			sql: "age INTEGER NULL",
		},
		{
			def: ddl.ColumnDef(&u.age, dbtypes.Int32[int]()).NotNull(),
			sql: "age INTEGER NOT NULL",
		},
		{
			def: ddl.ColumnDef(&u.age, dbtypes.Int32[int]()).PrimaryKey(),
			sql: "age INTEGER PRIMARY KEY",
		},
		{
			def: ddl.ColumnDef(&u.name, dbtypes.VarChar[string](20)).Collate("NOCASE"),
			sql: "name VARCHAR(20) COLLATE NOCASE",
		},
		{
			def: ddl.Unique(&u.age),
			sql: "UNIQUE (age)",
		},
		{
			def: ddl.Unique(&u.age, &u.name),
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
