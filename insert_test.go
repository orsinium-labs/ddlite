package qb_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/qb"
)

func TestInsertSQL(t *testing.T) {
	is := is.New(t)

	type User struct {
		Name string
		Age  int
	}
	u := User{}
	q := qb.Insert(&u, &u.Name, &u.Age)
	q = q.Values(User{"Aragorn", 88})
	sql, _, err := qb.SQL(q)
	is.NoErr(err)
	// is.Equal(args, []any{"Aragorn", 88})
	is.Equal(sql, "INSERT INTO user (Name,Age) VALUES ($1,$2)")
}
