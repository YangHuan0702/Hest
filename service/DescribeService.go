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
	ser.dao.Dao.Begin()
	ser.dao.BatchInsertBase(dto.Bases)
	ser.dao.BatchInsertLine(dto.Lines)
	ser.dao.BatchInsertParam(dto.Params)
	ser.dao.BatchInsertLineOfParam(dto.Replace)
	ser.dao.Dao.Commit()
}

func (ser *DescribeServiceStr) Sync(dto *BaseDescribeSyncDTO) {
	ser.dao.Dao.Begin()
	ser.dao.ClearBase()
	ser.dao.ClearLine()
	ser.dao.ClearLineOfParam()
	ser.dao.ClearParam()
	ser.dao.Dao.Commit()
	ser.Save(dto)
}
