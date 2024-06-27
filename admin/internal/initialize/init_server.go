package initialize

import (
	"fmt"
	"mime"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/issueye/pitaya_admin/internal/config"
	"github.com/issueye/pitaya_admin/internal/global"
	internalMid "github.com/issueye/pitaya_admin/internal/middleware"
	"github.com/issueye/pitaya_admin/internal/router"
	"github.com/issueye/pitaya_admin/pkg/middleware"
	orange_validator "github.com/issueye/pitaya_admin/pkg/validator"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/antage/eventsource.v1"
)

func InitServer() {
	mode := config.GetParam(config.CfgServerMode, "release").String()
	gin.SetMode(mode)

	// gin引擎对象
	global.Router = gin.New()
	// 注册一个form表单验证器
	orange_validator.RegisterValidator()

	global.SSE = eventsource.New(nil, nil)

	// 加载中间件
	global.Router.Use(middleware.CORSMiddleware([]string{}))       // 处理前端跨域
	global.Router.Use(middleware.GinLogger(global.Logger))         // 日志记录
	global.Router.Use(internalMid.Intercept(global.Router))        // 处理问题
	global.Router.Use(middleware.GinRecovery(global.Logger, true)) // 服务恐慌处理

	// staticFs, err := fs.Sub(global.PageStatic, "admin")
	// if err != nil {
	// 	panic(err)
	// }
	// global.Router.StaticFS("/admin", http.FS(staticFs))

	// 设置一个静态文件服务器
	global.Router.Static("/www", "./runtime/static")
	global.Router.Static("/resources", "./runtime/static/resources")

	if strings.ToLower(mode) == "debug" {
		// 设置 swagger
		global.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 告诉服务文件的MIME类型
	_ = mime.AddExtensionType(".js", "application/javascript")
	_ = mime.AddExtensionType(".css", "text/css")
	_ = mime.AddExtensionType(".eot", "text/plain")
	_ = mime.AddExtensionType(".woff", "application/font-woff")
	_ = mime.AddExtensionType(".woff2", "application/font-woff2")
	_ = mime.AddExtensionType(".ttf", "application/font-ttf")
	_ = mime.AddExtensionType(".eot", "application/vnd.ms-fontobject")

	// 注册路由
	router.InitRouter(global.Router)
	// 端口号为命令行提供
	port := config.GetParam(config.CfgServerPort, "10065").Int()
	global.HttpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: global.Router,
	}
}
