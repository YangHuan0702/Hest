package repository

import (
	"Hest/model/entity"
	"bytes"
	"fmt"
	"gorm.io/gorm"
)

type BaseDescribeRepoStr struct {
	Dao *gorm.DB
}

type BaseDescribeRepo interface {
	InsertBase(base *entity.BaseDescribe) int
	UpdateBase(base *entity.BaseDescribe) int
	DelBase(base *entity.BaseDescribe) int
	BatchInsertBase(bases []*entity.BaseDescribe) int
	BatchDeleteBase(ids []uint) int
	ClearBase()
}

func (repo *BaseDescribeRepoStr) InsertBase(base *entity.BaseDescribe) int {
	repo.Dao.Debug().Create(base)
	return 1
}

func (repo *BaseDescribeRepoStr) UpdateBase(base *entity.BaseDescribe) int {
	repo.Dao.Debug().Save(base)
	return 1
}

func (repo *BaseDescribeRepoStr) DelBase(base *entity.BaseDescribe) int {
	repo.Dao.Debug().Delete(base).Unscoped()
	return 1
}

func (repo *BaseDescribeRepoStr) BatchInsertBase(bases []*entity.BaseDescribe) int {
	sql := "insert into base_describe (id,name,url,sver,location) values "
	var buffer bytes.Buffer
	if _, err := buffer.WriteString(sql); err != nil {
		panic(err)
	}
	for i, e := range bases {
		if i == len(bases)-1 {
			buffer.WriteString(fmt.Sprintf("(%d,%s,%s,%s,%s),", e.Id, e.Name, e.Url, e.Sver, e.Location))
		} else {
			buffer.WriteString(fmt.Sprintf("(%d,%s,%s,%s,%s)", e.Id, e.Name, e.Url, e.Sver, e.Location))
		}
	}
	repo.Dao.Debug().Exec(buffer.String())
	return len(bases)
}

func (repo *BaseDescribeRepoStr) BatchDeleteBase(ids []uint) int {
	repo.Dao.Debug().Delete(entity.BaseDescribe{}, ids)
	return len(ids)
}

func (repo *BaseDescribeRepoStr) ClearBase() {
	repo.Dao.Debug().Where("1 = 1").Delete(entity.BaseDescribe{})
}

type LineDescribeRepo interface {
	InsertLine(line *entity.LineDescribe) int
	UpdateLine(line *entity.LineDescribe) int
	DelLine(line *entity.LineDescribe) int
	BatchInsertLine(bases []*entity.LineDescribe) int
	BatchDeleteLine(ids []uint) int
	ClearLine()
}

func (repo *BaseDescribeRepoStr) InsertLine(line *entity.LineDescribe) int {
	repo.Dao.Debug().Create(line)
	return 1
}

func (repo *BaseDescribeRepoStr) UpdateLine(line *entity.LineDescribe) int {
	repo.Dao.Debug().Updates(line)
	return 1
}

func (repo *BaseDescribeRepoStr) DelLine(line *entity.LineDescribe) int {
	repo.Dao.Debug().Delete(line)
	return 1
}

func (repo *BaseDescribeRepoStr) BatchInsertLine(lines []*entity.LineDescribe) int {
	sql := "INSERT INTO line_describe(id,consumes,operation_id,params_id,produces,summary,tags) values"
	var buffer bytes.Buffer

	if _, err := buffer.WriteString(sql); err != nil {
		panic(nil)
	}
	for i, l := range lines {
		if i == len(lines)-1 {
			buffer.WriteString(fmt.Sprintf("(%d,%s,%s,%s,%s,%s,%s),", l.Id, l.Consumes, l.OperationId, l.ParamsId, l.Produces, l.Summary, l.Tags))
		} else {
			buffer.WriteString(fmt.Sprintf("(%d,%s,%s,%s,%s,%s,%s)", l.Id, l.Consumes, l.OperationId, l.ParamsId, l.Produces, l.Summary, l.Tags))
		}
	}
	repo.Dao.Debug().Exec(buffer.String())
	return len(lines)
}

func (repo *BaseDescribeRepoStr) BatchDeleteLine(ids []uint) int {
	repo.Dao.Debug().Delete(entity.LineDescribe{}, ids)
	return len(ids)
}

