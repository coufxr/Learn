package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Printf("%d\n", <-c) //这里接受channel c，并打印
		}
		quit <- 1 //发送channel quit 的值1
	}() //闭包函数
	testnum(c, quit) //执行函数

	chs := [3]chan int{
		make(chan int, 3),
		make(chan int, 3),
		make(chan int, 3),
	}
	rand.Seed(time.Now().UnixNano())
	index1 := rand.Intn(3) // 随机生成0-2之间的数字
	fmt.Printf("随机索引/数值: %d\n", index1)
	chs[index1] <- rand.Int() // 向通道发送随机数字

	index2 := rand.Intn(3)
	fmt.Printf("随机索引/数值: %d\n", index2)
	chs[index2] <- rand.Int()

	index3 := rand.Intn(3)
	fmt.Printf("随机索引/数值: %d\n", index3)
	chs[index3] <- rand.Int()

	// 哪一个通道中有值，哪个对应的分支就会被执行
	for i := 0; i < 3; i++ {
		select {
		case num, ok := <-chs[0]:
			if !ok {
				break
			}
			fmt.Println("第一个条件分支被选中: chs[0]=>", num)
		case num, ok := <-chs[1]:
			if !ok {
				break
			}
			fmt.Println("第二个条件分支被选中: chs[1]=>", num)
		case num, ok := <-chs[2]:
			if !ok {
				break
			}
			fmt.Println("第三个条件分支被选中: chs[2]=>", num)
		default:
			fmt.Println("没有分支被选中")
		}
	}
}

//testnum 测试
func testnum(c, quit chan int) {
	x, y := 0, 1
	for {
		select { //监控？？
		case c <- x: //这里将x发送至channel c
			x, y = y, x+y
		case <-quit: //接受channel quit 为真，打印结束
			fmt.Print("quit")
			return

		}
	}
}
