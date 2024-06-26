package initialize

import (
	"fmt"
	"path/filepath"

	"github.com/issueye/pitaya_admin/pkg/utils"
)

func InitRuntime() {
	// 检查本地是否存在runtime文件夹
	// 获取当前程序的路径
	path := utils.GetWorkDir()
	rtPath := isExistsCreatePath(path, "runtime")
	dataPath := isExistsCreatePath(rtPath, "data")

	// 存放 http 报文数据
	isExistsCreatePath(dataPath, "http")

	isExistsCreatePath(rtPath, "logs")
	staticPath := isExistsCreatePath(rtPath, "static")
	isExistsCreatePath(staticPath, "admin")
	isExistsCreatePath(staticPath, "resources")
}

func isExistsCreatePath(path, name string) string {
	p := filepath.Join(path, name)
	exists, err := utils.PathExists(p)
	if err != nil {
		panic(err.Error())
	}

	if !exists {
		panic(fmt.Errorf("创建【%s】文件夹失败", name))
	}

	return p
}
