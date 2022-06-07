package backend

import (
	"fmt"
	"linux/GetSystemInfo/cmd"
	"linux/GetSystemInfo/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CloudController struct {
}

// Docker服务
func (dc CloudController) DockerService(c *gin.Context) {
	c.HTML(http.StatusOK, "view/cloud/containerService.html", gin.H{
		"title": "Docker Service",
	})
}

// 在线容器
func (dc CloudController) ContainerListRunning(c *gin.Context) {
	// 获取容器信息
	GetContainerInfoCmd := "sudo chmod +x ./sh/getContainerStatus.sh && sudo /bin/bash ./sh/getContainerStatus.sh"
	_, _, err := cmd.ShellOut(GetContainerInfoCmd)
	if err != nil {
		errOut := "脚本 [getContainerStatus.sh] 执行失败，请检查。"
		tools.LogFileErr(errOut)
		fmt.Println(err.Error())
	}

	containerRunNamePath := "./textfile/containerName.txt"
	containerRunName, err := tools.ReadFile(containerRunNamePath)
	if len(containerRunName) == 0 || err != nil {
		errOut := "没有获取到容器名，请登录到服务器检查。"
		tools.LogFileErr(errOut)
	}

	containerRunIdPath := "./textfile/dockerRunId.txt"
	containerRunId, err := tools.ReadFile(containerRunIdPath)
	if len(containerRunId) == 0 || err != nil {
		errOut := "没有获取到容器ID，请登录到服务器检查。"
		tools.LogFileErr(errOut)
	}

	// 遍历ContainerStatus
	containerStatusPath := "./textfile/dockerStatus.txt"
	containerStatus, err := tools.ReadFile(containerStatusPath)
	if len(containerStatus) == 0 || err != nil {
		errOut := "没有获取到容器运行状态，请登录到服务器检查。"
		tools.LogFileErr(errOut)
	}

	// 在线容器操作 重启、停止、删除、日志
	//重启
	containerRunResPath := "./textfile/dockerRunRes.txt"
	containerRunRes, err := tools.ReadFile(containerRunResPath)
	if len(containerRunRes) == 0 || err != nil {
		errOut := "没有获取到容器（Ruuning状态）重启按钮，请检查。"
		tools.ReadFile(errOut)
	}
	// 停止
	containerRunStopPath := "./textfile/dockerRunStop.txt"
	containerRunStop, err := tools.ReadFile(containerRunStopPath)
	if len(containerRunStop) == 0 || err != nil {
		errOut := "没有获取到容器（Ruuning状态）停止按钮，请检查。"
		tools.ReadFile(errOut)
	}
	// 删除
	containerRunDelPath := "./textfile/dockerRunDel.txt"
	containerRunDel, err := tools.ReadFile(containerRunDelPath)
	if len(containerRunDel) == 0 || err != nil {
		errOut := "没有获取到容器（Ruuning状态）删除按钮，请检查。"
		tools.ReadFile(errOut)
	}
	//日志
	containerRunLogPath := "./textfile/dockerRunLog.txt"
	containerRunLog, err := tools.ReadFile(containerRunLogPath)
	if len(containerRunLog) == 0 || err != nil {
		errOut := "没有获取到容器（Ruuning状态）日志按钮，请检查。"
		tools.ReadFile(errOut)
	}

	c.HTML(http.StatusOK, "view/cloud/containerListRunning.html", gin.H{
		"title":            "在线容器",
		"containerRunName": containerRunName,
		"containerRunId":   containerRunId,
		"containerStatus":  containerStatus,
		"containerRunRes":  containerRunRes,
		"containerRunStop": containerRunStop,
		"containerRunDel":  containerRunDel,
		"containerRunLog":  containerRunLog,
	})
}

// 离线容器
func (dc CloudController) ContainerListExit(c *gin.Context) {
	containerExitNamePath := "./textfile/containerExitName.txt"
	containerExitName, err := tools.ReadFile(containerExitNamePath)
	if len(containerExitName) == 0 || err != nil {
		errOut := "没有获取到容器名，请登录到服务器检查。"
		tools.LogFileErr(errOut)
	}

	containerExitIdPath := "./textfile/dockerExitId.txt"
	containerExitId, err := tools.ReadFile(containerExitIdPath)
	if len(containerExitId) == 0 || err != nil {
		errOut := "没有获取到离线容器ID，请登录到服务器检查。"
		tools.LogFileErr(errOut)
	}

	// 遍历ContainerStatus
	containerExitStatusPath := "./textfile/dockerExitStatus.txt"
	containerExitStatus, err := tools.ReadFile(containerExitStatusPath)
	if len(containerExitStatus) == 0 || err != nil {
		errOut := "没有获取到离线容器运行状态信息，请登录到服务器检查。"
		tools.LogFileErr(errOut)
	}

	// 离线容器操作 重启、停止、删除、日志
	//重启
	containerExitResPath := "./textfile/dockerExitRes.txt"
	containerExitRes, err := tools.ReadFile(containerExitResPath)
	if len(containerExitRes) == 0 || err != nil {
		errOut := "没有获取到容器（Exit状态）重启按钮，请检查。"
		tools.ReadFile(errOut)
	}
	// 停止
	containerExitStopPath := "./textfile/dockerExitStop.txt"
	containerExitStop, err := tools.ReadFile(containerExitStopPath)
	if len(containerExitStop) == 0 || err != nil {
		errOut := "没有获取到容器（Exit状态）停止按钮，请检查。"
		tools.ReadFile(errOut)
	}
	// 删除
	containerExitDelPath := "./textfile/dockerExitDel.txt"
	containerExitDel, err := tools.ReadFile(containerExitDelPath)
	if len(containerExitDel) == 0 || err != nil {
		errOut := "没有获取到容器（Exit状态）删除按钮，请检查。"
		tools.ReadFile(errOut)
	}
	//日志
	containerExitLogPath := "./textfile/dockerExitLog.txt"
	containerExitLog, err := tools.ReadFile(containerExitLogPath)
	if len(containerExitLog) == 0 || err != nil {
		errOut := "没有获取到容器（Exit状态）日志按钮，请检查。"
		tools.ReadFile(errOut)
	}

	c.HTML(http.StatusOK, "view/cloud/containerListExit.html", gin.H{
		"title":               "离线容器",
		"containerExitName":   containerExitName,
		"containerExitId":     containerExitId,
		"containerExitStatus": containerExitStatus,
		"containerExitRes":    containerExitRes,
		"containerExitStop":   containerExitStop,
		"containerExitDel":    containerExitDel,
		"containerExitLog":    containerExitLog,
	})
}

// k8s核心组件
func (k8s CloudController) KubernetesService(c *gin.Context) {
	c.HTML(http.StatusOK, "view/cloud/kubernetesService.html", gin.H{
		"title": "Kubernete Service",
	})
}

// 业务Pod
func (k8s CloudController) PodService(c *gin.Context) {
	c.HTML(http.StatusOK, "view/cloud/podService.html", gin.H{
		"title": "Kubernete Service",
	})
}
