package dbconf

import "strconv"

// Placeholder is the style of placeholder for variable binding.
type Placeholder interface {
	Make(position int) string
}

type (
	question rune
	dollar   rune
	colon    rune
	atp      rune
)

const (
	// Question is the default question placeholder (e.g. ?)
	Question question = '?'

	// Dollar is a dollar-prefixed positional placeholder (e.g. $1, $2, $3).
	Dollar dollar = '$'

	// Colon is a colon-prefixed positional placeholder (e.g. :1, :2, :3).
	Colon colon = ':'

	// AtP is a "@p"-prefixed positional placeholder (e.g. @p1, @p2, @p3).
	AtP atp = '@'
)

func (question) Make(pos int) string {
	return "?"
}

func (dollar) Make(pos int) string {
	return "$" + strconv.Itoa(pos+1)
}

func (colon) Make(pos int) string {
	return ":" + strconv.Itoa(pos+1)
}

func (atp) Make(pos int) string {
	return "@p" + strconv.Itoa(pos+1)
}
