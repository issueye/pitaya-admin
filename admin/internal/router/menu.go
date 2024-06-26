package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/issueye/pitaya_admin/internal/controller/v1"
	"github.com/issueye/pitaya_admin/internal/global"
)

type MenuRouter struct {
	Name    string
	control *v1.MenuController
}

func NewMenuRouter() *MenuRouter {
	return &MenuRouter{
		Name:    string(global.RGN_menu),
		control: v1.NewMenuController(),
	}
}

func (Menu MenuRouter) Register(group *gin.RouterGroup, auth gin.HandlerFunc) {
	f := group.Group(Menu.Name, auth)
	f.GET("", Menu.control.List)
	f.GET("tree", Menu.control.TreeList)
	f.GET(":id", Menu.control.GetById)
	f.POST("", Menu.control.Create)
	f.PUT(":id", Menu.control.Modify)
	f.PUT("state/:id", Menu.control.ModifyState)
	f.DELETE(":id", Menu.control.Delete)
}

type MenuGroupRouter struct {
	Name    string
	control *v1.GroupMenuController
}

func NewMenuGroupRouter() *MenuGroupRouter {
	return &MenuGroupRouter{
		Name:    string(global.RGN_menuGroup),
		control: v1.NewGroupMenuController(),
	}
}

func (menu *MenuGroupRouter) Register(group *gin.RouterGroup, auth gin.HandlerFunc) {
	f := group.Group(menu.Name, auth)
	f.GET("", menu.control.List)
	f.GET("getMenu/:groupId", menu.control.GetMenu)
	f.GET(":id", menu.control.GetById)
	f.POST("", menu.control.Create)
	f.PUT(":id", menu.control.Modify)
	f.PUT("state/:id", menu.control.ModifyState)
	f.PUT("auth/:id", menu.control.ModifyGroupMenuState)
	f.DELETE(":id", menu.control.Delete)
}
