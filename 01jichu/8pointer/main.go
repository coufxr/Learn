package main

import "fmt"

func main() {
	a := 10
	i := &a //取a的内存地址
	s := *i //指针根据内存地址进行寻址查找
	fmt.Println(a, i, s)
	//
	var a1 *int   //声明一个为nil（null）的指针
	a1 = new(int) //通过new关键字进行申请，申请一个内存地址
	*a1 = 100     //对指针进行赋值
	fmt.Println(a1)

	//new和make的区别
	//1.二者都为分配内存地址而使用
	//2.new多用于类型的地址分配，内存对应值多为类型零值，返回的是指向类型的指针
	//3.make只用于slice（切片），map，channe的初始化，返回的是这三者本身
}
