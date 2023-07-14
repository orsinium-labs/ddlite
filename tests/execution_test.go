package tests

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/matryer/is"
	_ "github.com/mattn/go-sqlite3"
	"github.com/orsinium-labs/qb"
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

	schema := qb.CreateTable(
		&p,
		qb.ColumnDef(&p.Country, qb.Text()),
		qb.ColumnDef(&p.City, qb.Text()).Null(),
		qb.ColumnDef(&p.TelCode, qb.Integer()),
	)
	schemaSQL, err := schema.SQL()
	is.NoErr(err)
	db.MustExec(schemaSQL)
	tx := db.MustBegin()
	defer func() {
		is.NoErr(tx.Rollback())
	}()

	tx.MustExec(
		"INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)",
		"United States", "New York", "1")
	tx.MustExec(
		"INSERT INTO place (country, telcode) VALUES ($1, $2)",
		"Hong Kong", "852")

	q := qb.Select(&p, &p.City, &p.Country).Where(qb.E(&p.TelCode, 1))
	r, err := qb.FetchOne[Place](tx, q)
	is.NoErr(err)
	is.Equal(r.City, "New York")
}
