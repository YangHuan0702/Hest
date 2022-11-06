package model

type BaseConfig struct {
	Id      string `json:"id" gorm:"column:id"`
	TypeUrl string `json:"typeUrl" gorm:"column:type_url'"`
}
