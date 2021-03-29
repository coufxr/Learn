package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//GET请求网站，返回response.包含了页面元素
	resp, err := http.Get("https://live.bilibili.com/363186")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	haed := resp.Header                //请求头
	proto := resp.Proto                //通讯协议
	major := resp.ProtoMajor           //？
	Minor := resp.ProtoMinor           //?
	reqs := resp.Request               //请求体
	status := resp.Status              //状态码
	statusCode := resp.StatusCode      //状态码
	tls := resp.TLS                    //?
	Uncompressed := resp.Uncompressed  //?
	body, err := io.ReadAll(resp.Body) //读取
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Print(string(body))
	fmt.Println(haed)
	fmt.Println(proto)
	fmt.Println(major)
	fmt.Println(Minor)
	fmt.Println(reqs)
	fmt.Println(status)
	fmt.Println(statusCode)
	fmt.Println(*tls)
	fmt.Println(Uncompressed)
}
