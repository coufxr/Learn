package main

import (
	"fmt"
	"sync/atomic"
)

//协程是通过 CPU 的调度不断从运行状态切换到非运行状态，或者从非运行状态切换到运行状态
//中断：我们把代码从运行状态切换到非运行状态称之为中断。
//中断的时机很多，比如任何两条语句执行的间隙，甚至在某条语句执行的过程中都是可以的，即使这些语句在临界区内也是如此
//原子操作
//原子操作通常是 CPU 和操作系统提供支持的，由于执行过程中不会中断，所以可以完全消除竞态条件，从而绝对保证并发安全性
//原子操作包括加法（Add）、比较并交换（Compare And Swap，简称 CAS）、加载（Load）、存储（Store）和交换（Swap）

func main() {
	//加减法	第一个参数是操作数对应的指针，第二个参数是加/减值
	var num int32
	_, _ = fmt.Scanf("%d", &num)
	var i int32 = 1
	atomic.AddInt32(&i, num)
	fmt.Printf("%d + %d =%d", i, num, i)
}
