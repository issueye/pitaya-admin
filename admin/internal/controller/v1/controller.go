package v1

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/issueye/pitaya_admin/internal/common/controller"
)

// DepositoryPost doc
//
//	@tags		    http报文信息
//	@Summary		查询HTTP请求信息
//	@Description	查询HTTP请求信息
//	@Produce		json
//	@Success		200		{object}	controller.Base			"code: 200 成功"
//	@Failure		500		{object}	controller.Base			"错误返回内容"
//	@Router			/api/v1/depository [post]
//	@Security		ApiKeyAuth
func DepositoryPost(ctx *gin.Context) {
	c := controller.NewA(ctx)

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.FailByMsgf("读取数据失败: %v", err)
		return
	}

	fmt.Println("data", string(data))

	c.Success()
}
