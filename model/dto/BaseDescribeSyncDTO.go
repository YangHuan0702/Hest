package dto

import "Hest/model/entity"

type BaseDescribeSyncDTO struct {
	Bases   []*entity.BaseDescribe
	Lines   []*entity.LineDescribe
	Params  []*entity.Param
	Replace []*entity.LineOfParam
}

func parseByJsonString(json string) *BaseDescribeSyncDTO {
	return nil
}
