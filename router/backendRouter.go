package router

import (
	"linux/GetSystemInfo/controller/backend"

	"github.com/gin-gonic/gin"
)

// 后端分组路由
func BackendRouter(r *gin.Engine) {
	// 日志接口
	backendApi := r.Group("/admin")
	{
		backendApi.GET("/log/requestapi", backend.LogController{}.HttpRequest)
		backendApi.GET("/log/errlog", backend.LogController{}.ErrLog)
	}

	// 云服务
	cloudApi := r.Group("cloud")
	{
		cloudApi.GET("/container/svc", backend.CloudController{}.DockerService)
		cloudApi.GET("/container/listRunning", backend.CloudController{}.ContainerListRunning)
		cloudApi.GET("/container/listExit", backend.CloudController{}.ContainerListExit)
		cloudApi.GET("/kubernetes/svc", backend.CloudController{}.KubernetesService)
		cloudApi.GET("/kubernetes/pod", backend.CloudController{}.PodService)
	}
}
