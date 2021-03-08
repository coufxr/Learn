package main

import (
	"fmt"

	"music.com/mats"
)

//init类似于python中的__name__=__main__
//init会在main之前执行，如果是导入包则在导入时就执行
//不能主动调用
//层级：pkg-init()>全局变量>init()>main()
func init() {

}
func main() {
	ret := mats.Add(10, 20)
	fmt.Println(ret)
}
