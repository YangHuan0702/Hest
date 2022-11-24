package repository

import (
	"Hest/model/entity"
	"gorm.io/gorm"
)

type BaseConfigRepoStr struct {
	Db *gorm.DB
}

type BaseConfigRepo interface {
	Insert(config *entity.BaseConfig) int
	Update(config *entity.BaseConfig) int
	Del(id int) int
}

func (repo *BaseConfigRepoStr) Insert(config *entity.BaseConfig) int {
	repo.Db.Debug().Create(config)
	return 1
}

func (repo *BaseConfigRepoStr) Update(config *entity.BaseConfig) int {
	repo.Db.Debug().Save(config)
	return 1
}

func (repo *BaseConfigRepoStr) Del(id int) int {
	repo.Db.Debug().Delete(entity.BaseConfig{}, id)
	return 1
}
