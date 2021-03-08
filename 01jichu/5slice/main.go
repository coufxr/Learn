package main

import "fmt"

func main() {
	var s1 []int //声明切片。。也就是数组
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //是否为空？？
	//fmt.Println(s1 == null)不存在null关键字
	s1 = []int{1, 2, 3, 4, 5} //初始化。。前方可不写长度
	fmt.Println(s1 == nil)    //是否为空？？
	fmt.Printf("len(s1):%d,cap(s1):%d\n", len(s1), cap(s1))
	//cap()原数组的容积，从切片数组的第一项开始算起到最后一项元素结束。

}
