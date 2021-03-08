package main

import "fmt"

type book struct {
	id        int
	name      string
	auther    string
	chubanshe string
	update    int
}

//type bookstr []map[...] string
func bookprint(book1 book) {
	fmt.Printf("名称:%s  作者:%s  出版社：%s  字数：%d \n",
		book1.name, book1.auther, book1.chubanshe, book1.update)
}
func (b *book) updata(new int) {
	b.update = new
}

func main() {
	var book1 book
	book1.id = 1
	book1.name = "《MySQL从删库到跑路》"
	book1.auther = "xxx"
	book1.chubanshe = "铁道出版社"
	book1.update = 50 * 10000

	var book2 book
	book2.id = 2
	book2.name = "《java从入门到入土》"
	book2.auther = "xxx"
	book2.chubanshe = "北京大学出版社"
	book2.update = 80 * 10000

	bookprint(book1)
	bookprint(book2)
	book2.updata(60 * 10000)
	bookprint(book2)

}
