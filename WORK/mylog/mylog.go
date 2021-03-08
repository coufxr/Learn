package mylog

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

// NewLog 构造函数
func NewLog(loglvStr string) Logger {
	level, err := parseLoglevel(loglvStr) //转化级别
	if err != nil {
		panic(err)
	}
	return Logger{
		Loglv: level, //输出级别
	}
}

//调用函数时的行号方法
func getInfo(skip int) (funcName, filename string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip) //调用函数时的位置，skip为调用层数
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	filename = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}

// Debug debug
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Trace debug
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

// Info debug
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning debug
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error debug
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal debug
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
