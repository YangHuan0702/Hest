package service

import (
	"Hest/model/entity"
	"Hest/repository"
	"Hest/util"
	"errors"
	"gorm.io/gorm"
)

type BaseConfigServiceStr struct {
	Repo *repository.BaseConfigRepoStr
	Util *util.OperatorUtil
}

type BaseConfigService interface {
	Insert(config *entity.BaseConfig) int
	Update(config *entity.BaseConfig) int
	Del(id int) int
}

func (service *BaseConfigServiceStr) Insert(config *entity.BaseConfig) int {
	id := service.Util.GetId()
	config.Id = id
	config.CommonId = id
	config.IsCover = 0
	config.Del = 0
	service.Repo.Insert(config)
	return id
}

func (service *BaseConfigServiceStr) Update(config *entity.BaseConfig) (int, error) {
	id := service.Util.GetId()
	newConfig := entity.BaseConfig{
		Id: id, TypeUrl: config.TypeUrl, CommonId: id, IsCover: 0, Del: 0,
	}
	config.IsCover = 1
	config.Del = 1
	err := service.Repo.Db.Transaction(func(tx *gorm.DB) error {
		service.Repo.Update(config)
		service.Repo.Insert(&newConfig)
		tx.Commit()
		return nil
	})
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return id, nil
}

func (service *BaseConfigServiceStr) Del(id int) int {
	return service.Repo.Del(id)
}
