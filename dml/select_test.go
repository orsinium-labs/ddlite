package dml_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dbfuncs"
	"github.com/orsinium-labs/sequel/dialects"
	"github.com/orsinium-labs/sequel/dml"
)

func TestSelectString(t *testing.T) {
	is := is.New(t)

	type User struct {
		name string
		age  int
	}
	u := User{}
	q := dml.Select(&u, &u.name, &u.age)
	q = q.Where(dml.Gt(dml.C(&u.age), dml.V(18)))
	q = q.And(dml.Gt(dml.C(&u.age), dbfuncs.Abs(dml.V(-18))))
	conf := dbconf.New(dialects.SQLite)
	sql, _, err := sequel.SQL(conf, q)
	is.NoErr(err)
	is.Equal(sql, "SELECT name, age FROM user WHERE age > ? AND age > abs(?)")
}
