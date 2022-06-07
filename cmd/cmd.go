package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
)

const ShellToUse = "/bin/bash"

// 执行Linux命令函数
func ShellOut(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

// 标准输出
func PrintOut() {
	fmt.Println("---- stdout ----")
}

// 标准错误输出
func PrintErr() {
	fmt.Println("---- stderr ----")
}
