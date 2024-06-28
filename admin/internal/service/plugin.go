package service

import (
	"github.com/issueye/pitaya_admin/internal/common/model"
	"github.com/issueye/pitaya_admin/internal/common/service"
	"github.com/issueye/pitaya_admin/internal/repository"
)

type Plugin struct {
	service.BaseService
}

func (owner *Plugin) Self() *Plugin {
	return owner
}

func (owner *Plugin) SetBase(base service.BaseService) {
	owner.BaseService = base
}

func NewPlugin(args ...service.ServiceContext) *Plugin {
	return service.NewService(&Plugin{}, args...)
}

// Create
// 创建信息
func (s *Plugin) Create(data *repository.CreatePlugin) error {
	info := model.PluginInfo{}.New()
	info.Name = data.Name
	info.Path = data.Path
	info.Version = data.Version
	info.Key = data.Key
	info.Value = data.Value

	return s.GetDB().Model(info).Create(info).Error
}
