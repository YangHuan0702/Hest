package main

import (
	hinit "Hest/init"
	_ "github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/v12"
)

func main() {
	application := iris.New()

	// 初始化Controller
	hinit.Handlers(application)

	application.Run(iris.Addr(":8080"))
	iris.WithoutServerError(iris.ErrServerClosed)

}
