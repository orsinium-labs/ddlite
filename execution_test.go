package ddl_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/matryer/is"
	_ "github.com/mattn/go-sqlite3"
	ddl "github.com/orsinium-labs/sequel"
	"github.com/orsinium-labs/sequel/dbtypes"
	"github.com/orsinium-labs/sequel/dialects"
)

func TestFetchOne(t *testing.T) {
	is := is.New(t)
	db, err := sqlx.Connect("sqlite3", ":memory:")
	is.NoErr(err)

	// CREATE TABLE
	schema := ddl.CreateTable(
		"place",
		ddl.Column("country", dbtypes.Text()),
		ddl.Column("city", dbtypes.Text()).Null(),
		ddl.Column("tel_code", dbtypes.Int(32)),
	)
	_, err = ddl.Exec(dialects.SQLite, db, schema)
	is.NoErr(err)
	tx := db.MustBegin()
	defer func() {
		is.NoErr(tx.Rollback())
	}()
}
