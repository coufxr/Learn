package main

import "fmt"

//匿名字段，只有类型，实际是以类型为名
//字段较少并简单的场景，不常用
type person struct {
	string
	int
}
type dog struct {
	jiao string
	sex  string
}
type song struct {
	name string
	age  int
	dog  //匿名字段嵌套
	//如果存在两个字段嵌套会冲突
}

func main() {
	p := person{
		"赵玉",
		5000,
	}
	fmt.Println(p)
	fmt.Println(p.string) //类型即名
	s := song{
		"憨憨",
		15,
		dog{
			"汪汪汪",
			"男",
		},
	}
	fmt.Println(s)
	fmt.Println(s.jiao) //可直接找到，
}
