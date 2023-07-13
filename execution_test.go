package sequel_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/matryer/is"
	_ "github.com/mattn/go-sqlite3"
	sq "github.com/orsinium-labs/sequel"
)

func TestFetchOne(t *testing.T) {
	is := is.New(t)
	db, err := sqlx.Connect("sqlite3", ":memory:")
	is.NoErr(err)
	var schema = `
		CREATE TABLE place (
			Country text,
			City text NULL,
			TelCode integer
		);
	`
	db.MustExec(schema)
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

	type Place struct {
		Country string
		City    string
		TelCode int
	}
	p := Place{}
	q := sq.Select(&p, &p.City, &p.Country).Where(sq.Eq(sq.C(&p.TelCode), sq.V(1)))
	r, err := sq.FetchOne[Place](tx, q)
	is.NoErr(err)
	is.Equal(r.City, "New York")
}
