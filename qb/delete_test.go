package qb_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel"
	"github.com/orsinium-labs/sequel/qb"
)

func TestDeleteSmoke(t *testing.T) {
	is := is.New(t)

	type User struct {
		name string
	}
	u := User{}
	q := qb.Delete(&u).Where(qb.E(&u.name, "Aragorn"))
	sql, _, err := sequel.SQL(q)
	is.NoErr(err)
	is.Equal(sql, "DELETE FROM user WHERE (name = $1)")
}
