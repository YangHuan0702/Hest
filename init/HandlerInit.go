package init

import (
	"Hest/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Handlers(application *iris.Application) {
	party := application.Party("/describe")
	mvc.Configure(party, func(mvc *mvc.Application) {
		mvc.Handle(new(controller.DescribeControllerStr))
	})
}
