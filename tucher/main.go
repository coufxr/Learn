package main

import "fmt"

type do interface {
	jiao()
	ranse()
}
type animal interface {
}
type cat struct {
	name  string
	age   int
	color color
}
type dog struct {
	name  string
	age   int
	color color
}
type color struct {
	R int
	G int
	B int
}

func (c cat) jiao() {
	fmt.Printf("%s 会喵喵喵的叫 \n", c.name)
}
func (d dog) jiao() {
	fmt.Printf("%s 会嗷呜嗷呜的叫 \n", d.name)
}

//interface 接口在使用时需要加上{},前可跟...,[]等约束类
func (c color) ranse(i interface{}) {
	fmt.Printf("%s 是#%d%d%d颜色的 \n", i, c.R, c.G, c.B)
}

func main() {
	var d dog
	d.name = "erha"
	d.age = 3
	d.color = color{52, 46, 37}
	//可简写成以下形式：
	var aa = cat{"hhh", 2, color{34, 43, 92}}
	c := cat{"bleak C", 2, color{12, 56, 18}}

	c.jiao()
	d.jiao()
	aa.jiao()
	c.color.ranse(c.name)
	d.color.ranse(d.name)

}
