package controller

import (
	"Hest/model/web"
	"Hest/rpc"
	"Hest/service"
	"Hest/util"
	"github.com/kataras/iris/v12"
)

type DescribeControllerStr struct {
	Service *service.DescribeServiceStr
	Util    *util.OperatorUtil
}

// GetBaseSync 同步 /base/sync
func (contr *DescribeControllerStr) GetBaseSync(context iris.Context) {
	baseDescribe, err := rpc.ReadBaseDescribe()
	if err != nil {
		_ = context.JSON(web.ERROR(err.Error()))
	}
	length := len(*baseDescribe)
	for i := 0; i < length; i++ {
		base := (*baseDescribe)[i]
		base.Id = contr.Util.GetId()
	}
	//service.
	_ = context.JSON(web.SUCCESS())
}
