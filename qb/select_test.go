package qb_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbfuncs"
	"github.com/orsinium-labs/sequel/qb"
)

func TestSelectString(t *testing.T) {
	is := is.New(t)

	type User struct {
		name string
		age  int
	}
	u := User{}
	q := qb.Select(&u, &u.name, &u.age)
	q = q.Where(qb.Gt(qb.C(&u.age), qb.V(18)))
	q = q.And(qb.Gt(qb.C(&u.age), dbfuncs.Abs(qb.V(-18))))
	is.Equal(q.String(), "SELECT name, age FROM user WHERE (age > ? AND age > abs(?))")
}
