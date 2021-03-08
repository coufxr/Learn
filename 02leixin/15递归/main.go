package main

import "fmt"

var (
	i = 0
	b = 0
	x int
)

func sum(i int) int {
	b = b + i
	return b
}
func dg(ipt int) {
	i++
	// println(i)
	x = sum(i)
	if i >= ipt {
		println("stop")
		fmt.Println(x)
		return //????为啥是return而不是break
	}
	dg(ipt)
}
func main() {
	var ipt int
	fmt.Scanf("%d", &ipt)
	dg(ipt)
}
