package qb_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel"
	"github.com/orsinium-labs/sequel/dbconfig"
	"github.com/orsinium-labs/sequel/qb"
)

func TestDeleteSmoke(t *testing.T) {
	is := is.New(t)

	type User struct {
		name string
	}
	u := User{}
	q := qb.Delete(&u).Where(qb.E(&u.name, "Aragorn"))
	conf := dbconfig.New("postgres")
	sql, _, err := sequel.SQL(conf, q)
	is.NoErr(err)
	is.Equal(sql, "DELETE FROM user WHERE (name = $1)")
}
