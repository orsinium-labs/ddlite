package dml

import (
	"fmt"

	"github.com/orsinium-labs/sequel/dbconf"
	"github.com/orsinium-labs/sequel/internal"
	"github.com/orsinium-labs/sequel/internal/tokens"
)

type tInsert[T internal.Model] struct {
	model  *T
	fields []any
	items  []T
}

func Insert[T internal.Model](model *T, fields ...any) tInsert[T] {
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

func (i tInsert[T]) Tokens(conf dbconf.Config) (tokens.Tokens, error) {
	conf = conf.WithModel(i.model)
	ts := tokens.New(
		tokens.Keyword("INSERT INTO"),
		tokens.TableName(internal.GetTableName(conf, i.model)),
		tokens.LParen(),
	)
	// get column names
	fieldNames := make([]string, 0, len(i.fields))
	for i, field := range i.fields {
		if i > 0 {
			ts.Add(tokens.Comma())
		}
		fieldName, err := internal.GetFieldName(conf, field)
		if err != nil {
			return tokens.New(), fmt.Errorf("get column name: %v", err)
		}
		fieldNames = append(fieldNames, fieldName)
		ts.Add(tokens.ColumnName(conf.ToColumn(fieldName)))
	}
	ts.Add(
		tokens.RParen(),
		tokens.Keyword("VALUES"),
	)

	// set values to insert
	for i, item := range i.items {
		if i > 0 {
			ts.Add(tokens.Comma())
		}
		ts.Add(tokens.LParen())
		for i, fname := range fieldNames {
			if i > 0 {
				ts.Add(tokens.Comma())
			}
			value, err := internal.GetField(&item, fname)
			if err != nil {
				return tokens.New(), fmt.Errorf("get field value: %v", err)
			}
			ts.Add(tokens.Bind(value))
		}
		ts.Add(tokens.RParen())
	}

	return ts, nil
}
