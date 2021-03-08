package main

import "fmt"

func main() {
	n1 := [...]int{1, 3, 5, 7, 8}

	sum := 0
	for i := 0; i < len(n1); i++ {
		sum += n1[i]
	}
	// println(sum)
	fmt.Println(sum, n1)
	n1[3] = 6 //可以重新进行赋值
	for i := 0; i < len(n1); i++ {
		println(n1[i])
	}
}
