package pgext_test

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/matryer/is"
	sq "github.com/orsinium-labs/sequel"
	"github.com/orsinium-labs/sequel/pgext"
)

type Squirrler interface {
	Squirrel(...sq.Model) squirrel.Sqlizer
}

func TestNumericFuncsSQL(t *testing.T) {
	testCases := []struct {
		given Squirrler
		sql   string
		args  []any
	}{
		{
			given: pgext.Abs(sq.V(12)),
			sql:   "abs(?)",
			args:  []any{12},
		},
		{
			given: pgext.Ceil[float64, int](sq.V(12.0)),
			sql:   "ceil(?)",
			args:  []any{12.0},
		},
		{
			given: pgext.Div[int, int](sq.V(12), sq.V(5)),
			sql:   "div(?, ?)",
			args:  []any{12, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.sql, func(t *testing.T) {
			is := is.New(t)
			sqlizer := tc.given.Squirrel()
			sql, args, err := sqlizer.ToSql()
			is.NoErr(err)
			is.Equal(sql, tc.sql)
			is.Equal(args, tc.args)
		})
	}
}
