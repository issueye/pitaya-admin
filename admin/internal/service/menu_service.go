package service

import (
	"fmt"

	"github.com/issueye/pitaya_admin/internal/common/model"
	"github.com/issueye/pitaya_admin/internal/common/service"
	"github.com/issueye/pitaya_admin/internal/repository"
	"gorm.io/gorm"
)

type Menu struct {
	service.BaseService
}

func (owner *Menu) Self() *Menu {
	return owner
}

func (owner *Menu) SetBase(base service.BaseService) {
	owner.BaseService = base
}

func NewMenu(args ...service.ServiceContext) *Menu {
	return service.NewService(&Menu{}, args...)
}

// CreateAdminNonExistent
// 创建管理员用户，如果不存在
func (Menu *Menu) CreateAdminNonExistent() error {
	// isHave := int64(0)
	// err := Menu.Db.Model(&model.Menu{}).Where("account = ?", global.AdminName).Where("id = ?", global.AdminId).Count(&isHave).Error
	// if err != nil {
	// 	return err
	// }

	// if isHave == 0 {
	// 	info := new(model.Menu)
	// 	info.ID = global.AdminId
	// 	info.Name =

	// 	return Menu.Db.Create(info).Error
	// } else {
	// 	return nil
	// }

	return nil
}

// Create
// 创建用户信息
func (Menu *Menu) Create(data *repository.CreateMenu) error {
	info := model.Menu{}.New()
	info.Copy(&data.MenuBase)

	return Menu.GetDB().Create(info).Error
}

// Create
// 创建用户信息
func (Menu *Menu) CreateNoExistent(data *model.MenuBase) string {
	info := new(model.Menu)
	err := Menu.GetDB().Model(info).Where("route = ?", data.Route).Where("level = ?", data.Level).Find(info).Error
	if err != nil {
		panic("查询菜单失败 " + err.Error())
	}

	if info.ID != "" {
		return info.ID
	}

	menuInfo := model.Menu{}.New()
	menuInfo.Copy(data)

	err = Menu.GetDB().Create(menuInfo).Error
	if err != nil {
		panic("创建菜单失败 " + err.Error())
	}

	return menuInfo.ID
}

// GetById
// 根据用户ID查找用户信息
func (Menu *Menu) GetById(id string) (*model.Menu, error) {
	info := new(model.Menu)
	err := Menu.GetDB().Model(info).Where("id = ?", id).Find(info).Error
	return info, err
}

// Modify
// 修改用户信息
func (Menu *Menu) Modify(id string, info *repository.ModifyMenu) error {
	m := make(map[string]any)
	m["name"] = info.Name
	m["route"] = info.Route
	m["title"] = info.Title
	m["icon"] = info.Icon
	m["auth"] = info.Auth
	m["level"] = info.Level
	m["parent_id"] = info.ParentId

	return Menu.GetDB().Model(&model.Menu{}).Where("id = ?", id).Updates(m).Error
}

// Status
// 修改用户信息
func (Menu *Menu) Status(info *repository.StatusMenu) error {
	return Menu.GetDB().
		Model(&model.Menu{}).
		Where("id = ?", info.ID).
		Update("state", info.State).
		Error
}

// Delete
// 删除用户信息
func (Menu *Menu) Delete(id string) error {
	return Menu.GetDB().Where("id = ?", id).Delete(&model.Menu{}).Error
}

// List
// 获取用户列表
func (Menu *Menu) List(info *repository.QueryMenu) ([]*model.Menu, error) {
	MenuInfo := new(model.Menu)
	list := make([]*model.Menu, 0)
	err := Menu.DataFilter(MenuInfo.TableName(), info, &list, func(db *gorm.DB) (*gorm.DB, error) {
		query := db.Order("[order]")

		// 通用统一条件
		if info.Condition != "" {
			query = query.
				Where("name like ?", fmt.Sprintf("%%%s%%", info.Condition)).
				Or("title like ?", fmt.Sprintf("%%%s%%", info.Condition)).
				Or("route like ?", fmt.Sprintf("%%%s%%", info.Condition))
		}

		if info.Level > -1 {
			query = query.Where("level = ?", info.Level)
		}

		return query, nil
	})
	return list, err
}

// List
// 获取用户列表
func (Menu *Menu) TreeList(info *repository.QueryMenu) ([]*repository.ResTreeMenus, error) {
	MenuInfo := new(model.Menu)
	list := make([]*model.Menu, 0)
	err := Menu.DataFilter(MenuInfo.TableName(), info, &list, func(db *gorm.DB) (*gorm.DB, error) {
		query := db.Order("[order]")

		// 通用统一条件
		if info.Condition != "" {
			query = query.
				Where("name like ?", fmt.Sprintf("%%%s%%", info.Condition)).
				Or("title like ?", fmt.Sprintf("%%%s%%", info.Condition)).
				Or("route like ?", fmt.Sprintf("%%%s%%", info.Condition))
		}

		if info.Level > -1 {
			query = query.Where("level = ?", info.Level)
		}

		return query, nil
	})

	if err != nil {
		return nil, err
	}

	resDatas := make([]*repository.ResTreeMenus, 0)

	findFirst := func(id string) *repository.ResTreeMenus {
		for _, element := range resDatas {
			if element.ID == id {
				return element
			}
		}

		return nil
	}

	for _, element := range list {
		if element.Level == 0 {
			menu := new(repository.ResTreeMenus)
			menu.MenuBase.Copy(&element.MenuBase)
			menu.ID = element.ID
			resDatas = append(resDatas, menu)
		}
	}

	for _, element := range list {
		if element.Level == 1 {
			menu := new(repository.ResTreeMenus)
			menu.MenuBase.Copy(&element.MenuBase)
			menu.ID = element.ID
			parent := findFirst(menu.ParentId)
			if parent != nil {
				parent.Children = append(parent.Children, &menu.Menu)
			}
		}
	}

	for _, element := range resDatas {
		if len(element.Children) > 0 {
			element.HasChildren = true
		}
	}

	return resDatas, nil
}
