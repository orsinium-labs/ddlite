package ddl_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel-ddl/ddl"
)

func TestCreateTable(t *testing.T) {
	is := is.New(t)
	q := ddl.CreateTable("user",
		ddl.Column("name", ddl.Text, ddl.Null),
		ddl.Column("age", ddl.Integer, ddl.Null),
	).Constraints(
		ddl.Constraint("", ddl.Unique(), "name"),
	)
	sql, err := ddl.SQL(q)
	is.NoErr(err)
	is.Equal(sql, "CREATE TABLE user (name TEXT, age INTEGER, UNIQUE (name))")
}
