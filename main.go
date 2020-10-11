package main

import (
	"github.com/iris-app/handlers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func main() {
	app := iris.New()

	app.Use(logger.New())

	app.Handle("GET", "/products", handlers.GetProductsHandle)

	app.Run(iris.Addr(":8080")) 
}
