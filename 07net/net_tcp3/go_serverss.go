package main

import (
	"Learn/07net/proto"
	"bufio"
	"fmt"
	"io"
	"net"
)

//进行封包和拆包可以使粘包得到解决
//封包：包头和包体；包头为固定长度，包含了包体的长度信息以及其他信息，就和请求头一样；包体：包含信息体，同请求体

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
