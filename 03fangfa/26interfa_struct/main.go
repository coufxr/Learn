package main

import "fmt"

//接口嵌套
type animal interface {
	mover
	eater
}

//同一结构体实现多个接口
type mover interface {
	move()
}
type eater interface {
	eat(string)
}
type cat struct {
	name string
	feet int
}

func (c cat) move() {
	fmt.Println("走猫步")
}
func (c cat) eat(food string) {
	fmt.Printf("%s吃%s\n", c.name, food)
}
func main() {
	c2 := cat{"星猫", 4} //指针值
	fmt.Println(c2)
}
