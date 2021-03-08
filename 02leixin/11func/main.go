package main

import "fmt"

func sum(x, y int) int { //参数类型可简写，同类型只保留最后一个声明
	return x + y
}
func f1(x, y, z string, i, j bool) { //比如像这样

}
func f2(x string, y ...int) { //y为可变长参数，即可传多个值（即[]切片）
	fmt.Println(x)
	i := y[3] //得到值的方法和切片获取一样
	fmt.Println(i)

}
func mul(x, y int) int {
	return x * y
}
func f3() int {
	fmt.Println("ssssgg")
	return 210
}

func f4(f func() int) { //函数作为参数时，传入的函数类型要满足参数函数，比如参数函数的参数，返回值类型以及各自的数量
	b := f()
	fmt.Println(b)
}

//满足此函数形式即可传参
func f5(f func(int, int) int, i, j int) { //此时的f()函数内部有参，我们传入的函数也要像这样有参
	x := f(i, j) //我们将参数传入至f()即mul()函数，mul将执行后的结果返回至此，并用x接受
	fmt.Println(x)
}

//同理函数也可做为返回值
func f6(x, y int) func(int, int) int {
	ret := func(int, int) int {
		b := x * y
		return b
	}
	return ret
}
func main() {
	// fmt.Println(sum(10, 18))
	// fmt.Println(mul(29, 13))
	// f2("上海", 1, 34, 56, 7, 57, 868)
	f4(f3) //传入函数时，不用写（）
	f5(mul, 4, 6)
	x := f6(4, 5)
	fmt.Printf("%T", x) //f6()的返回值类型为函数
}
