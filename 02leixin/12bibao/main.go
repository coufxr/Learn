package main

import "fmt"

//匿名函数：//主函数内无法声明带名字的函数
var f1 = func() {
	fmt.Println("这是个匿名函数")
}

//闭包
func f3(f func()) {
	fmt.Println("这是f3")
	f()
}
func f4(x, y int) {
	fmt.Println("这是f4")
	fmt.Println(x + y)
}

//应用:
func bibao(f func(int, int), m, n int) func() { //实则是将f4这个函数进行封装变形，使之成为能被f3识别的格式
	tmp := func() {
		f(m, n) //这里是将m，n两个值传入f4这个函数中去
	}
	return tmp //这里实际上返回的是f4()执行后的结果
	//上面的代码实际上是：
	//x(m, n)
}
func main() {
	//匿名函数
	f2 := func(x, y int) {
		fmt.Println(x + y)
	}
	f2(2, 3) //需要立即执行此函数才可以
	//但只需要执行1次的函数可以写成“闭包”格式，即：
	func(x, y int) {
		fmt.Println(x * y)
	}(2, 3) //就像这样
	//
	f3(bibao(f4, 100, 200))

}
