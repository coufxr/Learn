package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	// fmt.Println(t.Year())
	// fmt.Println(t.Month())
	// fmt.Println(t.Day())
	// fmt.Println(t.Date())//几月几号
	// fmt.Println(t.Hour())
	// fmt.Println(t.Minute())
	// fmt.Println(t.Second())
	//格式化输入：200612345
	n := t.Format("2006-01-02 15:04:05")
	fmt.Println(n)
	//延时
	// for {
	// 	time.Sleep(time.Duration(2) * time.Second)
	// 	fmt.Println("延时2秒")
	// }
	//定时器
	tt := time.NewTicker(10 * time.Second)
	fmt.Println(tt.C)
	// time.Tick(time.Duration(10) * time.Second)

	// fmt.Println("10秒后执行:", n)

	// ticker := time.NewTicker(time.Second * 3)
	// for t := range ticker.C {
	// 	fmt.Println("Tick at", t)
	// }
	// ticker.Stop()
	// fmt.Println("Ticker stopped")

}
