package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/issueye/pitaya_admin/internal/common/controller"
	"github.com/issueye/pitaya_admin/internal/logic"
)

type DevelopController struct{}

// GetTables doc
//
//	@tags			开发管理
//	@Summary		获取所有表
//	@Description	获取所有表
//	@Produce		json
//	@Success		200		{object}	controller.Full{[]model.Menu}	"code: 200 成功"
//	@Failure		500		{object}	controller.Base					"错误返回内容"
//	@Router			/api/v1/develop/getAllTable [get]
//	@Security		ApiKeyAuth
func (DevelopController) GetTables(ctx *gin.Context) {
	control := controller.NewA(ctx)

	lc := &logic.Develop{}

	tables, err := lc.GetTables()
	if err != nil {
		control.FailByMsgf("获取表失败: %s", err.Error())
		return
	}

	control.SuccessData(tables)
}

// GetTableFields doc
//
//	@tags			开发管理
//	@Summary		获取表所有字段
//	@Description	获取表所有字段
//	@Produce		json
//	@Success		200		{object}	controller.Full{[]model.Menu}	"code: 200 成功"
//	@Failure		500		{object}	controller.Base					"错误返回内容"
//	@Router			/api/v1/develop/getTableColumns/{table} [get]
//	@Security		ApiKeyAuth
func (DevelopController) GetTableFields(ctx *gin.Context) {
	control := controller.NewA(ctx)

	table := control.Param("table")
	if table == "" {
		control.FailByMsgf("参数错误")
		return
	}

	lc := &logic.Develop{}

	tables, err := lc.GetTableFields(table)
	if err != nil {
		control.FailByMsgf("获取表失败: %s", err.Error())
		return
	}

	control.SuccessData(tables)
}
