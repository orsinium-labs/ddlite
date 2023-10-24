package dml_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dml"
)

func TestCmpOps(t *testing.T) {
	type User struct {
		name string
		age  int
	}
	u := User{}
	testCases := []struct {
		given dml.Expr[bool]
		sql   string
		args  []any
	}{
		{
			given: dml.Eq(dml.C(&u.name), dml.V("Aragorn")),
			sql:   "name = ?",
			args:  []any{"Aragorn"},
		},
		{
			given: dml.Eq(dml.C(&u.age), dml.V(18)),
			sql:   "age = ?",
			args:  []any{18},
		},
		{
			given: dml.E(&u.age, 18),
			sql:   "age = ?",
			args:  []any{18},
		},
		{
			given: dml.Neq(dml.C(&u.age), dml.V(18)),
			sql:   "age <> ?",
			args:  []any{18},
		},
		{
			given: dml.Gt(dml.C(&u.age), dml.V(18)),
			sql:   "age > ?",
			args:  []any{18},
		},
		{
			given: dml.Gte(dml.C(&u.age), dml.V(18)),
			sql:   "age >= ?",
			args:  []any{18},
		},
		{
			given: dml.Lt(dml.C(&u.age), dml.V(18)),
			sql:   "age < ?",
			args:  []any{18},
		},
		{
			given: dml.Lte(dml.C(&u.age), dml.V(18)),
			sql:   "age <= ?",
			args:  []any{18},
		},
		{
			given: dml.And(dml.E(&u.age, 18), dml.E(&u.age, 19)),
			sql:   "age = ? AND age = ?",
			args:  []any{18, 19},
		},
		{
			given: dml.Or(dml.E(&u.age, 18), dml.E(&u.age, 19)),
			sql:   "age = ? OR age = ?",
			args:  []any{18, 19},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.sql, func(t *testing.T) {
			is := is.New(t)
			conf := dbconf.New("postgres").WithModel(&u)
			sql, args, err := tc.given.Tokens(conf).SQL(conf)
			is.NoErr(err)
			is.Equal(sql, tc.sql)
			is.Equal(args, tc.args)
		})
	}

}
