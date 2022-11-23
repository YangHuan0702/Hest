package init

import (
	"Hest/controller"
	"Hest/repository"
	"Hest/service"
	"Hest/util"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gorm.io/gorm"
)

// utils
var operatorUtil *util.OperatorUtil

// services
var describeService *service.DescribeServiceStr

func Controllers(application *iris.Application) {
	party := application.Party("/describe")
	mvc.Configure(party, func(mvc *mvc.Application) {
		mvc.Handle(controller.DescribeControllerStr{
			Service: describeService,
			Util:    operatorUtil,
		})
	})
}

func MvcInit(DbCli *gorm.DB, RedisClient *redis.Conn) {
	operatorUtil = &util.OperatorUtil{RedisClient: RedisClient}

	describeService = &service.DescribeServiceStr{
		Dao: &repository.BaseDescribeRepoStr{Dao: DbCli},
	}
}
