package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 缓冲区读取，可一行一行的读取
func filebufio() {
	//打开
	strObj, err := os.Open("src/mats/add.go")
	if err != nil {
		fmt.Printf("open file failde,err:%v\n", err)
		return
	}
	// 延时关闭
	defer strObj.Close()
	// bufio读取
	reader := bufio.NewReader(strObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("reader file failde,err:%v\n", err)
			return
		}
		fmt.Print(line)
	}

}

//字节读取，
func fileio() {
	//打开
	strObj, err := os.Open("src/mats/add.go")
	if err != nil {
		fmt.Printf("open file failde,err:%v\n", err)
		return
	}
	//延时关闭
	defer strObj.Close()
	//读取
	var tmp [1024]byte
	for {
		n, err := strObj.Read(tmp[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file failde,err:%v\n", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}

//直接读取整个文件，实质是利用bufio进行的操作
func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("src/mats/add.go")
	if err != nil {
		fmt.Printf("read file failde,err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

//file操作
func main() {
	// filebufio()
	fileio()
	readFromFileByIoutil()
	//readFromFile_Ioutil()
}
