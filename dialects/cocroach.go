package dialects

var CocroachDB Dialect = cocroach{}

type cocroach struct{}

func (cocroach) Features() Features {
	return cocroachFeatures
}

func (cocroach) Int(bits uint8) DataType {
	// https://www.cockroachlabs.com/docs/v23.1/int
	if bits <= 16 {
		return "INT2"
	}
	if bits <= 32 {
		return "INT4"
	}
	if bits <= 64 {
		return "INT8"
	}
	return ""
}

func (cocroach) UInt(bits uint8) DataType {
	return ""
}

func (cocroach) Float(precision uint8) DataType {
	// https://www.cockroachlabs.com/docs/v23.1/float
	if precision <= 24 {
		return "REAL"
	}
	if precision <= 53 {
		return "DOUBLE PRECISION"
	}
	return ""
}

func (cocroach) Decimal(precision uint8, scale uint8) DataType {
	return call2("DECIMAL", precision, scale)
}

func (cocroach) Text() DataType {
	return "STRING"
}

func (cocroach) FixedChar(size uint32) DataType {
	return call("STRING", size)
}

func (cocroach) VarChar(size uint32) DataType {
	return call("STRING", size)
}

func (cocroach) Enum(members []string) DataType {
	return callVar("ENUM", members)
}

func (cocroach) Blob() DataType {
	return "BYTES"
}

func (cocroach) Interval() DataType {
	return "INTERVAL"
}

func (cocroach) Date() DataType {
	return "DATE"
}

func (cocroach) DateTime() DataType {
	return "TIMESTAMP"
}

func (cocroach) Time() DataType {
	return "TIME"
}

func (cocroach) Bool() DataType {
	return "BOOL"
}

func (cocroach) String() string {
	return "CocroachDB"
}

var cocroachFeatures = Features{
	"DELETE FROM":  true,
	"CREATE TABLE": true,
	"DROP TABLE":   true,

	// https://www.cockroachlabs.com/docs/stable/truncate
	"TRUNCATE TABLE": true,

	// https://www.cockroachlabs.com/docs/stable/alter-table
	"ALTER TABLE":                                 true,
	"ALTER TABLE / IF EXISTS":                     true,
	"ALTER TABLE / RENAME TO":                     true,
	"ALTER TABLE / SET SCHEMA":                    true,
	"ALTER TABLE / OWNER TO":                      true,
	"ALTER TABLE / SPLIT AT":                      true,
	"ALTER TABLE / ADD COLUMN":                    true,
	"ALTER TABLE / ADD CONSTRAINT":                true,
	"ALTER TABLE / RENAME COLUMN":                 true,
	"ALTER TABLE / RENAME CONSTRAINT":             true,
	"ALTER TABLE / ALTER COLUMN":                  true,
	"ALTER TABLE / ALTER COLUMN / SET DEFAULT":    true,
	"ALTER TABLE / ALTER COLUMN / SET ON UPDATE":  true,
	"ALTER TABLE / ALTER COLUMN / SET VISIBLE":    true,
	"ALTER TABLE / ALTER COLUMN / SET NOT NULL":   true,
	"ALTER TABLE / ALTER COLUMN / DROP DEFAULT":   true,
	"ALTER TABLE / ALTER COLUMN / DROP ON UPDATE": true,
	"ALTER TABLE / ALTER COLUMN / DROP NOT NULL":  true,
	"ALTER TABLE / ALTER COLUMN / DROP STORED":    true,
	"ALTER TABLE / ALTER COLUMN / TYPE":           true,
	"ALTER TABLE / DROP COLUMN":                   true,
	"ALTER TABLE / DROP CONSTRAINT":               true,
	"ALTER TABLE / VALIDATE CONSTRAINT":           true,
}
