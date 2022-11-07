package dto

import "Hest/model/entity"

type BaseDescribeSyncDTO struct {
	bases   []*entity.BaseDescribe
	lines   []*entity.LineDescribe
	params  []*entity.Param
	replace []*entity.LineOfParam
}

func parseByJsonString(json string) *BaseDescribeSyncDTO {

	return nil
}
