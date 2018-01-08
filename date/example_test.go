package date_test

import (
	"fmt"
	"time"
)

//字符串转为time.Time类型
func ExampleTime() {
	t := time.Date(2018, 1, 1, 10, 0, 0, 0, time.Local)
	fmt.Println(t)
	// Output: 2018-01-01 10:00:00 +0800 CST
}

//根据字符串获取时间
func ExampleFormat() {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-01-02 10:06:00", time.Local)
	fmt.Println(t)
	// Output: 2018-01-02 10:06:00 +0800 CST
}

//获取当前时间
func ExampleNow() {
	now := time.Now()
	fmt.Println(now)
	strtime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(strtime)
}
