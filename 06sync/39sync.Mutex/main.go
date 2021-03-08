package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//“sync”	同步包
//Mutex		互斥锁，保证channel的通讯安全和不会被污染
//注意事项：
//不要重复锁定互斥锁；
//不要忘记解锁互斥锁，必要时使用 defer 语句；
//不要对尚未锁定或者已解锁的互斥锁解锁；
//不要在多个函数之间直接传递互斥锁。

var counter int = 0

func add(a, b int, lock *sync.Mutex) {
	c := a + b
	lock.Lock() //表示上锁
	counter++
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
	lock.Unlock() //表示解锁
}

func main() {
	start := time.Now()
	lock := &sync.Mutex{} //声明一个Mutex 互斥锁
	for i := 0; i < 10; i++ {
		go add(1, i, lock) //使用协程进行计算，并传入了锁
	}

	for {
		lock.Lock() //防止读取阶段被其他并发操作更新
		c := counter
		lock.Unlock()
		runtime.Gosched() //将协程暂时切换出去
		//fmt.Println(c)
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
