package repository

import "github.com/issueye/pitaya_admin/internal/common/model"

type CreatePlugin struct {
	model.PluginBase
}

type ModifyPlugin struct {
	model.PluginInfo
}
