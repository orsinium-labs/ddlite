package dbconf_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dbconf"
)

func Test_Placeholder_Make(t *testing.T) {
	tests := []struct {
		ph   dbconf.Placeholder
		pos  int
		want string
	}{
		{dbconf.Question, 0, "?"},
		{dbconf.Question, 1, "?"},
		{dbconf.Question, 2, "?"},
		{dbconf.Dollar, 0, "$1"},
		{dbconf.Dollar, 1, "$2"},
		{dbconf.Dollar, 2, "$3"},
		{dbconf.Colon, 0, ":1"},
		{dbconf.Colon, 1, ":2"},
		{dbconf.Colon, 2, ":3"},
		{dbconf.AtP, 0, "@p1"},
		{dbconf.AtP, 1, "@p2"},
		{dbconf.AtP, 2, "@p3"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			is := is.New(t)
			got := tt.ph.Make(tt.pos)
			is.Equal(got, tt.want)
		})
	}
}
