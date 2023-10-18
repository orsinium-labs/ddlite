package dbconfig

import (
	"regexp"
	"strings"
)

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
	Models []any

	// Plaholder style for variable binding.
	Placeholder Placeholder

	// DriverName is the name of the package used as the DB driver.
	DriverName string

	// ToTable converts Go struct name into DB table name.
	ToTable NameMapper

	// ToColumn converts Go struct field name into DB column name.
	ToColumn NameMapper

	// ToField converts DB column name into Go struct field name.
	ToField NameMapper
}

func New(driver string) Config {
	return Config{
		Placeholder: placeholderForDriver(driver),
		DriverName:  driver,
		ToTable:     CamelToSnake,
		ToColumn:    CamelToSnake,
		ToField:     SnakeToCamel,
	}
}

func (c Config) WithModel(m any) Config {
	c.Models = append([]any{m}, c.Models...)
	return c
}

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func CamelToSnake(s string) string {
	s = matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	s = matchAllCap.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(s)
}

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
	switch driver {
	case "postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql", "nrpostgres", "cockroach":
		return Dollar
	case "mysql", "sqlite", "sqlite3", "nrmysql", "nrsqlite3":
		return Question
	case "oci8", "ora", "goracle", "godror":
		return Colon
	case "sqlserver":
		return AtP
	default:
		return Question
	}
}
