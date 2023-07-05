package sequel_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel"
)

func Test(t *testing.T) {
	is := is.New(t)

	type User struct {
		name string
		age  int
	}
	u := User{}
	q := sequel.Select(&u, &u.name, &u.age)
	is.Equal(q.String(), "SELECT name, age FROM user")
}
