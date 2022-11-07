package repository

import "Hest/model"

type BaseDescribeRepo interface {
	Insert(base *model.BaseDescribe) int
	Update(base *model.BaseDescribe) int
	Del(base *model.BaseDescribe) int
	BatchInsert(bases []*model.BaseDescribe) int
	BatchDelete(ids []uint) int
}

type LineDescribeRepo interface {
	Insert(base *model.LineDescribe) int
	Update(base *model.LineDescribe) int
	Del(base *model.LineDescribe) int
	BatchInsert(bases []*model.LineDescribe) int
	BatchDelete(ids []uint) int
}

type ParamRepo interface {
	Insert(base *model.Param) int
	Update(base *model.Param) int
	Del(base *model.Param) int
	BatchInsert(bases []*model.Param) int
	BatchDelete(ids []uint) int
}

type LineOfParamRepo interface {
	Insert(base *model.LineOfParam) int
	Update(base *model.LineOfParam) int
	Del(base *model.LineOfParam) int
	BatchInsert(bases []*model.LineOfParam) int
	BatchDelete(ids []uint) int
}
