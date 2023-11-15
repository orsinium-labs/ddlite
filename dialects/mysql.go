package dialects

var MySQL Dialect = mysql{}

type mysql struct{}

func (mysql) Features() Features {
	return mysqlFeatures
}

func (mysql) Int(bits uint8) DataType {
	// https://dev.mysql.com/doc/refman/8.2/en/integer-types.html
	if bits <= 8 {
		return "TINYINT"
	}
	if bits <= 16 {
		return "SMALLINT"
	}
	if bits <= 24 {
		return "MEDIUMINT"
	}
	if bits <= 32 {
		return "INT"
	}
	if bits <= 64 {
		return "BIGINT"
	}
	return ""
}

func (mysql) UInt(bits uint8) DataType {
	return MySQL.Int(bits) + " UNSIGNED"
}

func (mysql) Float(precision uint8) DataType {
	if precision > 53 {
		return ""
	}
	if precision == 24 {
		return "FLOAT"
	}
	if precision == 53 {
		return "DOUBLE"
	}
	return call("FLOAT", precision)
}

func (mysql) Decimal(precision uint8, scale uint8) DataType {
	return call2("DECIMAL", precision, scale)
}

func (mysql) Text() DataType {
	return "TEXT"
}

func (mysql) FixedChar(size uint32) DataType {
	// https://dev.mysql.com/doc/refman/8.2/en/char.html
	if size > 255 {
		return ""
	}
	return call("CHAR", size)
}

func (mysql) VarChar(size uint32) DataType {
	// https://dev.mysql.com/doc/refman/8.2/en/char.html
	if size > 65_535 {
		return ""
	}
	return call("VARCHAR", size)
}

func (mysql) Enum(members []string) DataType {
	return callVar("ENUM", members)
}

func (mysql) Blob() DataType {
	return "BLOB"
}

func (mysql) Interval() DataType {
	return "INTEGER"
}

func (mysql) Date() DataType {
	return "DATE"
}

func (mysql) DateTime() DataType {
	return "DATETIME"
}

func (mysql) Time() DataType {
	return "TIME"
}

func (mysql) Bool() DataType {
	return "TINYINT(1)"
}

func (mysql) String() string {
	return "MySQL"
}

var mysqlFeatures = Features{
	"DELETE FROM":  true,
	"CREATE TABLE": true,
	"DROP TABLE":   true,

	// https://dev.mysql.com/doc/refman/8.0/en/truncate-table.html
	"TRUNCATE TABLE": true,

	// https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
	"ALTER TABLE":                               true,
	"ALTER TABLE / ADD COLUMN":                  true,
	"ALTER TABLE / ADD INDEX":                   true,
	"ALTER TABLE / ADD KEY":                     true,
	"ALTER TABLE / ADD CONSTRAINT":              true,
	"ALTER TABLE / DROP CHECK":                  true,
	"ALTER TABLE / DROP CONSTRAINT":             true,
	"ALTER TABLE / ALTER CHECK":                 true,
	"ALTER TABLE / ALTER CONSTRAINT":            true,
	"ALTER TABLE / ALTER COLUMN":                true,
	"ALTER TABLE / ALTER COLUMN / SET DEFAULT":  true,
	"ALTER TABLE / ALTER COLUMN / DROP DEFAULT": true,
	"ALTER TABLE / CHANGE COLUMN":               true,
	"ALTER TABLE / DROP COLUMN":                 true,
	"ALTER TABLE / DROP INDEX":                  true,
	"ALTER TABLE / DROP KEY":                    true,
	"ALTER TABLE / DROP PRIMARY KEY":            true,
	"ALTER TABLE / DROP FOREIGN KEY":            true,
	"ALTER TABLE / MODIFY":                      true,
	"ALTER TABLE / RENAME COLUMN":               true,
	"ALTER TABLE / RENAME INDEX":                true,
	"ALTER TABLE / RENAME KEY":                  true,
	"ALTER TABLE / RENAME TO":                   true,
}
