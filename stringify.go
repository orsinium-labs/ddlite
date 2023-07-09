package sequel

import (
	"errors"
	"reflect"
	"strings"
)

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
