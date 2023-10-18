package dbtypes

import (
	"fmt"
	"time"
)

type ColumnType[T any] interface {
	Default() T
	SQL() string
}

type colType0[T any] struct {
	name string
}

func (c colType0[T]) Default() T {
	return *new(T)
}

func (c colType0[T]) SQL() string {
	return c.name
}

type colType1[T any] struct {
	name string
	arg  int
}

func (c colType1[T]) Default() T {
	return *new(T)
}

func (c colType1[T]) SQL() string {
	return fmt.Sprintf("%s(%d)", c.name, c.arg)
}

// -- NUMERIC -- //

// SMALLINT, small-range integer
func SmallInt() ColumnType[int] {
	return colType0[int]{"SMALLINT"}
}

// INTEGER, typical choice for integer
func Integer() ColumnType[int] {
	return colType0[int]{"INTEGER"}
}

// BIGINT, large-range integer
func BigInt() ColumnType[int] {
	return colType0[int]{"BIGINT"}
}

// REAL, variable-precision, inexact
func Real() ColumnType[float64] {
	return colType0[float64]{"REAL"}
}

// SMALLSERIAL, small autoincrementing integer
func SmallSerial() ColumnType[int] {
	return colType0[int]{"SMALLSERIAL"}
}

// SERIAL, autoincrementing integer
func Serial() ColumnType[int] {
	return colType0[int]{"SERIAL"}
}

// BIGSERIAL, large autoincrementing integer
func BigSerial() ColumnType[int] {
	return colType0[int]{"BIGSERIAL"}
}

// -- CHARACTER -- //

// TEXT, variable unlimited length
func Text() ColumnType[string] {
	return colType0[string]{"TEXT"}
}

// CHARACTER, fixed-length, blank padded
func Character(n int) ColumnType[string] {
	return colType1[string]{"CHARACTER", n}
}

// VARCHAR, variable-length with limit
func VarChar(n int) ColumnType[string] {
	return colType1[string]{"VARCHAR", n}
}

// -- MISC -- //

// BOOLEAN, state of true or false
func Boolean() ColumnType[bool] {
	return colType0[bool]{"BOOLEAN"}
}

// DATE, date (no time of day)
func Date() ColumnType[time.Time] {
	return colType0[time.Time]{"DATE"}
}

// TIMESTAMP, both date and time
func TimeStamp() ColumnType[time.Time] {
	return colType0[time.Time]{"TIMESTAMP"}
}
