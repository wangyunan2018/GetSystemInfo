package router

import (
	"linux/GetSystemInfo/controller/frontend"

	"github.com/gin-gonic/gin"
)

// 前端路由
func FrontEndRoutes(r *gin.Engine) {
	// 分组路由
	frontEnd := r.Group("/")
	{
		frontEnd.GET("/", frontend.FrontendController{}.Home)
		frontEnd.GET("/cpuinfo", frontend.LeftMenuController{}.CpuInfo)
		frontEnd.GET("/diskinfo", frontend.LeftMenuController{}.DiskInfo)
		frontEnd.GET("/gpuinfo", frontend.LeftMenuController{}.GpuInfo)
		frontEnd.GET("/loginfo", frontend.LeftMenuController{}.LogInfo)
		frontEnd.GET("/meminfo", frontend.LeftMenuController{}.MemInfo)
		frontEnd.GET("/container", frontend.LeftMenuController{}.ContainerInfo)
		frontEnd.GET("/kubernetes", frontend.LeftMenuController{}.KubernetsInfo)
	}
}
