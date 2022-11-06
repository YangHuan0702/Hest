package model

type BaseDescribe struct {
	name     string `json:"name" gorm:"column:name"`
	url      string `json:"url" gorm:"column:url"`
	sver     string `json:"sver" gorm:"column:sver"`
	location string `json:"location" gorm:"column:location"`
}


type 