package utils

import (
	"strconv"
	"strings"
	"time"
)

var Layout = "2006-01-02 15:04:05"
var ZeroTime = "1970-01-01 00:00:00"

// 获取当前日期时间, 如: 2023-12-21 23:13:00
func GetNow() string {
	return time.Now().Format(Layout)
}

func GetZeroTime() string {
	return ZeroTime
}

func GetLatestTime(str1 string, str2 string) (string, error) {
	t1, err := time.Parse(Layout, str1)
	if err != nil {
		return "", nil
	}

	t2, err := time.Parse(Layout, str2)
	if err != nil {
		return "", nil
	}

	if t1.After(t2) {
		return t1.Format(Layout), nil
	} else {
		return t2.Format(Layout), nil
	}
}

// 获取日期数字,形如: 20230821
func GetDateNum(t time.Time) int {
	s := t.Format("2006-01-02")
	s = strings.ReplaceAll(s, "-", "")
	ret, _ := strconv.Atoi(s)
	return ret
}

// 将 "2023-08-29T16:32:42+08:00"
// 转为 "2023-08-29 16:32:42" 格式
func FormatTime(rfc3339Time string) string {
	t, _ := time.Parse(time.RFC3339, rfc3339Time)
	return t.Format(Layout)
}

// 将 "2007-12-03T00:00:00+08:00"
// 转为 "2007-12-03" 格式
func FormatDate(rfc3339Time string) string {
	t, _ := time.Parse(time.RFC3339, rfc3339Time)
	return strings.Split(t.Format(Layout), " ")[0]
}

// 将 "2023-09-04 14:46:00" 转为 20230904
func GetDateNumFromDatetime(datetime string) (int, error) {
	t, err := time.Parse(Layout, datetime)
	if err != nil {
		return 0, err
	}

	return GetDateNum(t), nil
}

// 将 "2023-09-04" 转为 20230904
func GetDateNumFromDate(s string) int {
	s = strings.ReplaceAll(s, "-", "")
	ret, _ := strconv.Atoi(s)
	return ret
}
