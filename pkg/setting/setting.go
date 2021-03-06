package setting

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

var (
	Cfg *ini.File

	RUNMODE string

	//app
	PAGE_SIZE  int
	JWT_SECRET string

	//server
	HTTPPORT     int
	READTIMEOUT  int
	WRITETIMEOUT int

	//database
	TYPE         string
	USER         string
	PASSWORD     string
	HOST         string
	NAME         string
	TABLE_PREFIX string

	//log
	SysLog_FILE_PATH   string
	InnerLog_FILE_PATH string
	SysLog_FILE_DIR    string
	InnerLog_FILE_DIR  string

	//mongo
	MonHost string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/config.ini")
	if err != nil {
		log.Fatal("cant't load config.ini.", err)
	}
	LoadBase()
	LoadApp()
	LoadServer()
	LoadDatabase()
	LoadLog()
	LogMongo()
	//检测日志目录是否存在
	_, err = os.Stat(SysLog_FILE_DIR)
	if os.IsNotExist(err) {
		mkdir(SysLog_FILE_DIR)
	}
	_, err = os.Stat(InnerLog_FILE_DIR)
	if os.IsNotExist(err) {
		mkdir(InnerLog_FILE_DIR)
	}
}

func mkdir(path string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func LoadBase() {
	RUNMODE = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadApp() {
	PAGE_SIZE = Cfg.Section("app").Key("PAGE_SIZE").MustInt(15)
	JWT_SECRET = Cfg.Section("app").Key("JWT_SECRET").MustString("love$vesan")
}

func LoadServer() {
	HTTPPORT = Cfg.Section("server").Key("HTTP_PORT").MustInt(8080)
	READTIMEOUT = Cfg.Section("server").Key("READ_TIMEOUT").MustInt(60)
	WRITETIMEOUT = Cfg.Section("server").Key("WRITE_TIMEOUT").MustInt(60)
}

func LoadDatabase() {
	TYPE = Cfg.Section("database").Key("TYPE").MustString("mysql")
	USER = Cfg.Section("database").Key("USER").MustString("root")
	PASSWORD = Cfg.Section("database").Key("PASSWORD").MustString("root")
	HOST = Cfg.Section("database").Key("HOST").MustString("127.0.0.1:3306")
	NAME = Cfg.Section("database").Key("NAME").MustString("w3fy")
	TABLE_PREFIX = Cfg.Section("database").Key("TABLE_PREFIX ").MustString(" w3fy_")
}

func LoadLog() {
	SysLog_FILE_DIR = Cfg.Section("log").Key("sysLog_FILE_DIR").MustString("runtime/logs/SystemLogs")
	InnerLog_FILE_DIR = Cfg.Section("log").Key("innerLog_FILE_DIR").MustString("runtime/logs/ErrorLogs")
	SysLog_FILE_PATH = Cfg.Section("log").Key("sysLog_FILE_PATH").MustString("runtime/logs/SystemLogs/syslog")
	InnerLog_FILE_PATH = Cfg.Section("log").Key("innerLog_FILE_PATH").MustString("runtime/logs/ErrorLogs/errlog")
}

func LogMongo() {
	MonHost = Cfg.Section("mongo").Key("HOST").MustString("mongodb://localhost:27017")
}
