package main

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/logger"
	"foodio/db"
	"foodio/api"
	"os"
)

func main() {


	app := iris.New()
	app.Use(logger.New())

	app.Get("/status", func(ctx *iris.Context){
		ctx.JSON(iris.StatusOK, api.Status{"OK"})
	})

	app.Get("/recipes", func(ctx *iris.Context) {
		limit, err:= ctx.URLParamInt("limit")
		if err != nil || limit > 100 {
			limit = 100
		}
		
		session := db.NewSession()
		defer session.Close()

		query := db.NewQuery()
		query.Amount = limit
		
		model := []api.Result{}
		col := session.Collection("recipes", model)
		results, _ := col.Find(query)

		data := results.([]api.Result)

		ctx.JSON(iris.StatusOK, data)
	})

	app.Listen(getPort())

}

// GetPort - Get the port number set
// by the environment PORT
// or return default 9000
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "localhost:9000"
	}
	return ":" + port
}
