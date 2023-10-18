package dbfuncs_test

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconfig"
	"github.com/orsinium-labs/sequel/dbfuncs"
	"github.com/orsinium-labs/sequel/qb"
)

type Squirrler interface {
	Squirrel(dbconfig.Config) squirrel.Sqlizer
}

func TestNumericFuncsSQL(t *testing.T) {
	testCases := []struct {
		given Squirrler
		sql   string
		args  []any
	}{
		{
			given: dbfuncs.Abs(qb.V(12)),
			sql:   "abs(?)",
			args:  []any{12},
		},
		{
			given: dbfuncs.Ceil[int](qb.V(12.0)),
			sql:   "ceil(?)",
			args:  []any{12.0},
		},
		{
			given: dbfuncs.Div[int](qb.V(12), qb.V(5)),
			sql:   "div(?, ?)",
			args:  []any{12, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.sql, func(t *testing.T) {
			is := is.New(t)
			conf := dbconfig.New("sqlite3")
			sqlizer := tc.given.Squirrel(conf)
			sql, args, err := sqlizer.ToSql()
			is.NoErr(err)
			is.Equal(sql, tc.sql)
			is.Equal(args, tc.args)
		})
	}
}
