package mylog

import (
	"errors"
	"strings"
)

// Loglevel 级别
type Loglevel uint16

//默认的级别数值常量
const (
	UNKNOWN Loglevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger 日志结构体
type Logger struct {
	Loglv Loglevel
}

//将传入的字符串转化为数值，好进行数值比较，以便级别输出
func parseLoglevel(s string) (Loglevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的级别")
		return UNKNOWN, err
	}

}

//将数值类型转为字符串，使其格式化输出
func unparseLoglevel(lv Loglevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}
