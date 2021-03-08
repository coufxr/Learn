package main

import (
	"fmt"
	"sync"
	"time"
)

//sync.Cond	条件变量	在对应的共享资源状态发生变化时，通知其它因此而阻塞的线程
//长与互斥锁组合使用，互斥锁为共享资源的访问提供互斥支持，而条件变量可以就共享资源的状态变化向相关线程发出通知，重在「协调」
//sync.Cond 存在3个方法：
//Wait()	等待
//Signal()	单发
//Broadcast()	广播
//调用 Wait() 函数前，需要先获得条件变量的成员锁，原因是需要互斥地变更条件变量的等待队列。在 Wait() 返回前，会重新上锁。
//条件变量和锁结合使用，在并发时如果逻辑不严谨容易发生死锁，所以尽量不要使用条件变量，推荐用 sync.WaitGroup 来实现并发时 Go 程间的同步。

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func main() {
	for i := 0; i < 10; i++ {
		go func(x int) {
			cond.L.Lock()         //获取锁
			defer cond.L.Unlock() //释放锁
			cond.Wait()           //等待通知，阻塞当前 goroutine

			// do something. 这里仅打印
			fmt.Println(x)
		}(i)
	}
	time.Sleep(time.Second * 1) // 睡眠 1 秒，等待所有 goroutine 进入 Wait 阻塞状态
	fmt.Println("Signal...")
	cond.Signal()

	time.Sleep(time.Second * 1) // 1 秒后下发下一个通知给已经获取锁的 goroutine
	fmt.Println("Signal...")
	cond.Signal()

	time.Sleep(time.Second * 1) // 1 秒后下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")
	cond.Broadcast()
	time.Sleep(time.Second * 1) // 睡眠 1 秒，等待所有 goroutine 执行完毕
}
