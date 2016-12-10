package api

import "gopkg.in/mgo.v2/bson"

type Status struct {
	Status string `json:"status"`
}

type Result struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	CookingTime string `json:"cookingTime" bson:"cookingTime"`
	Servings    string `json:"servings" bson:"servingsText"`
	Thumbnail   string `json:"thumbnail" bson:"thumbnail"`
	Methods     []Method `json:"methods" bson:"methods"`
	Stages      []Stage `json:"stages" bson:"stages"`
}

type Method struct {
	Text string `json:"step" bson:"text"`
}

type Stage struct {
	Title       string `json:"title" bson:"title"`
	Ingredients []Ingredient `json:"ingredients" bson:"ingredients"`
}

type Ingredient struct {
	Text  string `json:"text" bson:"text"`
	Foods []Food `json:"foods" bson:"foods"`
}

type Food struct {
	Title string `json:"title" bson:"title"`
}