package handlers

import (
	"github.com/iris-app/repository"
	"github.com/kataras/iris/v12"
)

func GetProductsHandle(ctx iris.Context) {

	productList := repository.GetProducts()

	ctx.JSON(productList)

	iris.New().Logger().Info("get products request")

}
