package dbfuncs

type Args []any

func Abs(val int) Func[int] {
	return Func[int]{
		Name: "abs",
		Args: []any{val},
	}
}
