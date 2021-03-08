package main

import "fmt"

//值接收者和指针接收者

type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int
}

//使用值接受实现方法。能实现所有值
// func (c cat) move() {
// 	fmt.Println("走猫步")
// }
// func (c cat) eat(food string) {
// 	fmt.Printf("%s吃%s\n", c.name, food)
// }
//使用指针接受实现方法。只能实现指针
func (c *cat) move() {
	fmt.Println("走猫步")
}
func (c *cat) eat(food string) {
	fmt.Printf("%s吃%s\n", c.name, food)
}
func main() {
	var a animal
	// c1 := cat{"哈嘿", 4}  //值
	c2 := &cat{"星猫", 4} //指针值
	// a = c1              //c1是个指针值，无法被animal识别
	fmt.Println(a)
	a = c2
	fmt.Println(a)
}
