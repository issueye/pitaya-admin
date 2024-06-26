package repository

import "github.com/issueye/pitaya_admin/internal/common/model"

type CreateCert struct {
	Name    string `json:"name" binding:"required" label:"证书名称"` // 名称
	Public  string `json:"public" label:"公钥地址"`                  // 公有证书路径
	Private string `json:"private" label:"私钥地址"`                 // 私有证书路径
	Mark    string `json:"mark" label:"备注"`                      // 备注
}

type ModifyCert struct {
	ID      string `json:"id" binding:"required" label:"编码"`     // 编码
	Name    string `json:"name" binding:"required" label:"证书名称"` // 名称
	Public  string `json:"public" label:"公钥地址"`                  // 公有证书路径
	Private string `json:"private" label:"私钥地址"`                 // 私有证书路径
	Mark    string `json:"mark" label:"备注"`                      // 备注
}

// 查询信息
type QueryCert struct {
	Condition string `json:"condition" form:"condition"` // 条件
	model.Page
}
