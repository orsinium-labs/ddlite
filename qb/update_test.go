package qb_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/qb"
)

func TestUpdateSQL(t *testing.T) {
	is := is.New(t)

	type User struct {
		name string
		age  int
	}
	u := User{}
	q := qb.Update(&u, qb.Set(&u.age, qb.V(88)))
	q = q.Where(qb.E(&u.name, "Aragorn"))
	conf := dbconf.New("postgres")
	sql, _, err := sequel.SQL(conf, q)
	is.NoErr(err)
	// is.Equal(args, []any{88, "Aragorn"})
	is.Equal(sql, "UPDATE user SET age = $1 WHERE (name = $2)")
}
