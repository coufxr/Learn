package main

import "fmt"

func a() {
	fmt.Println("aaa")
}
func b() {
	defer func() { //这里需配合使用？？//recover必须配合defer使用，一定要在可能产生panic之前定义
		err := recover() //错误解决机制，即使出错也继续执行
		fmt.Println(err)
		fmt.Println("我正在连接。。。")
	}()
	panic("这里出错了") //类似try。。错误抛出机制
	//fmt.Println(" bbb ") //永远到达不了
}
func c() {
	fmt.Println("ccc")
}
func main() {
	a()
	b()
	c()
}
