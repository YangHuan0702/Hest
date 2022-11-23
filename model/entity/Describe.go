package entity

type BaseDescribe struct {
	Id       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Url      string `json:"url" gorm:"column:url"`
	Sver     string `json:"swaggerVersion" gorm:"column:sver"`
	Location string `json:"location" gorm:"column:location"`
}
type LineDescribe struct {
	Id          uint   `json:"id" gorm:"column:id"`
	Consumes    string `json:"consumes" gorm:"column:consumes"`
	OperationId string `json:"operation_id" gorm:"column:operation_id"`
	ParamsId    string `json:"params_id" gorm:"column:params_id"`
	Produces    string `json:"produces" gorm:"column:produces"`
	Summary     string `json:"summary" gorm:"column:summary"`
	Tags        string `json:"tags" gorm:"column:tags"`
}

type Param struct {
	Id          uint   `json:"id" gorm:"column:id"`
	Description string `json:"description" gorm:"column:description"`
	Format      string `json:"format" gorm:"column:format"`
	In          string `json:"in" gorm:"column:in"`
	Name        string `json:"name" gorm:"column:name"`
	Required    string `json:"required" gorm:"column:required"`
	Types       string `json:"types" gorm:"column:type"`
}

type LineOfParam struct {
	Id      uint `json:"id" gorm:"column:id"`
	LineId  uint `json:"line_id" gorm:"column:line_id"`
	ParamId uint `json:"param_id" gorm:"column:param_id"`
}
