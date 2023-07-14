package tests

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/qb"
	"github.com/orsinium-labs/qb/pgext"
)

func Test(t *testing.T) {
	is := is.New(t)

	type User struct {
		name string
		age  int
	}
	u := User{}
	q := qb.Select(&u, &u.name, &u.age)
	q = q.Where(qb.Gt(qb.C(&u.age), qb.V(18)))
	q = q.And(qb.Gt(qb.C(&u.age), pgext.Abs(qb.V(-18))))
	is.Equal(q.String(), "SELECT name, age FROM user WHERE (age > $1 AND age > abs($2))")
}
