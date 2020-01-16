package main

import (
	"w3fy/routes"
)

func main() {
	// 禁用控制台颜色, 将日志写入文件时不需要控制台颜色
	r := routes.InitRoute()
	r.Run(":8080")
}
