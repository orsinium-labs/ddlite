package dbfuncs_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/dbfuncs"
	"github.com/orsinium-labs/sequel/dialects"
	"github.com/orsinium-labs/sequel/dml"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type Squirrler interface {
	Tokens(dbconf.Config) tokens.Tokens
}

func TestNumericFuncsSQL(t *testing.T) {
	testCases := []struct {
		given Squirrler
		sql   string
		args  []any
	}{
		{
			given: dbfuncs.Abs(dml.V(12)),
			sql:   "abs(?)",
			args:  []any{12},
		},
		{
			given: dbfuncs.Ceil[int](dml.V(12.0)),
			sql:   "ceil(?)",
			args:  []any{12.0},
		},
		{
			given: dbfuncs.Div[int](dml.V(12), dml.V(5)),
			sql:   "div(?, ?)",
			args:  []any{12, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.sql, func(t *testing.T) {
			is := is.New(t)
			conf := dbconf.New(dialects.SQLite)
			sql, args, err := tc.given.Tokens(conf).SQL(conf)
			is.NoErr(err)
			is.Equal(sql, tc.sql)
			is.Equal(args, tc.args)
		})
	}
}
