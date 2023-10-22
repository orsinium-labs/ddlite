package qb

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/orsinium-labs/sequel/dbconf"
)

type tInsert[T Model] struct {
	model  *T
	fields []any
	items  []T
}

func Insert[T Model](model *T, fields ...any) tInsert[T] {
	return tInsert[T]{
		model:  model,
		fields: fields,
		items:  make([]T, 0, 1), // most often, exactly one item will be inserted
	}
}

func (i tInsert[T]) Values(items ...T) tInsert[T] {
	i.items = append(i.items, items...)
	return i
}

func (i tInsert[T]) Squirrel(conf dbconf.Config) (squirrel.Sqlizer, error) {
	conf = conf.WithModel(i.model)
	// get column names
	fnames := make([]string, 0, len(i.fields))
	cnames := make([]string, 0, len(i.fields))
	for _, f := range i.fields {
		fname, err := getFieldName(conf, f)
		if err != nil {
			return squirrel.InsertBuilder{}, fmt.Errorf("get column name: %v", err)
		}
		fnames = append(fnames, fname)
		cnames = append(cnames, conf.ToColumn(fname))
	}

	// make builder, set column names and table name
	q := squirrel.Insert(getModelName(i.model))
	q = q.PlaceholderFormat(conf.SquirrelPlaceholder())
	q = q.Columns(cnames...)

	// set values to insert
	for _, item := range i.items {
		values := make([]any, 0, len(fnames))
		for _, fname := range fnames {
			value, err := getField(&item, fname)
			if err != nil {
				return q, fmt.Errorf("get field value: %v", err)
			}
			values = append(values, value)
		}
		q = q.Values(values...)
	}

	return q, nil
}
