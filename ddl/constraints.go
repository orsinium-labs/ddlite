package ddl

import (
	"errors"

	"github.com/orsinium-labs/sequel-ddl/dialects"
	"github.com/orsinium-labs/sequel-ddl/internal/tokens"
)

type Constraint interface {
	isConstraint()
	tokens(dialect dialects.Dialect) tokens.Tokens
}

type tUnique struct {
	names []Safe
}

func Unique(names ...Safe) Constraint {
	return tUnique{names: names}
}

func (def tUnique) isConstraint() {}

func (def tUnique) tokens(dialects.Dialect) tokens.Tokens {
	if len(def.names) == 0 {
		err := errors.New("unique index must have at least one column specified")
		return tokens.New(tokens.Err(err))
	}
	ts := tokens.New(
		tokens.Keyword("UNIQUE"),
		tokens.LParen(),
		tokens.Raws(def.names...),
		tokens.RParen(),
	)
	return ts
}
