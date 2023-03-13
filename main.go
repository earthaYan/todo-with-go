package main

import "to-do-list/conf"

func main() {
	// 从配置文件读入配置
	conf.Init()
	// 转载路由 swag init -g main.go
}
