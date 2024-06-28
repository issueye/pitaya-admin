package service

import (
	"fmt"

	"github.com/issueye/pitaya_admin/internal/common/model"
	"github.com/issueye/pitaya_admin/internal/common/service"
)

type Develop struct {
	service.BaseService
}

func (owner *Develop) SetBase(base service.BaseService) {
	owner.BaseService = base
}

func NewDevelop(args ...service.ServiceContext) *Develop {
	return service.NewService(&Develop{}, args...)
}

// GetAllTable
// 获取所有 sqlite 的数据表
func (s *Develop) GetAllTable() (res []string, err error) {
	err = s.GetDB().Raw("SELECT name FROM sqlite_master WHERE type='table';").Scan(&res).Error
	return
}

// 获取所有 sqlite 的数据表
func (s *Develop) GetTableFields(name string) (data *model.SqliteTableInfo, err error) {
	data = new(model.SqliteTableInfo)
	data.TableName = name
	data.Fields = make([]*model.SqliteTableFieldInfo, 0)
	err = s.GetDB().Raw(fmt.Sprintf("PRAGMA table_info(%s)", name)).Scan(&data.Fields).Error
	return
}
