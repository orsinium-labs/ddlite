package sequel

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/Masterminds/squirrel"
)

// asSquirrel converts the given value to a squirrel expression.
//
// The given value can be a DB function, model field, or a literal value.
func asSquirrel(m any, f any) squirrel.Sqlizer {
	// function
	switch fn := f.(type) {
	case Func[int]:
		args := make([]any, 0, len(fn.Args))
		for _, arg := range fn.Args {
			args = append(args, asSquirrel(m, arg))
		}
		return squirrel.Expr(fmt.Sprintf("%s(?)", fn.Name), args...)
	default:
	}

	// model field
	fname, err := getFieldName(m, f)
	if err == nil {
		return squirrel.Expr(fname)
	}

	// literal value
	return squirrel.Expr("?", f)
}

func getModelName(model any) string {
	t := reflect.ValueOf(model).Elem().Type()
	return strings.ToLower(t.Name())
}

func getFieldName(model any, field any) (string, error) {
	target := reflect.ValueOf(field)
	if target.Kind() != reflect.Pointer {
		return "", errors.New("the field is not a pointer")
	}
	tpointer := target.Pointer()
	rmodel := reflect.ValueOf(model).Elem()
	rtype := rmodel.Type()
	for i := 0; i < rtype.NumField(); i++ {
		if rmodel.Field(i).Addr().Pointer() == tpointer {
			return rtype.Field(i).Name, nil
		}
	}
	return "", errors.New("field not found")
}

// Ref converts the given value to a pointer.
//
// Convenient for making a pointer to a literal value.
func Ref[T any](val T) *T {
	return &val
}
