package backend

import (
	"linux/GetSystemInfo/tools"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 日志接口
type LogController struct {
}

// 获取接口请求日志
func (loginfo LogController) HttpRequest(c *gin.Context) {
	path := "log/"
	resLog := "_request.log"
	logFilePath := path + tools.FormatDate(time.Now(), tools.YYYYMMDD) + resLog
	requestSlic, _ := tools.ReadFile(logFilePath)

	c.HTML(http.StatusOK, "view/log/apiRequest.html", gin.H{
		"title":       "api接口日志",
		"httpRequest": requestSlic,
	})
}

// 获取程序运行错误日志
func (loginfo LogController) ErrLog(c *gin.Context) {
	path := "log/"
	errLog := "_err.log"
	logFilePath := path + tools.FormatDate(time.Now(), tools.YYYYMMDD) + errLog
	errLogSlic, _ := tools.ReadFile(logFilePath)

	c.HTML(http.StatusOK, "view/log/errLog.html", gin.H{
		"title":  "错误日志",
		"errLog": errLogSlic,
	})
}
