package main

import (
	"to-do-list/conf"
	"to-do-list/routes"
)

func main() {
	// 从配置文件读入配置
	conf.Init()
	// 路由配置
	r := routes.NewRouter()
	// 运行在3000端口
	_ = r.Run(conf.HttpPort)
}
