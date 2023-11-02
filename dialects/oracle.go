package dialects

import (
	"fmt"
	"math"
	"strconv"
)

var Oracle Dialect = oracle{}

type oracle struct{}

// Placeholder implements [Dialect].
func (oracle) Placeholder(pos int) string {
	return ":" + strconv.Itoa(pos+1)
}

// Precedence implements [Dialect].
func (oracle) Precedence(op string) uint8 {
	return oraclePrec[op]
}

func (oracle) True() string {
	return "1"
}

func (oracle) False() string {
	return "0"
}

func (oracle) Int(bits uint8) string {
	// https://docs.oracle.com/en/database/oracle/oracle-database/23/sqlrf/Data-Types.html
	digits := int(math.Log10(math.Pow(2, float64(bits))))
	return fmt.Sprintf("NUMBER(%d)", digits)
}

func (oracle) UInt(bits uint8) string {
	return Oracle.Int(bits + 1)
}

func (oracle) Interval() string {
	return "INTERVAL"
}

func (oracle) String() string {
	return "Oracle"
}

// https://docs.oracle.com/en//database/oracle/oracle-database/21/sqlrf/About-SQL-Operators.html
// https://docs.oracle.com/en//database/oracle/oracle-database/21/sqlrf/About-SQL-Conditions.html
var oraclePrec = map[string]uint8{
	"COLLATE": 20,
	"PRIOR":   20,

	"*": 18,
	"/": 18,

	"+":  17,
	"-":  17,
	"||": 17,

	"=":  16,
	"!=": 16,
	"<":  16,
	">":  16,
	"<=": 16,
	">=": 16,
	"<>": 16, // is it supported?

	"IS":          15,
	"IS NULL":     15,
	"IS NOT NULL": 15,
	"IS NOT":      15,
	"LIKE":        15,
	"NOT LIKE":    15, // is it supported?
	"BETWEEN":     15,
	"NOT BETWEEN": 15,
	"IN":          15,
	"NOT IN":      15,
	"EXISTS":      15,
	"IS OF":       15,

	"NOT": 14,
	"AND": 13,
	"OR":  12,
}
