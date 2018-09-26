package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	oneDayAfter, _ := time.ParseDuration("86400s")
	oneDayBefore, _ := time.ParseDuration("-86400s")
	oneHourBefore, _ := time.ParseDuration("1h")
	tomorrow := now.Add(oneDayAfter)
	// 格式化
	fmt.Println("now 格式化：", now.Format("2006-01-02 15:04:05"))
	// Unix
	nowUnix := time.Now().Unix()
	fmt.Println("microtime", time.Now().UnixNano()/int64(time.Millisecond))
	last6 := nowUnix - 6*3600
	fmt.Println("6*3600 秒前格式化：", time.Unix(last6, 0).Format("2006-01-02 15:04:05"))
	// 增加时间
	fmt.Println("oneHourBefore 后格式化：", now.Add(oneHourBefore).Format("2006-01-02 15:04:05"))
	fmt.Println("oneDayBefore 格式化：", now.Add(oneDayBefore).Format("2006-01-02 15:04:05"))
	// 星期几
	weekDay := now.Weekday()
	tomorrowWeek := tomorrow.Weekday()
	fmt.Println("今天星期几：", weekDay, weekDay.String())
	fmt.Println("明天星期几：", tomorrowWeek, tomorrowWeek == 0, tomorrowWeek.String())
}
