package dialects

import "fmt"

type Dialect interface {
	fmt.Stringer

	// Precedence of operators and keywords.
	//
	// The precedence is used when generating SQL for expressions
	//  to add parenthesis to sub-expressions to avoid ambiguity.
	//
	// If precedence for the given operator is unknown, zero (the lowest precedence)
	// should be returned. In this case, the operation is almost always wrapped
	// in parenthesis.
	Precedence(string) uint8

	// Placeholder for variable binding.
	Placeholder(pos int) string

	// True is TRUE literal that can be assigned to a boolean field.
	//
	// Typically, if the database doesn't support bool data type, 1 is used instead.
	True() string

	// False is FALSE literal that can be assigned to a boolean field.
	//
	// Typically, if the database doesn't support bool data type, 0 is used instead.
	False() string

	// Int is data type that can fit an integer value of the given maximum size in bits.
	//
	// Bits indicate not the maximum allowed value but the maximum size in bits needed
	// to store it. One bit is always used to store the sign.
	// That is, Int(8) fits numbers only up to 2^7-1=127.
	//
	// The Go type int8 is equivalent to the DB type Int(8), int16 to Int(16), etc.
	Int(bits uint8) string

	UInt(bits uint8) string

	Interval() string

	Date() string
}
