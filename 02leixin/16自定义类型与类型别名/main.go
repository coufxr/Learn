package main

import "fmt"

type myInt int    //自定义类型
type youInt = int //类型别名，只在编写中有作用，编译后无效（显示原类型）
func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n) //此时的n的类型为main.myInt类型

	var m youInt
	m = 10
	fmt.Println(m)
	fmt.Printf("%T\n", m) //m的类型为int

	var r rune //是int32的别名，指示为中文字符
	r = '中'
	fmt.Println(r)
	fmt.Printf("%T\n", r)
}
