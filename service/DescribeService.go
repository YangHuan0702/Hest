package service

import (
	"Hest/model/dto"
	"Hest/repository"
)

type DescribeServiceStr struct {
	dao *repository.BaseDescribeRepoStr
}

type BaseDescribeSyncDTO dto.BaseDescribeSyncDTO

type DescribeService interface {
	// Save 存储
	Save(dto *BaseDescribeSyncDTO) bool

	// Sync 执行前会先清理一遍数据
	Sync(dto *BaseDescribeSyncDTO) bool
}

func (ser *DescribeServiceStr) Save(dto *BaseDescribeSyncDTO) {

}

func (ser *DescribeServiceStr) Sync(dto *BaseDescribeSyncDTO) {

}
