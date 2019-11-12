package main

import (
	"foodio/api"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func main() {
	app := iris.New()
	app.Use(logger.New())

	app.Get("/status", api.GetStatus)
	app.Get("/recipes", api.GetRecipes)
	app.Get("/diets", api.GetDiets)

	app.Run(iris.Addr(getPort()))
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
