package dialects_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/sequel/dialects"
)

func Test_Placeholder_Make(t *testing.T) {
	tests := []struct {
		ph   dialects.Placeholder
		pos  int
		want string
	}{
		{dialects.Question, 0, "?"},
		{dialects.Question, 1, "?"},
		{dialects.Question, 2, "?"},
		{dialects.Dollar, 0, "$1"},
		{dialects.Dollar, 1, "$2"},
		{dialects.Dollar, 2, "$3"},
		{dialects.Colon, 0, ":1"},
		{dialects.Colon, 1, ":2"},
		{dialects.Colon, 2, ":3"},
		{dialects.AtP, 0, "@p1"},
		{dialects.AtP, 1, "@p2"},
		{dialects.AtP, 2, "@p3"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			is := is.New(t)
			got := tt.ph.Make(tt.pos)
			is.Equal(got, tt.want)
		})
	}
}
