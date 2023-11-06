package dialects

import "fmt"

type Dialect interface {
	fmt.Stringer

	// Int is data type that can fit an integer value of the given maximum size in bits.
	//
	// Bits indicate not the maximum allowed value but the maximum size in bits needed
	// to store it. One bit is always used to store the sign.
	// That is, Int(8) fits numbers only up to 2^7-1=127.
	//
	// The Go type int8 is equivalent to the DB type Int(8), int16 to Int(16), etc.
	Int(bits uint8) string

	UInt(bits uint8) string

	Float(precision uint8) string

	Interval() string

	Date() string
}
