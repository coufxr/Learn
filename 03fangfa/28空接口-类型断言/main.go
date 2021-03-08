package main

import "fmt"

//类型断言
//a.(type)

//例子1
func f1(a interface{}) {
	fmt.Printf("%T\n", a)
	str := a.(string)
	fmt.Println("这是字符串：", str, "!")
}

//例子2
//可以通过判断空接口接受的类型来进行下一步的操作
func f2(a interface{}) {
	switch t := a.(type) {
	case string:
		fmt.Println("这是字符串：", t)
	case int:
		fmt.Println("这是int：", t)
	case bool:
		fmt.Println("这是bool：", t)
	case byte:
		fmt.Println("这是byte：", t)

	}
}

func main() {
	// f1("十九世纪")
	f2(1)
}
