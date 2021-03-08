package main

import "fmt"

var (
	n uint64
)

//一个台阶，每次只能走一步或者两步，问一共有多少种走法
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}
func main() {
	fmt.Println("请输入台阶数量：")
	fmt.Scanf("%d", &n)
	ret := taijie(n)
	fmt.Println(ret)
}
