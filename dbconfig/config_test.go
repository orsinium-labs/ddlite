package dbconfig_test

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconfig"
)

func TestCamelToSnake(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	f := dbconfig.CamelToSnake
	is.Equal(f("User"), "user")
	is.Equal(f("UserName"), "user_name")
	is.Equal(f("Url"), "url")
	is.Equal(f("URL"), "url")
	is.Equal(f("URLAddr"), "url_addr")
	is.Equal(f("IPAddr"), "ip_addr")
	is.Equal(f("AUser"), "a_user")
}

func TestSnakeToCamel(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	f := dbconfig.SnakeToCamel
	is.Equal(f("user"), "User")
	is.Equal(f("user-name"), "UserName")
	is.Equal(f("user_name"), "UserName")
	is.Equal(f("url"), "Url")
	is.Equal(f("url_addr"), "UrlAddr")
	is.Equal(f("a_user"), "AUser")
}

func TestNew_Placeholder(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	n := dbconfig.New
	is.Equal(n("sqlite3").Placeholder, dbconfig.Question)
	is.Equal(n("sqlite").Placeholder, dbconfig.Question)
	is.Equal(n("mysql").Placeholder, dbconfig.Question)
	is.Equal(n("postgres").Placeholder, dbconfig.Dollar)
	is.Equal(n("goracle").Placeholder, dbconfig.Colon)
	is.Equal(n("sqlserver").Placeholder, dbconfig.AtP)
	is.Equal(n("SQLServer").Placeholder, dbconfig.AtP)
	is.Equal(n("").Placeholder, dbconfig.Question)
	is.Equal(n("unknown").Placeholder, dbconfig.Question)
}

func TestConfig_SquirrelPlaceholder(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	n := dbconfig.New
	is.Equal(n("sqlite3").SquirrelPlaceholder(), squirrel.Question)
	is.Equal(n("sqlite").SquirrelPlaceholder(), squirrel.Question)
	is.Equal(n("mysql").SquirrelPlaceholder(), squirrel.Question)
	is.Equal(n("postgres").SquirrelPlaceholder(), squirrel.Dollar)
	is.Equal(n("goracle").SquirrelPlaceholder(), squirrel.Colon)
	is.Equal(n("sqlserver").SquirrelPlaceholder(), squirrel.AtP)
	is.Equal(n("SQLServer").SquirrelPlaceholder(), squirrel.AtP)
	is.Equal(n("").SquirrelPlaceholder(), squirrel.Question)
	is.Equal(n("unknown").SquirrelPlaceholder(), squirrel.Question)

	c := n("unknown")
	c.Placeholder = 13
	is.Equal(c.SquirrelPlaceholder(), squirrel.Question)
}

func TestConfig_WithModel(t *testing.T) {
	t.Parallel()
	is := is.New(t)

	type User struct{}
	u1 := User{}
	u2 := User{}

	c1 := dbconfig.New("sqlite3")
	c2 := c1.WithModel(u1)
	c3 := c2.WithModel(u2)
	c4 := c1.WithModel(u2)

	is.Equal(len(c1.Models), 0)
	is.Equal(len(c2.Models), 1)
	is.Equal(len(c3.Models), 2)
	is.Equal(len(c4.Models), 1)
}
