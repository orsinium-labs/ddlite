package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
)

type ClauseDataType func(dialects.Dialect) dialects.DataType
type dl = dialects.Dialect
type dt = dialects.DataType

// Bool is a boolean type.
//
// If the database doesn't support boolean data type natively,
// the smallest integer type is used.
func Bool() ClauseDataType {
	return func(dialect dl) dt { return dialect.Bool() }
}
