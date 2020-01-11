package logging

import (
	"log"
	"os"
	"time"
)

//日志类型
const (
	Error = iota
	System
)

//日志文件前缀
const (
	err = "error"
	sys = "sys"
)

var (
	ErrorLogsPath  = "runtime/logs/ErrorLogs/"
	SystemLogsPath = "runtime/logs/SystemLogs/"
	LogsExt        = ".log"
)

//构造logs文件
func logsPath(ftype int) string {
	var lp string
	t := time.Now().Format("20060102")
	switch ftype {
	case 0: //错误日志
		lp = ErrorLogsPath + err + t + LogsExt
	case 1: //系统日志
		lp = SystemLogsPath + sys + t + LogsExt
	}
	return lp
}

func openLogs(filepath string) *os.File {
	_, err := os.Stat(filepath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("can't not open file.", err)
	}
	return file
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+ErrorLogsPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
