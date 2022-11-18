package main

import (
	_ "Hest/init"
	_ "github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/v12"
)

func main() {
	application := iris.New()

	application.Handle("GET", "/", func(context iris.Context) {
		context.HTML("<h1>Hello Hest</h1>")
	})

	application.Run(iris.Addr(":8080"))
}
