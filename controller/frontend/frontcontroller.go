package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FrontendController struct {
}

func (fc FrontendController) Home(c *gin.Context) {

	c.HTML(http.StatusOK, "view/default/home.html", gin.H{
		"title": "服务器硬件资源实时监控",
	})
}
