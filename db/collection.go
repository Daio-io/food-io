package db

import (
	"gopkg.in/mgo.v2"
	"reflect"
	"gopkg.in/mgo.v2/bson"
)

// Collection - Struct
type Collection struct {
	col   *mgo.Collection
	model interface{}
}

// Find - Return queried results list
func (c *Collection) Find(options QueryOptions) (interface{}, error) {
	modelType := reflect.New(reflect.TypeOf(c.model))

	query := bson.M{}

	if len(options.GetFilters()) > 0 {
		query["$and"] = options.GetFilters()
	}

	err := c.col.Find(query).Limit(options.Amount).All(modelType.Interface())
	if err != nil {
		return nil, err
	}
	return modelType.Elem().Interface(), nil
}

// Find - Return queried results list
func (c *Collection) Text(query string, options QueryOptions) (interface{}, error) {
	modelType := reflect.New(reflect.TypeOf(c.model))
	textSearch := bson.M{"$text": bson.M{"$search": query,  "$caseSensitive": false}}

	if len(options.GetFilters()) > 0 {
		textSearch["$and"] = options.GetFilters()
	}

	err := c.col.Find(textSearch).Limit(options.Amount).All(modelType.Interface())
	if err != nil {
		return nil, err
	}
	return modelType.Elem().Interface(), nil
}

