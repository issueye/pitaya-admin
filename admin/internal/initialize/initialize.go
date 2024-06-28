package initialize

import (
	"context"
	"fmt"

	"github.com/issueye/pitaya_admin/internal/config"
	"github.com/issueye/pitaya_admin/internal/global"
)

func Initialize() {
	ctx := context.Background()
	// 初始化运行文件
	InitRuntime()
	// 配置参数
	config.InitConfig()
	// 日志
	InitLog()
	// 数据
	InitData()
	// http服务
	InitServer()
	// 启动服务
	ShowInfo()
	// 监听服务
	err := global.HttpServer.ListenAndServe()
	if err != nil {
		fmt.Printf("启动服务失败：%v", err)
	}

	// 关闭服务
	global.HttpServer.Shutdown(ctx)

	// 关闭日志
	global.Logger.Sync()
	// 关闭监听
	ctx.Done()
}

var (
	AppName string
	Branch  string
	Commit  string
	Date    string
	Version string
)

func ShowInfo() {
	bannerStr := `
	▀      ▄                                           █           ▀          
	▄▄▄▄   ▄▄▄    ▄▄█▄▄   ▄▄▄   ▄   ▄   ▄▄▄           ▄▄▄    ▄▄▄█  ▄▄▄▄▄  ▄▄▄    ▄ ▄▄  
	█▀ ▀█    █      █    ▀   █  ▀▄ ▄▀  ▀   █         ▀   █  █▀ ▀█  █ █ █    █    █▀  █ 
	█   █    █      █    ▄▀▀▀█   █▄█   ▄▀▀▀█   ▀▀▀   ▄▀▀▀█  █   █  █ █ █    █    █   █ 
	██▄█▀  ▄▄█▄▄    ▀▄▄  ▀▄▄▀█   ▀█    ▀▄▄▀█         ▀▄▄▀█  ▀█▄██  █ █ █  ▄▄█▄▄  █   █ 
	█                            ▄▀                                                    
	▀                           ▀▀                                                     
   
	一个小工具 pitaya-admin
	`
	fmt.Println(bannerStr) // mona12 风格

	info := `
	AppName: %s
	Branch : %s
	Commit : %s
	Date   : %s
	Version: %s
	
	`
	fmt.Printf(info+"\n", AppName, Branch, Commit, Date, Version)
}
