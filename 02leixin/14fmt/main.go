package main

import "fmt"

func main() {
	// // 输出：
	// fmt.Print("aa")   //不会换行
	// fmt.Println("aa") //追加换行
	// aa := "sssss"
	// fmt.Printf("%s", aa) //不换行，后可追加\n进行换行

	//格式化输出：
	//%T:查看类型
	//%s:字符串
	//%d:十进制数
	//%b:二进制数
	//%o:八进制数
	//%x:十六进制数 小写a--f
	//$X:十六进制数 大写A--F
	//%c:字符//unicode对应的值
	//%U:unicode形式
	//%p:指针
	//%v:值
	//%f:浮点数
	//%x.yf:x为宽度，y为精度，x不写为默认，y不写为0
	//%t:布尔值
	//%%:百分号
	//%#v:go的语言显示(类型显示)
	//%q: 数字转译为ASCLL码，字符串则为用" "包裹起来
	//
	//输入：
	//总的来说，跟c一样输入没有前置提示词
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("输入的内容是：", s)
	var (
		name string
		age  int
		AD   string
	)
	// fmt.Scanf("%s %d %s", &name, &age, &AD) //限定输入类型
	// fmt.Println(name, age, AD)
	fmt.Println("请输入姓名，年龄，AD：")  //提示输入还是只能像这样
	fmt.Scanln(&name, &age, &AD) //只是多了一个换行判定
	fmt.Println(name, age, AD)

}
