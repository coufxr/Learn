package main

import (
	"fmt"
)

func main() {
	black := "\033[40;37m \033[0m "
	white := "\033[40;47m \033[0m "

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if (i+j)%2 == 0 {
				fmt.Printf(white)
			} else {
				fmt.Printf(black)
			}
		}
		fmt.Println()
	}
	fmt.Println("----------------------------")
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			if (i+j)%2 == 0 {
				fmt.Printf(white)
			} else {
				fmt.Printf(black)
			}
		}
		fmt.Println()
	}
}
