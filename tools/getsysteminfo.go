package tools

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// 这个文件可以单独执行，现在暂时没用了，已经把这里面的逻辑通过gin框架细分渲染。
// 单独执行只需把导入包修改成 package main, 把GetSystemInfo函数修改成main函数

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func GetSystemInfo() {
	fmt.Println("######服务器硬盘使用量######")
	DiskCheck()
	fmt.Println("######服务器CPU使用量######")
	CPUCheck()
	fmt.Println("######服务器内存使用量######")
	RAMCheck()
}

//服务器硬盘使用量
func DiskCheck() {
	u, _ := disk.Usage("/")
	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)
	fmt.Printf("磁盘总容量：%dMB, 约：%dGB\n", totalMB, totalGB)
	fmt.Printf("磁盘已使用：%d%%\n", usedPercent)
	fmt.Printf("可用磁盘容量：%dMB, 约：%dGB\n", totalMB-usedMB, totalGB-usedGB)
	fmt.Println("")
}

//CPU 使用量
func CPUCheck() {
	cores, _ := cpu.Counts(false)

	cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true)
	if err == nil {
		for i, c := range cpus {
			fmt.Printf("cpu%d : %f%%\n", i, c)
		}
	}
	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15
	fmt.Println("L1负载：", l1)
	fmt.Println("L5负载：", l5)
	fmt.Println("L15负载：", l15)
	fmt.Println("核心数：", cores)
	fmt.Println("")
}

//内存使用量
func RAMCheck() {
	u, _ := mem.VirtualMemory()
	usedMB := int(u.Used) / MB
	totalMB := int(u.Total) / MB
	usedPercent := int(u.UsedPercent)
	fmt.Printf("内存总量: %dMB\n已使用: %dMB\n使用占比: %d%%\n", totalMB, usedMB, usedPercent)
	fmt.Println("剩余可用内存: ", totalMB-usedMB, "MB")
}
