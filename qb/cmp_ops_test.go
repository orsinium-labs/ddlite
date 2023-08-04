package qb_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/qb"
)

func TestCmpOps(t *testing.T) {
	type User struct {
		name string
		age  int
	}
	u := User{}
	testCases := []struct {
		given qb.Expr[bool]
		sql   string
		args  []any
	}{
		{
			given: qb.Eq(qb.C(&u.name), qb.V("Aragorn")),
			sql:   "name = ?",
			args:  []any{"Aragorn"},
		},
		{
			given: qb.Eq(qb.C(&u.age), qb.V(18)),
			sql:   "age = ?",
			args:  []any{18},
		},
		{
			given: qb.E(&u.age, 18),
			sql:   "age = ?",
			args:  []any{18},
		},
		{
			given: qb.Neq(qb.C(&u.age), qb.V(18)),
			sql:   "age <> ?",
			args:  []any{18},
		},
		{
			given: qb.Gt(qb.C(&u.age), qb.V(18)),
			sql:   "age > ?",
			args:  []any{18},
		},
		{
			given: qb.Gte(qb.C(&u.age), qb.V(18)),
			sql:   "age >= ?",
			args:  []any{18},
		},
		{
			given: qb.Lt(qb.C(&u.age), qb.V(18)),
			sql:   "age < ?",
			args:  []any{18},
		},
		{
			given: qb.Lte(qb.C(&u.age), qb.V(18)),
			sql:   "age <= ?",
			args:  []any{18},
		},
		{
			given: qb.And(qb.E(&u.age, 18), qb.E(&u.age, 19)),
			sql:   "age = ? AND age = ?",
			args:  []any{18, 19},
		},
		{
			given: qb.Or(qb.E(&u.age, 18), qb.E(&u.age, 19)),
			sql:   "age = ? OR age = ?",
			args:  []any{18, 19},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.sql, func(t *testing.T) {
			is := is.New(t)
			sqlizer := tc.given.Squirrel(&u)
			sql, args, err := sqlizer.ToSql()
			is.NoErr(err)
			is.Equal(sql, tc.sql)
			is.Equal(args, tc.args)
		})
	}

}
