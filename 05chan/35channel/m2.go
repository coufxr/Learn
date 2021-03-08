package main

import (
	"fmt"
)

func main() {
	buf := make(chan int)
	flg := make(chan int)
	//启动两个线程
	go gen(buf)
	go consumer(buf, flg)
	//channel 管道传输信息，带有缓冲区。缓冲区大小是指能接收的值的个数，与值的大小无关
	//单向通讯如下所示
	//<- chan int 表示仅能接受,
	//chan <- int 表示仅能发送
	//双向通讯可以转化成单项通讯，但单向通讯无法转化成双向通讯。
	//ch1 := make(chan int)
	//ch2 := <-chan int(ch1)	接收
	//ch3 := chan<- int(ch1)	发送
	<-flg //接收f发送的打印数据

	c := make(chan int, 5) //后为缓冲区大小，接受和发送的数量？？与字节大小无关？缓冲区大小是指能接收的值的个数，与值的大小无关
	c <- 13254235425423
	c <- 2
	// fmt.Println()这里接受c的值
	fmt.Println(<-c)
	fmt.Println(<-c)
}
func gen(c chan<- int) {
	defer close(c) //延时关闭线程c
	for i := 0; i < 10; i++ {
		c <- i //每循环一次就会通过channel c向外传输一次
	}
}
func consumer(c <-chan int, f chan<- int) {
	for {
		if v, ok := <-c; ok { //判断channel c中是否有值，存在就打印，否则跳出if
			fmt.Printf("v=%v；\n", v)
		} else {
			break
		}
	}
	f <- 1 //向外发送打印的值

}
