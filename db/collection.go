package db

import (
	"gopkg.in/mgo.v2"
	"reflect"
)

// Collection - Struct
type Collection struct {
	col   *mgo.Collection
	model interface{}
}

// Find - Return queried results list
func (c *Collection) Find(options QueryOptions) (interface{}, error) {
	modelType := reflect.New(reflect.TypeOf(c.model))
	err := c.col.Find(options.GetFilters()).Limit(options.Amount).All(modelType.Interface())
	if err != nil {
		return nil, err
	}
	return modelType.Elem().Interface(), nil
}

