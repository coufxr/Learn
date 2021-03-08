package main

import "fmt"

//标识符大写表示对外部可见即公有类型

//Dog 一个结构体
type Dog struct {
	name string
}

func newDog(name string) Dog {
	return Dog{
		name: name,
	}
}

//方法，用于特定类型的函数，无法为其他包写方法，只能给自己的包添加
//接收者表示的是调用该方法的具体类型变量，多用类型名的首字母小写表示
//和this，self相似
func (d Dog) wang() {
	fmt.Printf("%s:汪汪汪~", d.name)
}

type person struct {
	name string
	age  int
}

func newperson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

//使用值接收者，拷贝值
func (p person) guonian() {
	p.age++
}

//使用指针才可以修改
func (p *person) zhenguonian() {
	p.age++
}
func main() {
	// newDog("赵宇").wang()
	pp := newperson("赵玉", 18)
	fmt.Println(pp.age)
	pp.guonian()
	fmt.Println(pp.age)
	pp.zhenguonian()
	fmt.Println(pp.age)
}
