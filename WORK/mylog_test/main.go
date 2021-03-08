package main

import "nguc.com/WORK/mylog"

func main() {
	// loger := mylog.NewLog("debug")
	loger := mylog.NewFilelogger("info", "./", "A.log", 10*1024^3)
	for {
		a := 2834
		b := "dsafa"
		loger.Debug("这是一条debug日志")
		loger.Trace("这是一条Trace日志,%v:%v", a, b)
		loger.Info("这是一条Info日志")
		loger.Warning("这是一条Warning日志")
		loger.Error("这是一条Error日志")
		loger.Fatal("这是一条Fatal日志")
	}
}
