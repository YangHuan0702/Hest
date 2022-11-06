package model

type BaseDescribe struct {
	Name     string `json:"name" gorm:"column:name"`
	Url      string `json:"url" gorm:"column:url"`
	Sver     string `json:"sver" gorm:"column:sver"`
	Location string `json:"location" gorm:"column:location"`
}
type LineDescribe struct {
	Consumes    string `json:"consumes" gorm:"column:consumes"`
	OperationId string `json:"operation_id" gorm:"column:operation_id"`
	ParamsId    string `json:"params_id" gorm:"column:params_id"`
	Produces    string `json:"produces" gorm:"column:produces"`
	Summary     string `json:"summary" gorm:"column:summary"`
	Tags        string `json:"tags" gorm:"column:tags"`
}

type Param struct {
	Description string `json:"description" gorm:"column:description"`
	Format      string `json:"format" gorm:"column:format"`
	In          string `json:"in" gorm:"column:in"`
	Name        string `json:"name" gorm:"column:name"`
	Required    string `json:"required" gorm:"column:required"`
	Types       string `json:"types" gorm:"column:type"`
}

type LineOfParam struct {
}
