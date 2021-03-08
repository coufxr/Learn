package main

import "fmt"

type person struct {
	name string
}

func (p person) jiao() {
	fmt.Printf("%s:嗷呜~\n", p.name)
}

type zhoayu struct {
	person
	age int
}

func (z zhoayu) xuexi() {
	fmt.Printf("%s会唱歌\n", z.name)
}
func main() {
	p := zhoayu{
		person: person{
			name: "赵玉",
		},
		age: 18,
	}
	fmt.Println(p)
	p.jiao()
	p.xuexi()
}
