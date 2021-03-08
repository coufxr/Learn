package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1
	fmt.Println(s1, s2)
	s1[1] = 100
	fmt.Println(s1, s2)
	for i, v := range s1 {
		fmt.Println(i, v)
	}
	// for i range (s1){
	// 	fmt.Println("sdd")//python写法
	// }
	s3 := []string{"广州", "上海", "苏州", "北京"}
	s4 := []string{"南京", "重庆", "成都", "西安"}
	fmt.Println(s3)
	//append必须使用变量接受
	s3 = append(s3, s4...) //append当数组未溢出时向后添加，如果溢出则将当前数组重新分配内存地址
	//s3 = append(s3, s1...)//只能同类型添加
	fmt.Println(s3)
	var c1 = make([]string, 10, 10) //make需要声明并赋给变量
	copy(c1, s3)                    //深拷贝？？
	fmt.Printf("%s,%s\n", s3, c1)

}
