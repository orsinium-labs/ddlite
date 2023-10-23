package sequel_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/matryer/is"
	_ "github.com/mattn/go-sqlite3"
	"github.com/orsinium-labs/sequel"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dbtypes"
	"github.com/orsinium-labs/sequel/ddl"
	"github.com/orsinium-labs/sequel/dml"
)

func TestFetchOne(t *testing.T) {
	is := is.New(t)
	db, err := sqlx.Connect("sqlite3", ":memory:")
	is.NoErr(err)

	type Place struct {
		Country string
		City    string
		TelCode int
	}
	p := Place{}

	// CREATE TABLE
	schema := ddl.CreateTable(
		"place",
		ddl.ColumnDef("country", dbtypes.Text()),
		ddl.ColumnDef("city", dbtypes.Text()).Null(),
		ddl.ColumnDef("tel_code", dbtypes.Int32()),
	)
	conf := dbconf.New("sqlite3")
	_, err = sequel.Exec(conf, db, schema)
	is.NoErr(err)
	tx := db.MustBegin()
	defer func() {
		is.NoErr(tx.Rollback())
	}()

	// INSERT
	_, err = sequel.Exec(conf, tx,
		dml.Insert(&p, &p.Country, &p.City, &p.TelCode).Values(
			Place{"United States", "New York", 1},
		),
	)
	is.NoErr(err)

	// INSERT
	_, err = sequel.Exec(conf, tx,
		dml.Insert(&p, &p.Country, &p.TelCode).Values(
			Place{Country: "Hong Kong", TelCode: 852},
		),
	)
	is.NoErr(err)

	// SELECT
	q := dml.Select(&p, &p.City, &p.Country).Where(dml.E(&p.TelCode, 1))
	r, err := sequel.FetchOne[Place](conf, tx, q)
	is.NoErr(err)
	is.Equal(r.City, "New York")
}
