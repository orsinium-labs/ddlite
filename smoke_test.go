package sequel_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel"
	"github.com/orsinium-labs/sequel/pgext"
)

func Test(t *testing.T) {
	is := is.New(t)

	type User struct {
		name string
		age  int
	}
	u := User{}
	q := sequel.Select(&u, &u.name, &u.age)
	q = q.Where(sequel.Gt(&u.age, 18))
	q = q.Where(sequel.GtF(&u.age, pgext.Abs(-18)))
	is.Equal(q.String(), "SELECT name, age FROM user WHERE (age > ? AND age > abs(?))")
}
