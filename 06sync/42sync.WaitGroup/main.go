package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//WaitGroup实现goroutine的同步，以达到推出
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()      // goroutine结束就-1
	r1 := rand.Intn(100) //随机生成100以内的值
	fmt.Println(r1, i)
}
func main() {
	rand.Seed(time.Now().UnixNano()) //随机数种子为时间的纳秒
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
