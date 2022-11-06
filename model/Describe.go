package model

type BaseDescribe struct {
	name     string `json:"name" gorm:"column:name"`
	url      string `json:"url" gorm:"column:url"`
	sver     string `json:"sver" gorm:"column:sver"`
	location string `json:"location" gorm:"column:location"`
}
type LineDescribe struct {
	consumes    string `json:"consumes" gorm:"column:consumes"`
	operationId string `json:"operation_id" gorm:"column:operation_id"`
	paramsId    string `json:"params_id" gorm:"column:params_id"`
	produces    string `json:"produces" gorm:"column:produces"`
	summary     string `json:"summary" gorm:"column:summary"`
	tags        string `json:"tags" gorm:"column:tags"`
}

type Param struct {
	description string `json:"description" gorm:"column:description"`
	format      string `json:"format" gorm:"column:format"`
	in          string `json:"in" gorm:"column:in"`
	name        string `json:"name" gorm:"column:name"`
	required    string `json:"required" gorm:"column:required"`
	types       string `json:"types" gorm:"column:type"`
}

type LineOfParam struct {
}
