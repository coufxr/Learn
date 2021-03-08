package main

import "fmt"

//空接口
//空接口可以作为函数参和类型参
func f(a interface{}) {
	fmt.Printf("%T\n:%v", a, a)
}

func main() {
	var m map[string]interface{}
	m = make(map[string]interface{}, 16)
	m["上海"] = "hhhhh"
	m["北京"] = 1534
	m["广州"] = true
	// fmt.Println(m)
	f(m)
}
