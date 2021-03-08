package mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

//向文件里面写相关代码

//FileLogger 日志文件结构体
type FileLogger struct {
	Level       Loglevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errfileObj  *os.File
	maxFileSize int64
}

//NewFilelogger 构造函数
func NewFilelogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	Loglevel, err := parseLoglevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       Loglevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

//initFile() 打开文件
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open failed,err:%s", err)
		// return err
	}
	errfileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open failed err,err:%s", err)
		// return err
	}
	//日志文件已打开了
	f.fileObj = fileObj
	f.errfileObj = errfileObj
	return nil
}

func (f *FileLogger) checkSizi(file *os.File) bool {
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err :%s\n", err)
		return false
	}
	//判断日志文件大小是否超过文件最大值
	return fileinfo.Size() >= f.maxFileSize
}

func (f *FileLogger) splitfile(file *os.File) (*os.File, error) {
	//切割日志文件
	//1.关闭当前文件
	// f.fileObj.Close()
	//2.备份一下	rename
	nowStr := time.Now().Format("20060102150405000")
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err :%s\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileinfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)

	f.fileObj.Close()

	os.Rename(logName, newLogName)
	//3.打开新文件并赋值给原值
	newFileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open failed,err:%s", err)
		return nil, err
	}
	return newFileObj, nil
}

//格式化输出2.0，使用空接口传参
func (f *FileLogger) log(lv Loglevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, flieName, lineNo := getInfo(3)
	if f.checkSizi(f.fileObj) {
		newfile, err := f.splitfile(f.fileObj)
		if err != nil {
			return
		}
		f.fileObj = newfile
	}

	fmt.Fprintf(f.fileObj, "[%s] [%7s] [%s:%s:%v] %s\n", now.Format("2006-01-02 15:04:05"), unparseLoglevel(lv), funcName, flieName, lineNo, msg)
	if lv >= ERROR {
		if f.checkSizi(f.errfileObj) {
			newfile, err := f.splitfile(f.errfileObj)
			if err != nil {
				return
			}
			f.errfileObj = newfile
		}
		//如果日志级别》=err，还需要在err中记录
		fmt.Fprintf(f.errfileObj, "[%s] [%7s] [%s:%s:%v] %s\n", now.Format("2006-01-02 15:04:05"), unparseLoglevel(lv), funcName, flieName, lineNo, msg)
	}

}

//关闭文件
func (f *FileLogger) close() {
	f.fileObj.Close()
	f.errfileObj.Close()
}
