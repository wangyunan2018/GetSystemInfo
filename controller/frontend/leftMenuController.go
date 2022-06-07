package frontend

import (
	"fmt"
	"linux/GetSystemInfo/cmd"
	"linux/GetSystemInfo/tools"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// 左侧菜单栏
type LeftMenuController struct {
}

type GetCpuInfo struct {
	CpuCores int
}

type GetDiskInfo struct {
	DiskTotal      int
	DiskPercentage int
	DiskFree       int
}

type GetMemInfo struct {
	MemTotal   int
	MemPercent int
	MemFree    int
}

// CPU利用率
func (menu LeftMenuController) CpuInfo(c *gin.Context) {
	// 获取CPU信息
	cores, _ := cpu.Counts(false)
	cpus, _ := cpu.Percent(time.Duration(200)*time.Millisecond, true)
	var cpuu []float64
	cpuu = append(cpuu, cpus...)

	// 获取CPU温度  面向Linux物理服务器
	cmd1 := "sudo apt install lm-sensors -y || yum install -y lm_sensors"
	_, _, err := cmd.ShellOut(cmd1)
	if err != nil {
		log.Printf("命令执行失败，请检查输入是否有误。%v", err)
	}

	cmd2 := "sensors | grep 'Core' > textfile/cpuTemp.txt"
	_, errout2, err := cmd.ShellOut(cmd2)
	if err != nil {
		log.Printf("命令执行失败，请检查输入是否有误。%v", err)
	}

	fileName := "textfile/cpuTemp.txt"

	newSlic, _ := tools.ReadFile(fileName)
	if len(newSlic) == 0 {
		logErr := "没有获取到CPU温度数据，请检查服务器是否存在sensors命令。"
		fmt.Println(logErr)
		tools.LogFileErr(logErr)
	}

	fmt.Println(errout2)

	// 实例化GetCpuInfo{}结构体
	var getCpu GetCpuInfo = GetCpuInfo{
		CpuCores: cores,
	}

	c.HTML(http.StatusOK, "view/ajaxFresh/cpuInfo.html", gin.H{
		"title":    "CPU信息",
		"cpuCores": getCpu.CpuCores,
		"cpuFree":  cpuu,
		"cpuTemp":  newSlic,
	})
}

// 磁盘利用率
func (menu LeftMenuController) DiskInfo(c *gin.Context) {
	// 获取磁盘信息
	d, _ := disk.Usage("/")
	totalMB := int(d.Total) / MB
	usedPercent := int(d.UsedPercent)
	freeMB := int(d.Free) / MB

	// 实例化结构体
	var getDisk GetDiskInfo = GetDiskInfo{
		DiskTotal:      totalMB,
		DiskPercentage: usedPercent,
		DiskFree:       freeMB,
	}
	c.HTML(http.StatusOK, "view/ajaxFresh/diskInfo.html", gin.H{
		"title":          "磁盘信息",
		"diskTotal":      getDisk.DiskTotal,
		"diskPercentage": getDisk.DiskPercentage,
		"diskFree":       getDisk.DiskFree,
	})
}

// 内存使用率
func (menu LeftMenuController) MemInfo(c *gin.Context) {
	// 获取内存信息
	u, _ := mem.VirtualMemory()
	memTotal := int(u.Total) / MB
	memPercent := int(u.UsedPercent)
	memUsed := int(u.Used) / MB
	memFree := memTotal - memUsed

	// 实例化结构体
	var getMem GetMemInfo = GetMemInfo{
		MemTotal:   memTotal,
		MemPercent: memPercent,
		MemFree:    memFree,
	}

	c.HTML(http.StatusOK, "view/ajaxFresh/memInfo.html", gin.H{
		"title":      "内存信息",
		"memTotal":   getMem.MemTotal,
		"memPercent": getMem.MemPercent,
		"memFree":    getMem.MemFree,
	})
}

// GPU利用率
func (menu LeftMenuController) GpuInfo(c *gin.Context) {
	cmdGpu := "sudo chmod +x ./gpuExec.sh && /bin/bash ./gpuExec.sh"
	_, errout2, err := cmd.ShellOut(cmdGpu)
	if err != nil {
		fmt.Println("命令执行失败，请检查！")
	}

	// 获取GPU驱动版本
	logErr1 := "没有获取到GPU驱动版本"
	gpuVersion := "textfile/gpuVersion.txt"
	gpuV, _ := tools.ReadFile(gpuVersion)
	if len(gpuV) == 0 {
		fmt.Println(logErr1)
		tools.LogFileErr(logErr1)
	}

	// 获取GPU风扇温度|功率|显存利用率
	gpuPower := "textfile/gpuPower.txt"
	logErr2 := "没有获取到GPU风扇温度|功率|显存利用率"
	gpuP, _ := tools.ReadFile(gpuPower)
	if len(gpuP) == 0 {
		fmt.Println(logErr2)
		tools.LogFileErr(logErr2)
	}

	// 获取GPU进程
	gpuProcesses := "textfile/gpuProcesses.txt"
	gpuProc, _ := tools.ReadFile(gpuProcesses)
	if len(gpuProc) == 0 {
		logErr3 := "没有获取到占用GPU的进程信息"
		fmt.Println(logErr3)
		tools.LogFileErr(logErr3)
	}

	fmt.Println(errout2)

	c.HTML(http.StatusOK, "view/ajaxFresh/gpuInfo.html", gin.H{
		"title":        "GPU信息",
		"gpuVersion":   gpuV,
		"gpuPower":     gpuP,
		"gpuProcesses": gpuProc,
	})
}

// 容器信息
func (container LeftMenuController) ContainerInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "view/ajaxFresh/container.html", gin.H{
		"title": "容器信息",
	})
}

// k8s信息
func (k8s LeftMenuController) KubernetsInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "view/ajaxFresh/kubernetes.html", gin.H{
		"title": "k8s信息",
	})
}

// 日志信息
func (menu LeftMenuController) LogInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "view/ajaxFresh/logInfo.html", gin.H{
		"title": "日志信息",
	})
}
