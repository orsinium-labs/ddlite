package dml_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dml"
)

func TestUpdateSQL(t *testing.T) {
	is := is.New(t)

	type User struct {
		name string
		age  int
	}
	u := User{}
	q := dml.Update(&u, dml.Set(&u.age, dml.V(88)))
	q = q.Where(dml.E(&u.name, "Aragorn"))
	conf := dbconf.New("postgres")
	sql, _, err := sequel.SQL(conf, q)
	is.NoErr(err)
	// is.Equal(args, []any{88, "Aragorn"})
	is.Equal(sql, "UPDATE user SET age = $1 WHERE name = $2")
}
