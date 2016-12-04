package main

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/logger"
	"os"
)

func main() {


	app := iris.New()
	app.Use(logger.New())

	app.Get("/status", func(ctx *iris.Context){
		ctx.JSON(iris.StatusOK, Status{"OK"})
	})

	app.Listen(getPort())

}

type Status struct {
	Status string `json:"status"`
}

// GetPort - Get the port number set
// by the environment PORT
// or return default 9000
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "localhost:9000"
	}
	return ":" + port
}
