package api

import (
	"github.com/kataras/iris"
	"foodio/db"
)

func GetRecipes(ctx *iris.Context) {
	limit, err := ctx.URLParamInt("limit")
	with := ctx.URLParam("with")
	if err != nil || limit > 100 {
		limit = 100
	}

	session := db.NewSession()
	defer session.Close()

	query := db.NewQuery()
	query.Amount = limit

	model := []Result{}
	col := session.Collection("recipes", model)
	results, _ := col.Text(with, query)

	data := results.([]Result)

	ctx.JSON(iris.StatusOK, data)
}

func GetStatus(ctx *iris.Context){
	ctx.JSON(iris.StatusOK, Status{"OK"})
}