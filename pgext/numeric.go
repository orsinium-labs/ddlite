package pgext

import "github.com/orsinium-labs/sequel"

type Args []any

func Abs(val int) sequel.Func[int] {
	return sequel.Func[int]{
		Name: "abs",
		Args: []any{val},
	}
}
