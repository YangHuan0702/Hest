package entity

type BaseConfig struct {
	Id       int    `json:"id" gorm:"column:id"`
	TypeUrl  string `json:"typeUrl" gorm:"column:type_url'"`
	CommonId int    `json:"commonId" gorm:"column:common_id"`
	IsCover  int    `json:"is_cover" gorm:"column:is_cover"`
	Del      int    `json:"del" gorm:"column:del"`
}
