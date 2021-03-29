package main

import (
	"fmt"
	"strconv"
)

type book struct {
	id        int
	name      string
	auther    string
	chubanshe string
	update    int
}

//printbook 打印当前的book
func (b book) printbook() {
	fmt.Printf("ID:%v 名称:%s  作者:%s  出版社：%s  字数：%d \n",
		b.id, b.name, b.auther, b.chubanshe, b.update)
}

//updata 通过字符串数组进行更新数据，可以点对点修改
func (b *book) updata(args []string) {
	for i := 0; i < len(args); i++ {
		msg := args[i] //取得当前下标的值
		if msg != "" {
			switch i { //使用swicth进行判定
			case 0:
				tmp, err := strconv.Atoi(msg)
				if err != nil {
					fmt.Println("id输入错误,id类型需为数字")
				} else {
					b.id = tmp //将字符串转为数字
				}
			case 1:
				b.name = msg
			case 2:
				b.auther = msg
			case 3:
				b.chubanshe = msg
			case 4:
				tmp, err := strconv.Atoi(msg)
				if err != nil {
					fmt.Println("update输入错误,update类型需为数字")
				} else {
					b.update = tmp //将字符串转为数字
				}
			}
		}
	}
}
func (b *book) newBook(id int, name string, auther string, chubanshe string, update int) {
	b.id = id
	b.name = name
	b.auther = auther
	b.chubanshe = chubanshe
	b.update = update
}

func main() {
	var book3 book
	book3.newBook(24, "《java 从入门到入土》", "java", "北京邮电出版社", 100000)
	book3.printbook()
	book3.updata([]string{"", "《java 从入门到入坟》"})
	book3.updata([]string{"dasf", "", "", "", "fasfas"})
	book3.printbook()
	//var book book

}
