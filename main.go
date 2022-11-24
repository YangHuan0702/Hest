package main

import (
	hinit "Hest/init"
	_ "github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/v12"
)

func main() {
	application := iris.New()

	application.RegisterView(iris.HTML("./views", ".html"))
	iris.WithoutServerError(iris.ErrServerClosed)

	// 初始化Controller
	hinit.Controllers(application)

	application.Run(iris.Addr(":8080"))

}
