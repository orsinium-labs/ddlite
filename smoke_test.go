package sequel_test

import (
	"testing"

	"github.com/matryer/is"
	sq "github.com/orsinium-labs/sequel"
	"github.com/orsinium-labs/sequel/pgext"
)

func Test(t *testing.T) {
	is := is.New(t)

	type User struct {
		name string
		age  int
	}
	u := User{}
	q := sq.Select(&u, &u.name, &u.age)
	q = q.Where(sq.Gt(sq.C(&u.age), sq.V(18)))
	q = q.And(sq.Gt(sq.C(&u.age), pgext.Abs(sq.V(-18))))
	is.Equal(q.String(), "SELECT name, age FROM user WHERE (age > $1 AND age > abs($2))")
}
