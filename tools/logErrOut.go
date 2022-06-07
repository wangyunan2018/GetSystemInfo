package tools

import (
	"fmt"
	"io"
	"os"
	"time"
)

func LogFileErr(logContent string) {
	errInfoFile := logContent
	path := "./log/"
	logName := "_err.log"
	fileName := path + FormatDate(time.Now(), YYYYMMDD) + logName

	var f *os.File
	var err error

	// 如果文件存在，则打开，反之，创建文件
	if checkIsFileExist((fileName)) {
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	} else {
		f, err = os.Create(fileName)
	}
	check(err)

	logText, err := io.WriteString(f, FormatDate(time.Now(), YYYY_MM_DD_HH_MM_SS)+":  "+errInfoFile+"\n")
	check(err)
	fmt.Println(logText)
}

func checkIsFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
