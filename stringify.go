package sequel

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/orsinium-labs/sequel/dbfuncs"
)

func stringify(m any, f any) string {
	switch fn := f.(type) {
	case dbfuncs.Func[int]:
		return fmtFunc(m, fn.Name, fn.Args)
	default:
	}

	fname, err := getFieldName(m, f)
	if err == nil {
		return fname
	}
	fname, _ = fmtLiteral(f)
	return fname
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

func fmtLiteral(val any) (string, error) {
	switch val.(type) {
	case int, *int:
		return fmt.Sprintf("%d", val), nil
	default:
		return "?", nil
	}
}

func fmtFunc(m any, name string, args []any) string {
	formatted := make([]string, 0, len(args))
	for _, arg := range args {
		farg := stringify(m, arg)
		formatted = append(formatted, farg)
	}
	joined := strings.Join(formatted, ", ")
	return fmt.Sprintf("%s(%s)", name, joined)
}
