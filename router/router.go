package router

import (
	"io"
	"linux/GetSystemInfo/tools"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() {
	// 把控制台请求日志输出到文件
	path := "log/"
	logFile, _ := os.Create(path + tools.FormatDate(time.Now(), tools.YYYYMMDD) + "_request.log")
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// 加载gin默认路由引擎
	r := gin.Default()

	// 加载静态资源
	r.LoadHTMLGlob("view/**/**")
	r.Static("/static", "static")

	// 前端路由
	FrontEndRoutes(r)
	// 后端路由
	BackendRouter(r)

	r.Run(":8080")
}
