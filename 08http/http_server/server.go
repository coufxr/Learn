package main

import (
	"io"
	"net/http"
)

//路由执行的函数
func f1(w http.ResponseWriter, r *http.Request) {
	//str:="<h1 style='color:red'>你好，golang！</h1>"
	//_, _ = w.Write([]byte(str))
	//官方推荐使用下面这个方式
	_, _ = io.WriteString(w, "<h1 style='color:red'>你好，golang！</h1>")
}

func main() {
	http.HandleFunc("/index/", f1)                 //路由
	_ = http.ListenAndServe("127.0.0.1:9090", nil) //指定链接ip地址
}
