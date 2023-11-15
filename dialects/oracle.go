package dialects

var Oracle Dialect = oracle{}

type oracle struct{}

func (oracle) Features() Features {
	return oracleFeatures
}

func (oracle) Int(bits uint8) DataType {
	// https://docs.oracle.com/en/database/oracle/oracle-database/23/sqlrf/Data-Types.html
	return call("NUMBER", bits)
}

func (oracle) UInt(bits uint8) DataType {
	return ""
}

func (oracle) Float(precision uint8) DataType {
	if precision > 53 {
		return ""
	}
	if precision == 24 {
		return "FLOAT(63)"
	}
	if precision == 53 {
		return "FLOAT"
	}
	return call("FLOAT", precision)
}

func (oracle) Decimal(precision uint8, scale uint8) DataType {
	return call2("NUMBER", precision, scale)
}

func (oracle) Text() DataType {
	return ""
}

func (oracle) Blob() DataType {
	return "BLOB"
}

func (oracle) FixedChar(size uint32) DataType {
	return call("CHAR", size)
}

func (oracle) VarChar(size uint32) DataType {
	return call("VARCHAR2", size)
}

func (oracle) Enum(members []string) DataType {
	return ""
}

func (oracle) Interval() DataType {
	return "INTERVAL"
}

func (oracle) Date() DataType {
	return "DATE"
}

func (oracle) DateTime() DataType {
	return "TIMESTAMP"
}

func (oracle) Time() DataType {
	return "INTERVAL"
}

func (oracle) Bool() DataType {
	return "NUMBER(1)"
}

func (oracle) String() string {
	return "Oracle"
}

var oracleFeatures = Features{
	"DELETE FROM":  true,
	"CREATE TABLE": true,
	"DROP TABLE":   true,

	// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/TRUNCATE-TABLE.html
	"TRUNCATE TABLE": true,

	// https://docs.oracle.com/database/121/SQLRF/statements_3001.htm
	"ALTER TABLE":                       true,
	"ALTER TABLE / RENAME TO":           true,
	"ALTER TABLE / ADD":                 true, // ADD COLUMN or ADD CONSTRAINT
	"ALTER TABLE / MODIFY":              true,
	"ALTER TABLE / MODIFY / DEFAULT":    true,
	"ALTER TABLE / MODIFY / CONSTRAINT": true,
	"ALTER TABLE / DROP COLUMN":         true,
	"ALTER TABLE / RENAME COLUMN":       true,
	"ALTER TABLE / MODIFY CONSTRAINT":   true,
	"ALTER TABLE / MODIFY PRIMARY KEY":  true,
	"ALTER TABLE / MODIFY UNIQUE":       true,
	"ALTER TABLE / RENAME CONSTRAINT":   true,
	"ALTER TABLE / DROP CONSTRAINT":     true,
	"ALTER TABLE / DROP PRIMARY KEY":    true,
	"ALTER TABLE / DROP UNIQUE":         true,
}
