package dml_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dml"
)

func TestInsertSQL(t *testing.T) {
	is := is.New(t)

	type User struct {
		Name string
		Age  int
	}
	u := User{}
	q := dml.Insert(&u, &u.Name, &u.Age)
	q = q.Values(User{"Aragorn", 88})
	conf := dbconf.New("postgres")
	sql, _, err := dml.SQL(conf, q)
	is.NoErr(err)
	// is.Equal(args, []any{"Aragorn", 88})
	is.Equal(sql, "INSERT INTO user (name, age) VALUES ($1, $2)")
}
