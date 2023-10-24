package dml_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dml"
)

func TestDeleteSmoke(t *testing.T) {
	is := is.New(t)

	type User struct {
		name string
	}
	u := User{}
	q := dml.Delete(&u).Where(dml.E(&u.name, "Aragorn"))
	conf := dbconf.New("postgres")
	sql, _, err := dml.SQL(conf, q)
	is.NoErr(err)
	is.Equal(sql, "DELETE FROM user WHERE name = $1")
}
