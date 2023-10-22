package sequel_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/matryer/is"
	_ "github.com/mattn/go-sqlite3"
	"github.com/orsinium-labs/sequel"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dbtypes"
	"github.com/orsinium-labs/sequel/qb"
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
	schema := qb.CreateTable(
		&p,
		qb.ColumnDef(&p.Country, dbtypes.Text[string]()),
		qb.ColumnDef(&p.City, dbtypes.Text[string]()).Null(),
		qb.ColumnDef(&p.TelCode, dbtypes.Int32[int]()),
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
		qb.Insert(&p, &p.Country, &p.City, &p.TelCode).Values(
			Place{"United States", "New York", 1},
		),
	)
	is.NoErr(err)

	// INSERT
	_, err = sequel.Exec(conf, tx,
		qb.Insert(&p, &p.Country, &p.TelCode).Values(
			Place{Country: "Hong Kong", TelCode: 852},
		),
	)
	is.NoErr(err)

	// SELECT
	q := qb.Select(&p, &p.City, &p.Country).Where(qb.E(&p.TelCode, 1))
	r, err := sequel.FetchOne[Place](conf, tx, q)
	is.NoErr(err)
	is.Equal(r.City, "New York")
}
