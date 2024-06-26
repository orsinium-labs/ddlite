package ddl

type DataType string

func Integer() DataType {
	return "INTEGER"
}

func Real() DataType {
	return "REAL"
}

func Numeric() DataType {
	return "NUMERIC"
}

func Text() DataType {
	return "TEXT"
}

func Blob() DataType {
	return "BLOB"
}

func Time() DataType {
	return "TIME"
}
