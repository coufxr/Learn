package main

import "fmt"

//接口，速度多少
type seeper interface {
	seep() //接口下的方法，只有实现了这个方法的变量即为seeper类型
}

type cat struct{}
type dog struct{}
type person struct{}

func (c cat) seep() {
	fmt.Println("我的速度10m/s")
}
func (d dog) seep() {
	fmt.Println("我的速度15m/s")
}
func (p person) seep() {
	fmt.Println("我的速度8m/s")
}
func main() {
	var (
		c1 cat
		d1 dog
		p1 person
	)
	c1.seep()
	d1.seep()
	p1.seep()
}
