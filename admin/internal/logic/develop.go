package logic

import (
	"github.com/issueye/pitaya_admin/internal/common/model"
	"github.com/issueye/pitaya_admin/internal/service"
)

type Develop struct{}

// GetTables
// 获取所有表
func (d *Develop) GetTables() (tables []string, err error) {
	devSrv := service.NewDevelop()
	tables, err = devSrv.GetAllTable()
	return
}

// GetTableFields
// 获取表字段
func (d *Develop) GetTableFields(table string) (data *model.SqliteTableInfo, err error) {
	devSrv := service.NewDevelop()
	data, err = devSrv.GetTableFields(table)
	return
}
