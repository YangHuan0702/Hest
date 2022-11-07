package repository

import (
	"Hest/model"
	"bytes"
	"fmt"
)
import "Hest/init"

type BaseDescribe model.BaseDescribe
type LineDescribe model.LineDescribe
type Param model.Param
type LineOfParam model.LineOfParam

type BaseDescribeRepo interface {
	Insert() int
	Update() int
	Del() int
	BatchInsert(bases []*BaseDescribe) int
	BatchDelete(ids []uint) int
}

func (obj *BaseDescribe) Insert() int {
	init.DbCli.Debug().Create(obj)
	return 1
}

func (obj *BaseDescribe) Update() int {
	init.DbCli.Debug().Save(obj)
	return 1
}

func (obj *BaseDescribe) Del() int {
	init.DbCli.Debug().Delete(obj).Unscoped()
	return 1
}

func (obj *BaseDescribe) BatchInsert(bases []*BaseDescribe) int {
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
	init.DbCli.Debug().Exec(buffer.String())
	return len(bases)
}

func (obj *BaseDescribe) BatchDelete(ids []uint) int {
	init.DbCli.Debug().Delete(obj, ids)
	return len(ids)
}

type LineDescribeRepo interface {
	Insert() int
	Update() int
	Del() int
	BatchInsert(bases []*LineDescribe) int
	BatchDelete(ids []uint) int
}

func (obj *LineDescribe) Insert() int {
	init.DbCli.Debug().Create(obj)
	return 1
}

func (obj *LineDescribe) Update() int {
	init.DbCli.Debug().Updates(obj)
	return 1
}

func (obj *LineDescribe) Del() int {
	init.DbCli.Debug().Delete(obj)
	return 1
}

func (obj *LineDescribe) BatchInsert(lines []*LineDescribe) int {
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
	init.DbCli.Debug().Exec(buffer.String())
	return len(lines)
}

func (obj *LineDescribe) BatchDelete(ids []uint) int {
	init.DbCli.Debug().Delete(obj, ids)
	return len(ids)
}

type ParamRepo interface {
	Insert() int
	Update() int
	Del() int
	BatchInsert(bases []*Param) int
	BatchDelete(ids []uint) int
}

func (obj *Param) Insert() int {
	init.DbCli.Debug().Create(obj)
	return 1
}

func (obj *Param) Update() int {
	init.DbCli.Debug().Updates(obj)
	return 1
}

func (obj *Param) Del() int {
	init.DbCli.Debug().Delete(obj)
	return 1
}

func (obj *Param) BatchInsert(bases []*Param) int {
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
	init.DbCli.Debug().Exec(buffer.String())
	return len(bases)
}

func (obj *Param) BatchDelete(ids []uint) int {
	init.DbCli.Debug().Delete(obj, ids)
	return len(ids)
}

type LineOfParamRepo interface {
	Insert() int
	Update() int
	Del() int
	BatchInsert(bases []*LineOfParam) int
	BatchDelete(ids []uint) int
}

func (obj *LineOfParam) Insert() int {
	init.DbCli.Debug().Create(obj)
	return 1
}

func (obj *LineOfParam) Update() int {
	if obj.Id == 0 {
		return 0
	}
	init.DbCli.Debug().Updates(obj)
	return 1
}

func (obj *LineOfParam) Del() int {
	init.DbCli.Debug().Delete(obj)
	return 1
}

func (obj *LineOfParam) BatchInsert(bases []*LineOfParam) int {
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
			buffer.WriteString(fmt.Sprint("(%d,%s,%s),", o.Id, o.LineId, o.ParamId))
		} else {
			buffer.WriteString(fmt.Sprint("(%d,%s,%s)", o.Id, o.LineId, o.ParamId))
		}
	}
	init.DbCli.Debug().Exec(buffer.String())
	return len(bases)
}

func (obj *LineOfParam) BatchDelete(ids []uint) int {
	init.DbCli.Debug().Delete(obj, ids)
	return len(ids)
}
