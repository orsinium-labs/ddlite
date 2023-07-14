package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/qb"
)

type sqlized interface {
	SQL(m qb.Model) (string, error)
}

func TestColumnDef(t *testing.T) {
	type User struct {
		name string
		age  int
	}
	u := User{}
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
			sql, err := testCase.def.SQL(&u)
			is.NoErr(err)
			is.Equal(sql, testCase.sql)
		})
	}
}
