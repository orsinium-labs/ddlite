package qb_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconfig"
	"github.com/orsinium-labs/sequel/qb"
)

type sqlized interface {
	SQL(dbconfig.Config) (string, error)
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
			def: qb.ColumnDef(&u.name, qb.Text()),
			sql: "name TEXT",
		},
		{
			def: qb.ColumnDef(&u.age, qb.Integer()),
			sql: "age INTEGER",
		},
		{
			def: qb.ColumnDef(&u.age, qb.Integer()).Unique(),
			sql: "age INTEGER UNIQUE",
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
