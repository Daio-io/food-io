package api

type Status struct {
	Status string `json:"status"`
}

type Result struct {
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