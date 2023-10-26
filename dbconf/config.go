package dbconf

import (
	"regexp"
	"strings"

	"github.com/orsinium-labs/sequel/dialects"
)

type Dialect rune

type NameMapper func(string) string

type Config struct {
	// Dialect indicates which syntax should be used for the database.
	//
	// The default dialect is detected based on the passed driver name.
	Dialect dialects.Dialect

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
func New(dialect dialects.Dialect) Config {
	if dialect == nil {
		panic("dialect must not be nil")
	}
	return Config{
		Dialect:  dialect,
		ToTable:  CamelToSnake,
		ToColumn: CamelToSnake,
		ToField:  SnakeToCamel,
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
