package ddl_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/matryer/is"
	_ "github.com/mattn/go-sqlite3"
	"github.com/orsinium-labs/sequel-ddl/ddl"
	"github.com/orsinium-labs/sequel-ddl/dialects"
)

func TestFetchOne(t *testing.T) {
	is := is.New(t)
	db, err := sqlx.Connect("sqlite3", ":memory:")
	is.NoErr(err)

	// CREATE TABLE
	schema := ddl.CreateTable(
		"place",
		ddl.Column("country", ddl.Text(), ddl.NotNull),
		ddl.Column("city", ddl.Text(), ddl.Null),
		ddl.Column("tel_code", ddl.Int(32), ddl.NotNull),
	)
	_, err = ddl.Exec(dialects.SQLite, db, schema)
	is.NoErr(err)
	tx := db.MustBegin()
	defer func() {
		is.NoErr(tx.Rollback())
	}()
}
