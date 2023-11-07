package ddl

import (
	"github.com/orsinium-labs/sequel-ddl/dialects"
)

type ColumnType interface {
	SQL(dialects.Dialect) dialects.DataType
}

type colType struct {
	callback func(dialects.Dialect) dialects.DataType
}

func (c colType) SQL(dialect dialects.Dialect) dialects.DataType {
	return c.callback(dialect)
}

// Bool is a boolean type.
//
// If the database doesn't support BOOL natively, the smallest integer type is used.
func Bool() ColumnType {
	callback := func(dialect dialects.Dialect) dialects.DataType {
		return dialect.Bool()
	}
	return colType{callback}
}
