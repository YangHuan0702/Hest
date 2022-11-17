package repository

import (
	"Hest/model/entity"
	"bytes"
	"fmt"
	"gorm.io/gorm"
)

type BaseDescribe entity.BaseDescribe
type LineDescribe entity.LineDescribe
type Param entity.Param
type LineOfParam entity.LineOfParam

type BaseDescribeRepoStr struct {
	dao *gorm.DB
}

type BaseDescribeRepo interface {
	InsertBase(base *BaseDescribe) int
	UpdateBase(base *BaseDescribe) int
	DelBase(base *BaseDescribe) int
	BatchInsertBase(bases []*BaseDescribe) int
	BatchDeleteBase(ids []uint) int
}

func (repo *BaseDescribeRepoStr) InsertBase(base *BaseDescribe) int {
	repo.dao.Debug().Create(base)
	return 1
}

func (repo *BaseDescribeRepoStr) UpdateBase(base *BaseDescribe) int {
	repo.dao.Debug().Save(base)
	return 1
}

func (repo *BaseDescribeRepoStr) DelBase(base *BaseDescribe) int {
	repo.dao.Debug().Delete(base).Unscoped()
	return 1
}

func (repo *BaseDescribeRepoStr) BatchInsertBase(bases []*BaseDescribe) int {
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
	repo.dao.Debug().Exec(buffer.String())
	return len(bases)
}

func (repo *BaseDescribeRepoStr) BatchDeleteBase(ids []uint) int {
	repo.dao.Debug().Delete(BaseDescribe{}, ids)
	return len(ids)
}

type LineDescribeRepo interface {
	Insert(line *LineDescribe) int
	Update(line *LineDescribe) int
	Del(line *LineDescribe) int
	BatchInsert(bases []*LineDescribe) int
	BatchDelete(ids []uint) int
}

func (repo *BaseDescribeRepoStr) Insert(line *LineDescribe) int {
	repo.dao.Debug().Create(line)
	return 1
}

func (repo *BaseDescribeRepoStr) Update(line *LineDescribe) int {
	repo.dao.Debug().Updates(line)
	return 1
}

func (repo *BaseDescribeRepoStr) Del(line *LineDescribe) int {
	repo.dao.Debug().Delete(line)
	return 1
}

func (repo *BaseDescribeRepoStr) BatchInsert(lines []*LineDescribe) int {
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
	repo.dao.Debug().Exec(buffer.String())
	return len(lines)
}

func (repo *BaseDescribeRepoStr) BatchDelete(ids []uint) int {
	repo.dao.Debug().Delete(LineDescribe{}, ids)
	return len(ids)
}

type ParamRepo interface {
	InsertParam(param *Param) int
	UpdateParam(param *Param) int
	DelParam(param *Param) int
	BatchInsertParam(bases []*Param) int
	BatchDeleteParam(ids []uint) int
}

func (repo *BaseDescribeRepoStr) InsertParam(param *Param) int {
	repo.dao.Debug().Create(param)
	return 1
}

func (repo *BaseDescribeRepoStr) UpdateParam(param *Param) int {
	repo.dao.Debug().Updates(param)
	return 1
}

func (repo *BaseDescribeRepoStr) DelParam(param *Param) int {
	repo.dao.Debug().Delete(param)
	return 1
}

func (repo *BaseDescribeRepoStr) BatchInsertParam(bases []*Param) int {
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
	repo.dao.Debug().Exec(buffer.String())
	return len(bases)
}

func (repo *BaseDescribeRepoStr) BatchDeleteParam(ids []uint) int {
	repo.dao.Debug().Delete(Param{}, ids)
	return len(ids)
}

type LineOfParamRepo interface {
	InsertLineOfParam(obj *LineOfParam) int
	UpdateLineOfParam(obj *LineOfParam) int
	DelLineOfParam(obj *LineOfParam) int
	BatchInsertLineOfParam(bases []*LineOfParam) int
	BatchDeleteLineOfParam(ids []uint) int
}

func (repo *BaseDescribeRepoStr) InsertLineOfParam(obj *LineOfParam) int {
	repo.dao.Debug().Create(obj)
	return 1
}

func (repo *BaseDescribeRepoStr) UpdateLineOfParam(obj *LineOfParam) int {
	repo.dao.Debug().Updates(obj)
	return 1
}

func (repo *BaseDescribeRepoStr) DelLineOfParam(obj *LineOfParam) int {
	repo.dao.Debug().Delete(obj)
	return 1
}

func (repo *BaseDescribeRepoStr) BatchInsertLineOfParam(bases []*LineOfParam) int {
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
	repo.dao.Debug().Exec(buffer.String())
	return len(bases)
}

func (repo *BaseDescribeRepoStr) BatchDeleteLineOfParam(ids []uint) int {
	repo.dao.Debug().Delete(LineOfParam{}, ids)
	return len(ids)
}
