package service

import (
	"fmt"

	"github.com/issueye/pitaya_admin/internal/common/model"
	"github.com/issueye/pitaya_admin/internal/common/service"
	"github.com/issueye/pitaya_admin/internal/repository"
	"gorm.io/gorm"
)

type GroupMenu struct {
	service.BaseService
}

func (owner *GroupMenu) Self() *GroupMenu {
	return owner
}

func (owner *GroupMenu) SetBase(base service.BaseService) {
	owner.BaseService = base
}

func NewGroupMenu(args ...service.ServiceContext) *GroupMenu {
	return service.NewServiceSelf(&GroupMenu{}, args...)
}

// Create
// 创建用户组菜单信息
func (GroupMenu *GroupMenu) Create(data *repository.CreateGroupMenu) error {
	info := model.GroupMenu{}.New()
	info.Copy(&data.GroupMenuBase)
	info.State = 1
	return GroupMenu.GetDB().Create(info).Error
}

// Create
// 创建用户组菜单信息
func (GroupMenu *GroupMenu) BatchCreate(datas *[]*model.GroupMenu) error {
	return GroupMenu.GetDB().Create(datas).Error
}

// GetById
// 根据用户组菜单ID查找用户组菜单信息
func (GroupMenu *GroupMenu) GetById(id string) (*model.GroupMenu, error) {
	info := new(model.GroupMenu)
	err := GroupMenu.GetDB().Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// Modify
// 修改用户组菜单信息
func (GroupMenu *GroupMenu) Modify(id string, info *repository.ModifyGroupMenu) error {
	m := make(map[string]any)
	m["name"] = info.Name
	m["route"] = info.Route
	m["title"] = info.Title
	m["icon"] = info.Icon
	m["auth"] = info.Auth
	m["level"] = info.Level
	m["parent_id"] = info.ParentId

	return GroupMenu.GetDB().Model(&model.GroupMenu{}).Where("id = ?", id).Updates(m).Error
}

// Status
// 修改用户组菜单信息
func (GroupMenu *GroupMenu) Status(info *repository.StatusGroupMenu) error {
	return GroupMenu.GetDB().
		Model(&model.GroupMenu{}).
		Where("id = ?", info.ID).
		Update("state", info.State).
		Error
}

// Status
// 修改用户组菜单信息
func (GroupMenu *GroupMenu) ModifyStateByGroupId(groupId string, state uint) error {
	return GroupMenu.GetDB().
		Model(&model.GroupMenu{}).
		Where("group_id = ?", groupId).
		Update("state", state).
		Error
}

// Delete
// 删除用户组菜单信息
func (GroupMenu *GroupMenu) Delete(id string) error {
	return GroupMenu.GetDB().Where("id = ?", id).Delete(&model.GroupMenu{}).Error
}

// Delete
// 删除用户组菜单信息
func (GroupMenu *GroupMenu) DelByGroupId(id string) error {
	return GroupMenu.GetDB().Where("group_id = ?", id).Delete(&model.GroupMenu{}).Error
}

// List
// 获取用户组菜单列表
func (GroupMenu *GroupMenu) List(info *repository.QueryGroupMenu) ([]*model.GroupMenu, error) {
	GroupMenuInfo := new(model.GroupMenu)
	list := make([]*model.GroupMenu, 0)
	err := GroupMenu.DataFilter(GroupMenuInfo.TableName(), info, &list, func(db *gorm.DB) (*gorm.DB, error) {
		query := db.Order("level").Order("[order]")

		if info.GroupId != "" {
			query = query.Where("group_id = ?", info.GroupId)
		}

		if info.Condition != "" {
			query = query.
				Where("name like ?", fmt.Sprintf("%%%s%%", info.Condition)).
				Or("title like ?", fmt.Sprintf("%%%s%%", info.Condition)).
				Or("route like ?", fmt.Sprintf("%%%s%%", info.Condition))
		}

		return query, nil
	})
	return list, err
}
