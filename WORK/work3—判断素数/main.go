package main

import (
	"fmt"
)

var (
	a, y int
)

func fog(a int) {
	for i := 1; i <= a; i++ {
		y = a % i
		if y == 0 {
			fmt.Println("因素有：", i)
		} else {
			fmt.Println("此数是素数")
			break
		}
	}
}
func main() {
	fmt.Scanf("%d", &a)
	fog(a) //取余判断是否为素数
}
