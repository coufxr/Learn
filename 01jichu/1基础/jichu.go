package main

import (
	"fmt"
)

var (
	names string = "小明"
)

const (
	pi = 3.1415926
	s1 = 100
	s2
	s3
)
const (
	n1 = iota
	n2
	n3
)
const (
	_  = iota
	kB = 1 << (10 * iota)
	mB = 1 << (10 * iota)
	gB = 1 << (10 * iota)
	tB = 1 << (10 * iota)
)

func main() {
	s := 456
	name := "小兰"
	println(name, names, s)
	fmt.Printf("%7.7f \n", pi*3) //仅限于fmt包？？
	println(s1, s2, s3)
	println(n1, n2, n3)
	println(kB, mB, gB, tB)
}
