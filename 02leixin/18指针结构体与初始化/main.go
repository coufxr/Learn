package main

import "fmt"

type person struct {
	name string
	age  int
}

func f1(x person) { //此时传入的值是拷贝值，并没有修改原值
	x.age = 45
}

//指针变量，用于修改已存在的成员量
func f2(x *person) {
	// (*x).age = 45 //()包裹用于强调为一个整体
	x.age = 45 //与上行同作用，语法糖，
}
func main() {
	var p person
	p.name = "中华"
	p.age = 35
	f1(p)
	fmt.Println(p)
	f2(&p) //调用是传参也需是指针的地址值
	fmt.Println(p)

	//结构体初始化1
	//可以使用key-value进行赋值
	var s = person{ //可直接写作&person来得到内存值
		name: "花花",
		age:  89, //最后一项需要加,
	}
	fmt.Println(s)
	//结构体初始化2
	//可直接按照顺序进行赋值
	var s2 = person{
		"天天",
		78,
	}
	fmt.Println(s2)
}
