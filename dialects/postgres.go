package dialects

var PostgreSQL Dialect = psql{}

type psql struct{}

func (psql) Features() Features {
	return psqlFeatures
}

func (psql) Int(bits uint8) DataType {
	// https://www.postgresql.org/docs/current/datatype-numeric.html
	if bits <= 16 {
		return "SMALLINT"
	}
	if bits <= 32 {
		return "INTEGER"
	}
	// BIGINT fits anything up to 2^8=256 which is +1 from what `bits` (uint8) can fit.
	return "BIGINT"
}

func (psql) UInt(bits uint8) DataType {
	return ""
}

func (psql) Float(precision uint8) DataType {
	// https://www.postgresql.org/docs/current/datatype-numeric.html#DATATYPE-FLOAT
	if precision > 53 {
		return ""
	}
	if precision == 24 {
		return "REAL"
	}
	if precision == 53 {
		return "DOUBLE PRECISION"
	}
	return call("FLOAT", precision)
}

func (psql) Decimal(precision uint8, scale uint8) DataType {
	return call2("NUMERIC", precision, scale)
}

func (psql) Text() DataType {
	return "TEXT"
}

func (psql) FixedChar(size uint32) DataType {
	return call("CHAR", size)
}

func (psql) VarChar(size uint32) DataType {
	return call("VARCHAR", size)
}

func (psql) Enum(members []string) DataType {
	return callVar("ENUM", members)
}

func (psql) Blob() DataType {
	return "BYTEA"
}

func (psql) Interval() DataType {
	return "INTERVAL"
}

func (psql) Date() DataType {
	return "DATE"
}

func (psql) DateTime() DataType {
	return "TIMESTAMP"
}

func (psql) Time() DataType {
	return "TIME"
}

func (psql) Bool() DataType {
	return "BOOLEAN"
}

func (psql) String() string {
	return "PostgreSQL"
}

var psqlFeatures = Features{
	"DELETE FROM": true,

	// https://www.postgresql.org/docs/current/sql-truncate.html
	"TRUNCATE TABLE": true,

	// https://www.postgresql.org/docs/current/sql-altertable.html
	"ALTER TABLE":                                  true,
	"ALTER TABLE / RENAME COLUMN":                  true,
	"ALTER TABLE / RENAME CONSTRAINT":              true,
	"ALTER TABLE / RENAME TO":                      true,
	"ALTER TABLE / SET SCHEMA":                     true,
	"ALTER TABLE / SET TABLESPACE":                 true,
	"ALTER TABLE / ATTACH PARTITION":               true,
	"ALTER TABLE / DETACH PARTITION":               true,
	"ALTER TABLE / ADD COLUMN":                     true,
	"ALTER TABLE / DROP COLUMN":                    true,
	"ALTER TABLE / ALTER COLUMN":                   true,
	"ALTER TABLE / ALTER COLUMN / TYPE":            true,
	"ALTER TABLE / ALTER COLUMN / SET DEFAULT":     true,
	"ALTER TABLE / ALTER COLUMN / DROP DEFAULT":    true,
	"ALTER TABLE / ALTER COLUMN / SET NOT NULL":    true,
	"ALTER TABLE / ALTER COLUMN / DROP NOT NULL":   true,
	"ALTER TABLE / ALTER COLUMN / DROP EXPRESSION": true,
	"ALTER TABLE / ALTER COLUMN / ADD GENERATED":   true,
	"ALTER TABLE / ALTER COLUMN / DROP IDENTITY":   true,
	"ALTER TABLE / ALTER COLUMN / SET STATISTICS":  true,
	"ALTER TABLE / ALTER COLUMN / RESET":           true,
	"ALTER TABLE / ALTER COLUMN / SET":             true,
	"ALTER TABLE / ADD":                            true,
	"ALTER TABLE / ALTER CONSTRAINT":               true,
	"ALTER TABLE / VALIDATE CONSTRAINT":            true,
	"ALTER TABLE / DROP CONSTRAINT":                true,
}
