package main

import (
	"fmt"
	"time"
)

func main() {

	// 获取本地时间：time.Now()
	now := time.Now()
	fmt.Println(now)

	// 获取其中年月日时分秒
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	min := now.Minute()
	sec := now.Second()
	fmt.Println("year = ", year)
	fmt.Println("month = ", month)
	// 获取的month为英文月份，可以用int()将其转为数字月份
	fmt.Println("month = ", int(month))
	fmt.Println("day = ", day)
	fmt.Println("hour = ", hour)
	fmt.Println("min = ", min)
	fmt.Println("sec = ", sec) 

	// 格式化时间
	// 1. 手动格式化
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, min, sec)
	datestr := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, min, sec)
	fmt.Println(datestr)
	// time.Format("2006-01-02 15:04:05"): 格式化时间
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))

	// 休眠时间：time.Sleep(t), 其中t用time定义的时间常量表示
	for i := 0; i < 4; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}

	// 1970.1.1 00:00:00 到当前时间的秒数/纳秒数：time.Unix() / time.UnixNano()
	fmt.Println(now.Unix(), now.UnixNano())

}