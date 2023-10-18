package qb_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconfig"
	"github.com/orsinium-labs/sequel/dbtypes"
	"github.com/orsinium-labs/sequel/qb"
)

type sqlized interface {
	SQL(dbconfig.Config) (string, error)
}

func TestCreateTable(t *testing.T) {
	is := is.New(t)
	type User struct {
		name string
		age  int
	}
	u := User{}
	conf := dbconfig.New("postgres")
	q := qb.CreateTable(&u,
		qb.ColumnDef(&u.name, dbtypes.Text()),
		qb.ColumnDef(&u.age, dbtypes.SmallInt()),
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
	conf := dbconfig.New("postgres").WithModel(&u)
	testCases := []struct {
		def sqlized
		sql string
	}{
		{
			def: qb.ColumnDef(&u.name, dbtypes.Text()),
			sql: "name TEXT",
		},
		{
			def: qb.ColumnDef(&u.age, dbtypes.Integer()),
			sql: "age INTEGER",
		},
		{
			def: qb.ColumnDef(&u.age, dbtypes.Integer()).Unique(),
			sql: "age INTEGER UNIQUE",
		},
		{
			def: qb.ColumnDef(&u.age, dbtypes.Integer()).Null(),
			sql: "age INTEGER NULL",
		},
		{
			def: qb.ColumnDef(&u.age, dbtypes.Integer()).NotNull(),
			sql: "age INTEGER NOT NULL",
		},
		{
			def: qb.ColumnDef(&u.age, dbtypes.Integer()).PrimaryKey(),
			sql: "age INTEGER PRIMARY KEY",
		},
		{
			def: qb.ColumnDef(&u.name, dbtypes.VarChar(20)).Collate("NOCASE"),
			sql: "name VARCHAR(20) COLLATE NOCASE",
		},
		{
			def: qb.Unique(&u.age),
			sql: "UNIQUE (age)",
		},
		{
			def: qb.Unique(&u.age, &u.name),
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
