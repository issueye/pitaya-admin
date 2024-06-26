package main

import (
	"embed"
	_ "embed"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/issueye/pitaya_admin/docs"
	"github.com/issueye/pitaya_admin/internal/global"
	"github.com/issueye/pitaya_admin/internal/initialize"
)

//	@title			代理管理服务
//	@version		V0.1
//	@description	代理管理服务

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//go:embed admin/*
var page embed.FS

func main() {

	global.PageStatic = page
	initialize.Initialize()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
