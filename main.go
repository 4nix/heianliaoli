package main

import (
	"github.com/kataras/iris"
	"./src/userinterface"
)

func main () {
	app := iris.Default()
	app.Get("/list", userinterface.List)

	app.Run(iris.Addr(":8082"))
}