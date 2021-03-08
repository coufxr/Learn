package main

import "fmt"

//面向对象，构建一个类设置对象的特征
type person struct {
	name  string
	age   int
	sex   string
	likes []string
}

func main() {
	var p person
	//使用.语法进行赋值
	p.name = "草渐青"
	p.age = 18
	p.sex = "女"
	p.likes = []string{"剑", "师兄"}
	fmt.Println(p)
	fmt.Println(p.age) //通过key获取值
	//匿名结构体，可在函数内进行声明构造，只使用一次
	var pp struct {
		name string
		age  int
	}
	pp.name = "哈哈哈"
	pp.age = 46
	fmt.Println(pp)
}
