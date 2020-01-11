package logging

import (
	"log"
	"os"
)

var (
	DefaultPrefix = ""
	EF            *os.File
	SF            *os.File
	elogger       *log.Logger
	slogger       *log.Logger
)

func init() {
	errorlosPath := logsPath(0)
	EF = openLogs(errorlosPath)
	systemlogsPath := logsPath(1)
	SF = openLogs(systemlogsPath)
	elogger = log.New(EF, DefaultPrefix, log.LstdFlags)
	slogger = log.New(SF, DefaultPrefix, log.LstdFlags)
}

func ErrorDebug(v ...interface{}) {
	elogger.Println(v)
}

func SystemDebug(v ...interface{}) {
	slogger.Println(v)
}
