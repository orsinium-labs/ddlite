package dbconfig

import (
	"regexp"
	"strings"

	"github.com/Masterminds/squirrel"
)

// Placeholder is the style of placeholder for variable binding.
type Placeholder uint8

const (
	// Question is the default question placeholder (e.g. ?)
	Question Placeholder = 1

	// Dollar is a dollar-prefixed positional placeholder (e.g. $1, $2, $3).
	Dollar Placeholder = 2

	// Colon is a colon-prefixed positional placeholder (e.g. :1, :2, :3).
	Colon Placeholder = 3

	// AtP is a "@p"-prefixed positional placeholder (e.g. @p1, @p2, @p3).
	AtP Placeholder = 4
)

type NameMapper func(string) string

type Config struct {
	// Placeholder style for variable binding.
	//
	// It varies for different databases. When config is constructed using [New],
	// the default placeholder style is detected based on the passed driver name.
	Placeholder Placeholder

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
	return Config{
		Placeholder: placeholderForDriver(driver),
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

// SquirrelPlaceholder returns Placeholder recognized by squirrel.
//
// The method is used by SQL query builders.
func (c Config) SquirrelPlaceholder() squirrel.PlaceholderFormat {
	switch c.Placeholder {
	case Question:
		return squirrel.Question
	case Dollar:
		return squirrel.Dollar
	case Colon:
		return squirrel.Colon
	case AtP:
		return squirrel.AtP
	default:
		return squirrel.Question
	}
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
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql", "nrpostgres", "cockroach":
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
