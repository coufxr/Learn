package main

import (
	"bufio"
	"fmt"
	"os"
)

func filewrite() {
	//打开文件,和open有区别
	fileObj, err := os.OpenFile("./a.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("open file failde,err:", err)
		return
	}
	defer fileObj.Close()
	//写入，通过字节写入
	fileObj.Write([]byte("海上生明月\n")) //需要类型强转
	//直接写入字符串
	fileObj.WriteString("天涯共此时\n") //没有换行，需要直接添加
	//??
	//fileObj.WriteAt([]byte("作者"), 65)
}
func bufioWrite() {
	fileObj, err := os.OpenFile("./a.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failde,err:", err)
		return
	}
	defer fileObj.Close()
	//创建一个bufio的写入对象
	w := bufio.NewWriter(fileObj)
	w.WriteString("作者是我") //此时将写入缓存中
	w.Flush()             //此时写入文件
}
func writefile() {
	fileObj, err := os.OpenFile("./a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failde,err:", err)
		return
	}
	defer fileObj.Close()
	count2, _ := fileObj.Seek(5, 0) // 计算文件起始位置和末尾之间的字符数（文件大小）。
	_, _ = fileObj.WriteAt([]byte("hello"), count2)
}
func main() {
	// filewrite()
	// bufioWrite()
	writefile()
}
