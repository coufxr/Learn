package main

import (
	"fmt"
	"io"
	"os"
)

// os.O_RDONLY // 只读
// 	  os.O_WRONLY // 只写
// 	  os.O_RDWR // 读写
// 	  os.O_APPEND // 追加（Append）
// 	  os.O_CREATE // 如果文件不存在则先创建
// 	  os.O_TRUNC // 文件打开时裁剪文件
// 	  os.O_EXCL // 和O_CREATE一起使用，文件不能存在
// 	  os.O_SYNC // 以同步I/O的方式打开
// 	第三个参数：权限(rwx:0-7)
// 	  0：没有任何权限
// 	  1：执行权限
// 	  2：写权限
// 	  3：写权限和执行权限
// 	  4：读权限
// 	  5：读权限和执行权限
// 	  6：读权限和写权限
// 	  7：读权限，写权限，执行权限
// 写入文件 从指定光标位置写入  WriteAt()
// count2, _ := fp.Seek(0, io.SeekEnd)  // 计算文件起始位置和末尾之间的字符数（文件大小）。
// count, _ = fp.WriteAt([]byte("hello"), count2)  // 从指定光标位置开始写入
func f() {
	//源文件
	fileObj, err := os.OpenFile("./a.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed err:", err)
		return
	}
	//临时文件
	ls, err := os.OpenFile("./l.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed err:", err)
		return
	}
	//读取原文件
	var ret [1]byte
	f, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Println("read file failed err:", err)
		return
	}
	//写入临时文件
	ls.Write(ret[:f])
	//写入插入值
	var n []byte
	n = []byte{'c'}
	ls.Write(n)
	//读取后面并写入临时文件
	var x [1024]byte
	for {
		a, err := fileObj.Read(x[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("reader file failde,err:%v\n", err)
			return
		}
		ls.Write(x[:a])
	}
	//源文件后续也写入临时文件
	fileObj.Close()
	ls.Close()
	os.Rename("./l.tmp", "./a.txt")
}
func main() {
	f()
}
