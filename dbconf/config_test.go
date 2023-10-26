package dbconf_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dialects"
)

func TestCamelToSnake(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	f := dbconf.CamelToSnake
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
	f := dbconf.SnakeToCamel
	is.Equal(f("user"), "User")
	is.Equal(f("user-name"), "UserName")
	is.Equal(f("user_name"), "UserName")
	is.Equal(f("url"), "Url")
	is.Equal(f("url_addr"), "UrlAddr")
	is.Equal(f("a_user"), "AUser")
}

func TestConfig_WithModel(t *testing.T) {
	t.Parallel()
	is := is.New(t)

	type User struct{}
	u1 := User{}
	u2 := User{}

	c1 := dbconf.New(dialects.SQLite)
	c2 := c1.WithModel(u1)
	c3 := c2.WithModel(u2)
	c4 := c1.WithModel(u2)

	is.Equal(len(c1.Models), 0)
	is.Equal(len(c2.Models), 1)
	is.Equal(len(c3.Models), 2)
	is.Equal(len(c4.Models), 1)
}