func (repo *BaseDescribeRepoStr) ClearLine() {
	repo.Dao.Debug().Where(" 1= 1").Delete(entity.LineDescribe{})
}

type ParamRepo interface {
	InsertParam(param *entity.Param) int
	UpdateParam(param *entity.Param) int
	DelParam(param *entity.Param) int
	BatchInsertParam(bases []*entity.Param) int
	BatchDeleteParam(ids []uint) int
	ClearParam()
}

func (repo *BaseDescribeRepoStr) InsertParam(param *entity.Param) int {
	repo.Dao.Debug().Create(param)
	return 1
}

func (repo *BaseDescribeRepoStr) UpdateParam(param *entity.Param) int {
	repo.Dao.Debug().Updates(param)
	return 1
}

func (repo *BaseDescribeRepoStr) DelParam(param *entity.Param) int {
	repo.Dao.Debug().Delete(param)
	return 1
}

func (repo *BaseDescribeRepoStr) BatchInsertParam(bases []*entity.Param) int {
	if len(bases) == 0 {
		return 0
	}
	var buffer bytes.Buffer
	sql := "INSERT INTO PARAM(id,description,format,in,name,required,type) VALUES"
	if _, err := buffer.WriteString(sql); err != nil {
		panic(err)
	}
	for i, p := range bases {
		if i == len(bases)-1 {
			buffer.WriteString(fmt.Sprintf("(%d,%s,%s,%s,%s,%s,%s),", p.Id, p.Description, p.Format, p.In, p.Name, p.Required, p.Types))
		}
		buffer.WriteString(fmt.Sprintf("(%d,%s,%s,%s,%s,%s,%s)", p.Id, p.Description, p.Format, p.In, p.Name, p.Required, p.Types))
	}
	repo.Dao.Debug().Exec(buffer.String())
	return len(bases)
}

func (repo *BaseDescribeRepoStr) BatchDeleteParam(ids []uint) int {
	repo.Dao.Debug().Delete(entity.Param{}, ids)
	return len(ids)
}

func (repo *BaseDescribeRepoStr) ClearParam() {
	repo.Dao.Debug().Where("1=1").Delete(entity.Param{})
}

type LineOfParamRepo interface {
	InsertLineOfParam(obj *entity.LineOfParam) int
	UpdateLineOfParam(obj *entity.LineOfParam) int
	DelLineOfParam(obj *entity.LineOfParam) int
	BatchInsertLineOfParam(bases []*entity.LineOfParam) int
	BatchDeleteLineOfParam(ids []uint) int
	ClearLineOfParam()
}

func (repo *BaseDescribeRepoStr) InsertLineOfParam(obj *entity.LineOfParam) int {
	repo.Dao.Debug().Create(obj)
	return 1
}

func (repo *BaseDescribeRepoStr) UpdateLineOfParam(obj *entity.LineOfParam) int {
	repo.Dao.Debug().Updates(obj)
	return 1
}

func (repo *BaseDescribeRepoStr) DelLineOfParam(obj *entity.LineOfParam) int {
	repo.Dao.Debug().Delete(obj)
	return 1
}

func (repo *BaseDescribeRepoStr) BatchInsertLineOfParam(bases []*entity.LineOfParam) int {
	if len(bases) == 0 {
		return 0
	}
	sql := "INSERT into line_of_param (id,line_id,param_id) values "
	var buffer bytes.Buffer
	if _, err := buffer.WriteString(sql); err != nil {
		panic(err)
	}
	for i, o := range bases {
		if i == len(bases)-1 {
			buffer.WriteString(fmt.Sprintf("(%d,%d,%d),", o.Id, o.LineId, o.ParamId))
		} else {
			buffer.WriteString(fmt.Sprintf("(%d,%d,%d)", o.Id, o.LineId, o.ParamId))
		}
	}
	repo.Dao.Debug().Exec(buffer.String())
	return len(bases)
}

func (repo *BaseDescribeRepoStr) BatchDeleteLineOfParam(ids []uint) int {
	repo.Dao.Debug().Delete(entity.LineOfParam{}, ids)
	return len(ids)
}

func (repo *BaseDescribeRepoStr) ClearLineOfParam() {
	repo.Dao.Debug().Where("1=1").Delete(entity.LineOfParam{})
}
