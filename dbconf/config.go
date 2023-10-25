package dbconf

import (
	"regexp"
	"strings"
)

type Dialect rune

const (
	CockroachDB Dialect = '🪳'
	MySQL       Dialect = '🐬'
	OracleDB    Dialect = '👁'
	PostgreSQL  Dialect = '🐘'
	SQLite      Dialect = '🪶'
	SQLServer   Dialect = '🪟'
)

type NameMapper func(string) string

type Config struct {
	// Placeholder style for variable binding.
	//
	// It varies for different databases. The default placeholder style
	// is detected based on the passed driver name.
	Placeholder Placeholder

	// Dialect indicates which syntax should be used for the database.
	//
	// The default dialect is detected based on the passed driver name.
	Dialect Dialect

	// DriverName is the name of the package used as the DB driver.
	DriverName string

	// ToTable converts Go struct name into DB table name.
	//
	// Default: [CamelToSnake].
	ToTable NameMapper

	// ToColumn converts Go struct field name into DB column name.
	//
	// Default: [CamelToSnake].
	ToColumn NameMapper

	// ToField converts DB column name into Go struct field name.
	//
	// Default: [SnakeToCamel].
	ToField NameMapper

	// Models is a list of struct instances used to specify columns in the query.
	//
	// Usually, you don't need to specify it explicitly, the query builders
	// automatically add all needed models.
	Models []any
}

// New creates a new config instance for the given DB driver name with good defaults.
func New(driver string) Config {
	dialect := dialectForDriver(driver)
	return Config{
		Placeholder: placeholderForDriver(driver),
		Dialect:     dialect,
		DriverName:  driver,
		ToTable:     CamelToSnake,
		ToColumn:    CamelToSnake,
		ToField:     SnakeToCamel,
	}
}

// WithModel returns a copy of the config with the given model added into the registry of models.
//
// The model must be a pointer.
func (c Config) WithModel(m any) Config {
	c.Models = append([]any{m}, c.Models...)
	return c
}

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// CamelToSnake transforms CamelCase into snake_case.
//
// Abbreviations are recognized, so "IPAddr" will be transformed into "ip_addr".
func CamelToSnake(s string) string {
	s = matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	s = matchAllCap.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(s)
}

// SnakeToCamel transforms snake_case into CamelCase.
//
// Can also transform kebab-case into CamelCase.
//
// Abbreviatins aren't recognized, so "ip_addr" will be transformed
// into "IpAddr", not "IPAddr".
func SnakeToCamel(s string) string {
	camelCase := ""
	isToUpper := false
	for k, v := range s {
		if k == 0 {
			camelCase = strings.ToUpper(string(s[0]))
		} else {
			if isToUpper {
				camelCase += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' || v == '-' {
					isToUpper = true
				} else {
					camelCase += string(v)
				}
			}
		}
	}
	return camelCase
}

// placeholderForDriver infers the correct placeholder for the given driver name.
func placeholderForDriver(driver string) Placeholder {
	switch strings.ToLower(driver) {
	case "postgres", "pgx", "pq", "pq-timeouts", "cloudsqlpostgres", "ql", "nrpostgres", "cockroach":
		return Dollar
	case "mysql", "sqlite", "sqlite3", "nrmysql", "nrsqlite3":
		return Question
	case "oci8", "ora", "goracle", "oracle", "godror":
		return Colon
	case "sqlserver":
		return AtP
	default:
		return Question
	}
}

// dialectForDriver infers the correct dialect for the given driver name.
func dialectForDriver(driver string) Dialect {
	switch strings.ToLower(driver) {
	case "cockroach":
		return CockroachDB
	case "mysql", "nrmysql":
		return MySQL
	case "oci8", "ora", "goracle", "oracle", "godror":
		return OracleDB
	case "postgres", "pgx", "pq", "pq-timeouts", "cloudsqlpostgres", "ql", "nrpostgres":
		return PostgreSQL
	case "sqlite", "sqlite3", "nrsqlite3":
		return SQLite
	case "sqlserver":
		return SQLServer
	default:
		return SQLite
	}
}
