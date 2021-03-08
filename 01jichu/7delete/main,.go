package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	//删除，go没有此方法，只能使用下面的方法
	s1 = append(s1[:1], s1[2:]...) //删除下标为1的元素
	fmt.Println(s1)
	fmt.Println(cap(s1)) //容量不会消失

}
