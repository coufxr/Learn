package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//RWMutex	读写锁，支持单写多读。目的在于提升性能
//	lock	写锁
//	Rlock 	读锁
//	Rlock()、lock()都需要相对应的Runlock()、unlock()来进行解锁
//	要保证lock有相对应的解锁，否则会导致goroutine陷入阻塞状态，最终陷入死锁状态

var counter int = 0

func add(a, b int, lock *sync.RWMutex) {
	c := a + b
	lock.Lock()
	counter++
	fmt.Printf("%d: %d + %d = %d ----\n", counter, a, b, c)
	lock.Unlock()
}

func main() {
	start := time.Now()
	lock := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		go add(1, i, lock)
	}

	for {
		lock.RLock()
		c := counter
		lock.RUnlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
