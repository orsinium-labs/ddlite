package internal

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/orsinium-labs/sequel/dbconf"
)

type Model any

// GetTableName converts struct name into table name.
func GetTableName(conf dbconf.Config, model Model) string {
	t := reflect.ValueOf(model).Elem().Type()
	return conf.ToTable(t.Name())
}

// GetField extracts the value from the given struct in the given struct field.
//
// `model` is the struct and `field` is the struct field name
// from which extract the vlaue.
func GetField(model Model, field string) (any, error) {
	vmodel := reflect.ValueOf(model)
	if vmodel.Kind() != reflect.Pointer {
		return "", errors.New("the model is not a pointer")
	}
	rmodel := vmodel.Elem()
	f := rmodel.FieldByName(field)
	zero := reflect.Value{}
	if f == zero {
		return "", fmt.Errorf("the model doesn't have field `%s`", field)
	}
	return f.Addr().Interface(), nil
}

// GetColumnName returns column name for the given struct field.
func GetColumnName(conf dbconf.Config, field any) (string, error) {
	name, err := GetFieldName(conf, field)
	name = conf.ToColumn(name)
	return name, err
}

// GetFieldName returns the name of the given struct field.
func GetFieldName(conf dbconf.Config, field any) (string, error) {
	target := reflect.ValueOf(field)
	if target.Kind() != reflect.Pointer {
		return "", errors.New("the field is not a pointer")
	}
	if len(conf.Models) == 0 {
		return "", errors.New("no models registered in the config")
	}
	tpointer := target.Pointer()
	for _, model := range conf.Models {
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
	}
	return "", errors.New("field not found")
}
