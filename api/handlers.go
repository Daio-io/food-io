package api

import (
	"github.com/kataras/iris"
	"foodio/db"
	"time"
	"math/rand"
	"strings"
)

func GetRecipes(ctx *iris.Context) {
	limit, err := ctx.URLParamInt("limit")
	with := ctx.URLParam("with")
	diets := strings.Split(ctx.URLParam("diets"), ",")
	if err != nil || limit > 500 {
		limit = 100
	}

	session := db.NewSession()
	defer session.Close()

	query := db.NewQuery()
	query.Amount = limit
	query.MultiFilter("diets.title", diets)

	model := []Result{}
	col := session.Collection("recipes", model)

	var results interface{}
	var queryErr error

	if with != "" {
		results, queryErr = col.Text(with, query)
	} else {
		results, queryErr = col.Find(query)
	}

	if queryErr != nil || results == nil {
		ctx.JSON(iris.StatusBadRequest, Status{"Request Failed"})
	} else {
		data := results.([]Result)
		ctx.JSON(iris.StatusOK, shuffle(data))
	}

}

func GetDiets(ctx *iris.Context) {
	model := []string{
		"gluten-free",
		"vegan",
		"vegetarian",
		"nut-free",
		"egg-free",
		"dairy-free",
		"pregnancy-friendly",
		"shellfish-free"}

	result := map[string][]string{"diets": model}
	ctx.JSON(iris.StatusOK, result)
}

func GetStatus(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, Status{"OK"})
}

func shuffle(data []Result) []Result {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))

	for i := len(data) - 1; i > 0; i-- {
		j := rand.Intn(i)
		data[i], data[j] = data[j], data[i]
	}

	return data
}