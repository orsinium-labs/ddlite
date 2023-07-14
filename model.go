package qb

import (
	"errors"
	"reflect"
	"strings"
)

type Model any

func getModelName(model Model) string {
	t := reflect.ValueOf(model).Elem().Type()
	return strings.ToLower(t.Name())
}

func getField(model any, field string) (any, error) {
	vmodel := reflect.ValueOf(model)
	if vmodel.Kind() != reflect.Pointer {
		return "", errors.New("the model is not a pointer")
	}
	rmodel := vmodel.Elem()
	f := rmodel.FieldByName(field)
	return f.Addr().Interface(), nil
}

func getFieldName(model any, field any) (string, error) {
	target := reflect.ValueOf(field)
	if target.Kind() != reflect.Pointer {
		return "", errors.New("the field is not a pointer")
	}
	tpointer := target.Pointer()
	vmodel := reflect.ValueOf(model)
	if vmodel.Kind() != reflect.Pointer {
		return "", errors.New("the model is not a pointer")
	}
	rmodel := vmodel.Elem()
	rtype := rmodel.Type()
	for i := 0; i < rtype.NumField(); i++ {
		if rmodel.Field(i).Addr().Pointer() == tpointer {
			return rtype.Field(i).Name, nil
		}
	}
	return "", errors.New("field not found")
}
