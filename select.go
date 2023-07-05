package sequel

import (
	"fmt"
	"reflect"
	"strings"
)

type selectQ[T any] struct {
	fields []any
	model  any
}

func Select[T any](model T, fields ...any) selectQ[T] {
	return selectQ[T]{model: model, fields: fields}
}

func (s selectQ[T]) String() string {
	fnames := make([]string, 0, len(s.fields))
	for _, f := range s.fields {
		fnames = append(fnames, getFieldName(s.model, f))
	}
	joined := strings.Join(fnames, ", ")
	table := getModelName(s.model)
	return fmt.Sprintf("SELECT %s FROM %s", joined, table)
}

func getModelName(model any) string {
	t := reflect.ValueOf(model).Elem().Type()
	return strings.ToLower(t.Name())
}

func getFieldName(model any, field any) string {
	target := reflect.ValueOf(field).Pointer()
	rmodel := reflect.ValueOf(model).Elem()
	rtype := rmodel.Type()
	for i := 0; i < rtype.NumField(); i++ {
		if rmodel.Field(i).Addr().Pointer() == target {
			return rtype.Field(i).Name
		}
	}
	panic("field not found")
}
