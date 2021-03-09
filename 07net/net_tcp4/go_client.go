package main

import (
	"Learn/07net/proto"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入：")
	input, _ := inputReader.ReadString('\n') // 读取用户输入
	inputInfo := strings.Trim(input, "\r\n")
	for i := 0; i < 20; i++ {
		//msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(inputInfo)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		_, _ = conn.Write(data)
	}
}
