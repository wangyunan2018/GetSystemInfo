package tools

import (
	"strings"
	"time"
)

type TimerDate string

const (
	YYYYMMDD            = "yyyyMMdd"
	YYYYMMDDHHMMSS      = "yyyyMMddHHmmss"
	YYYY_MM_DD_HH_MM_SS = "yyyy-MM-dd HH:mm:ss"
)

// 把日期转换成字符串
func FormatDate(date time.Time, dateStyle TimerDate) string {
	layout := string(dateStyle)
	layout = strings.Replace(layout, "yyyy", "2006", 1)
	layout = strings.Replace(layout, "yy", "06", 1)
	layout = strings.Replace(layout, "MM", "01", 1)
	layout = strings.Replace(layout, "dd", "02", 1)
	layout = strings.Replace(layout, "HH", "15", 1)
	layout = strings.Replace(layout, "mm", "04", 1)
	layout = strings.Replace(layout, "ss", "05", 1)
	layout = strings.Replace(layout, "SSS", "000", -1)
	return date.Format(layout)
}

//	调用
// tools.FormatDate(time.Now(), tools.YYYYMMDD)
