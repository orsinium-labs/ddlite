package dml_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dialects"
	"github.com/orsinium-labs/sequel/dml"
)

func TestExpr_SQL(t *testing.T) {
	type User struct {
		name string
		age  int
	}
	u := User{}
	age18 := dml.E(&u.age, 18)
	age19 := dml.E(&u.age, 19)
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
			given: age18,
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
			given: dml.And(age18, age19),
			sql:   "age = ? AND age = ?",
			args:  []any{18, 19},
		},
		{
			given: dml.Or(age18, age19),
			sql:   "age = ? OR age = ?",
			args:  []any{18, 19},
		},
		{
			given: dml.And(dml.And(age18, age19), dml.E(&u.name, "A")),
			sql:   "age = ? AND age = ? AND name = ?",
			args:  []any{18, 19, "A"},
		},
		{
			given: dml.And(age18, dml.And(age19, dml.E(&u.name, "A"))),
			sql:   "age = ? AND (age = ? AND name = ?)",
			args:  []any{18, 19, "A"},
		},
		{
			given: dml.Or(dml.And(age18, age18), dml.And(age19, age19)),
			sql:   "age = ? AND age = ? OR age = ? AND age = ?",
			args:  []any{18, 18, 19, 19},
		},
		{
			given: dml.And(dml.Or(age18, age18), dml.Or(age19, age19)),
			sql:   "(age = ? OR age = ?) AND (age = ? OR age = ?)",
			args:  []any{18, 18, 19, 19},
		},
		{
			given: dml.Not(age18),
			sql:   "NOT age = ?",
			args:  []any{18},
		},
		{
			given: dml.And(dml.Not(age18), age19),
			sql:   "NOT age = ? AND age = ?",
			args:  []any{18, 19},
		},
		{
			given: dml.Not(dml.And(age18, age19)),
			sql:   "NOT (age = ? AND age = ?)",
			args:  []any{18, 19},
		},
		{
			given: dml.IsNull(dml.C(&u.age)),
			sql:   "age IS NULL",
			args:  []any{},
		},
		{
			given: dml.Not(dml.IsNull(dml.C(&u.age))),
			sql:   "NOT age IS NULL",
			args:  []any{},
		},
		{
			given: dml.IsNull(age18),
			sql:   "age = ? IS NULL",
			args:  []any{18},
		},
		{
			given: dml.IsNull(dml.Not(age18)),
			sql:   "(NOT age = ?) IS NULL",
			args:  []any{18},
		},
		{
			given: dml.Eq(dml.C(&u.age), dml.L(13)),
			sql:   "age = 13",
			args:  []any{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.sql, func(t *testing.T) {
			is := is.New(t)
			conf := dbconf.New(dialects.SQLite).WithModel(&u)
			sql, args, err := tc.given.Tokens(conf).SQL(conf)
			is.NoErr(err)
			is.Equal(sql, tc.sql)
			is.Equal(args, tc.args)
		})
	}

}
