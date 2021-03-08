package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fileObj, err := os.Open("E:/YD/0729/fw1_rule_udp.log")
	if err != nil {
		fmt.Printf("open file failde,err:%v\n", err)
		return
	}
	// 延时关闭
	defer fileObj.Close()
	// bufio读取
	reader := bufio.NewReader(fileObj)
	for true {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("reader file failde,err:%v\n", err)
			return
		}
		strobj := strings.ReplaceAll(line, "=", " ")
		strobj2 := strings.Split(strobj, " ")
		fmt.Println(strobj2[1])
		//fmt.Println(str)
	}
}
