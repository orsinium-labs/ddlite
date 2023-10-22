// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package constraints defines a set of useful constraints to be used
// with type parameters.
package constraints

import (
	"database/sql"
	"database/sql/driver"
)

// Option is a Maybe monad used to represent a NULLable field.
//
// Compatible with Option defined in the `mo` package.
//
// https://github.com/samber/mo
type Option[T any] interface {
	Get() (T, bool)
}

// Signed is a constraint that permits any signed integer type.
// If future releases of Go add new predeclared signed integer types,
// this constraint will be modified to include them.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
// If future releases of Go add new predeclared unsigned integer types,
// this constraint will be modified to include them.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint that permits any integer type.
// If future releases of Go add new predeclared integer types,
// this constraint will be modified to include them.
type Integer interface {
	Signed | Unsigned
}

// Float is a constraint that permits any floating-point type.
// If future releases of Go add new predeclared floating-point types,
// this constraint will be modified to include them.
type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface {
	Integer | Float | ~string
}

// Decimal is an arbitrary precision fixed-point decimal.
//
// Supports [shopspring/decimal] and [ericlagergren/decimal].
//
// [shopspring/decimal]: https://pkg.go.dev/github.com/shopspring/decimal
// [ericlagergren/decimal]: https://pkg.go.dev/github.com/ericlagergren/decimal
type Decimal interface {
	sql.Scanner
	driver.Valuer
	Sign() int
}
