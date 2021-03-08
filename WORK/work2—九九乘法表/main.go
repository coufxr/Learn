package main

import "fmt"

func main() {
	//正序
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%2d ", j, i, j*i)
		}
		fmt.Println("")
	}
	//反序
	for i := 9; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%2d ", j, i, j*i)
		}
		fmt.Println("")
	}
}
