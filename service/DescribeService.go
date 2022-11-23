package service

import (
	"Hest/model/dto"
	"Hest/repository"
)

type DescribeServiceStr struct {
	Dao *repository.BaseDescribeRepoStr
}

type BaseDescribeSyncDTO dto.BaseDescribeSyncDTO

type DescribeService interface {
	// Save 存储
	Save(dto *BaseDescribeSyncDTO) bool

	// Sync 执行前会先清理一遍数据
	Sync(dto *BaseDescribeSyncDTO) bool
}

func (ser *DescribeServiceStr) Save(dto *BaseDescribeSyncDTO) {
	ser.Dao.Dao.Begin()
	ser.Dao.BatchInsertBase(dto.Bases)
	ser.Dao.BatchInsertLine(dto.Lines)
	ser.Dao.BatchInsertParam(dto.Params)
	ser.Dao.BatchInsertLineOfParam(dto.Replace)
	ser.Dao.Dao.Commit()
}

func (ser *DescribeServiceStr) Sync(dto *BaseDescribeSyncDTO) {
	ser.Dao.Dao.Begin()
	ser.Dao.ClearBase()
	ser.Dao.ClearLine()
	ser.Dao.ClearLineOfParam()
	ser.Dao.ClearParam()
	ser.Dao.Dao.Commit()
	ser.Save(dto)
}
