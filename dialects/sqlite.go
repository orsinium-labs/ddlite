package dialects

var SQLite Dialect = sqlite{}

type sqlite struct{}

func (sqlite) Features() Features {
	return sqliteFeatures
}

func (sqlite) Int(bits uint8) DataType {
	// https://www.sqlite.org/datatype3.html#boolean_datatype
	// INTEGER fits up to 8 bytes.
	return "INTEGER"
}

func (sqlite) UInt(bits uint8) DataType {
	return "INTEGER"
}

func (sqlite) Float(precision uint8) DataType {
	if precision <= 53 {
		return "REAL"
	}
	return ""
}

func (sqlite) Decimal(precision uint8, scale uint8) DataType {
	return "NUMERIC"
}

func (sqlite) Text() DataType {
	return "TEXT"
}

func (sqlite) FixedChar(size uint32) DataType {
	return "TEXT"
}

func (sqlite) VarChar(size uint32) DataType {
	return "TEXT"
}

func (sqlite) Enum(members []string) DataType {
	return "TEXT"
}

func (sqlite) Blob() DataType {
	return "BLOB"
}

func (sqlite) Interval() DataType {
	return ""
}

func (sqlite) Date() DataType {
	return ""
}

func (sqlite) DateTime() DataType {
	return ""
}

func (sqlite) Time() DataType {
	return "TIME"
}

func (sqlite) Bool() DataType {
	return "INTEGER"
}

func (sqlite) String() string {
	return "SQLite"
}

var sqliteFeatures = Features{
	// https://www.sqlite.org/lang_delete.html#the_truncate_optimization
	"DELETE FROM":    true,
	"TRUNCATE TABLE": false,

	// https://www.sqlite.org/lang_altertable.html
	"ALTER TABLE":                 true,
	"ALTER TABLE / RENAME TO":     true,
	"ALTER TABLE / RENAME COLUMN": true,
	"ALTER TABLE / ADD COLUMN":    true,
	"ALTER TABLE / DROP COLUMN":   true,
}
