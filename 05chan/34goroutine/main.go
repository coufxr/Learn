package main

import (
	"fmt"
	"runtime"
)

func main() {
	go hello() //启动并发执行hello（）
	go sb()    //后进先出？？
	//goroutine返回将结束main
	var s string
	fmt.Scan(&s) //防止程序结束
	// fmt.Println(runtime.NumCPU())
	// fmt.Println(runtime.NumGoroutine())

}

func hello() {
	for i := 0; i < 10; i++ {
		fmt.Print("hello ")
		runtime.Gosched() //计数信号量，可以用来记录并维护运行的 goroutine
	}
}
func sb() {
	for i := 0; i < 10; i++ {
		fmt.Println("sb")
		runtime.Gosched()

	}
}
