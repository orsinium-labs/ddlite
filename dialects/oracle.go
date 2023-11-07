package dialects

import (
	"fmt"
	"math"
	"strconv"
)

var Oracle Dialect = oracle{}

type oracle struct{}

func (oracle) Int(bits uint8) string {
	// https://docs.oracle.com/en/database/oracle/oracle-database/23/sqlrf/Data-Types.html
	digits := int(math.Log10(math.Pow(2, float64(bits))))
	return fmt.Sprintf("NUMBER(%d)", digits)
}

func (oracle) UInt(bits uint8) string {
	return Oracle.Int(bits + 1)
}

func (oracle) Float(precision uint8) string {
	if precision > 53 {
		return ""
	}
	if precision == 24 {
		return "FLOAT(63)"
	}
	if precision == 53 {
		return "FLOAT"
	}
	return "FLOAT(" + strconv.FormatInt(int64(precision), 10) + ")"
}

func (oracle) Decimal(precision uint8, scale uint8) string {
	return call2("NUMBER", precision, scale)
}

func (oracle) Text() string {
	return ""
}

func (oracle) Interval() string {
	return "INTERVAL"
}

func (oracle) Date() string {
	return "DATE"
}

func (oracle) String() string {
	return "Oracle"
}
