package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "VS Code VS Code"
	fmt.Println(name)
	ret := strings.Split(name, " ")
	fmt.Println(ret)                          //分割字符串，分割后是数组格式？
	fmt.Println(strings.Contains(name, "VS")) //检测是否包含
	fmt.Println(strings.HasPrefix(name, "V")) //检测前缀
	fmt.Println(strings.HasSuffix(name, "e")) //检测后缀，上三者皆为布尔值
	fmt.Println(strings.Index(name, " "))     //下标位置，第一次出现的下标
	fmt.Println(strings.LastIndex(name, " ")) //下标位置，最后一次出现的下标
	fmt.Println(strings.Join(ret, "+"))       //拼接数组或者字符串？string类型无法直接使用
	fmt.Printf("%T\n", ret)                   //格式检查及其他打印操作只能这样使用

}
