package main

import "fmt"

func main() {
	//这是个切片，切片中的元素是 map
	var s1 = make([]map[int]string, 10, 10) //无序的,(make()进行初始化)
	s1[0] = make(map[int]string, 10)        //声明下标为0的map类型
	s1[0][3] = "天地"                         //对0下标的map进行初始化。前[]为下标。后[]为key
	s1[0][4] = "天下"                         //即【4】为key
	s1[1] = make(map[int]string, 2)         //下标为1
	s1[1][1] = "市民"
	s1[1][2] = "教案"
	fmt.Printf("%T\n", s1)
	fmt.Println(s1)
	fmt.Println(s1[0]) //这里打印的是s1这个切片的第【0】个元素
	//值是切片的map
	var m1 = make(map[string][]string, 10)
	m1["北京"] = []string{"朝阳", "昌平", "通州", "密云"}
	fmt.Println(m1)
}
