package main

import "fmt"

type person struct {
	name string
	age  int
}

//构造函数
//需要判断返回的是结构体还是结构体指针
//当结构体中的结构值比较多的时候使用结构体指针，以此来减少内存的开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
func main() {
	p1 := newPerson("赵飞", 45)
	p2 := newPerson("刘超", 38)
	fmt.Println(p1, p2)
}
