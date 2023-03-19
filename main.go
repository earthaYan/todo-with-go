package main

import (
	"to-do-list/conf"
	"to-do-list/routes"

	_ "to-do-list/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title TodoList API
// @version 0.0.1
// @description Backend of Todo
// @BasePath /api/v1
// @Host 116.204.108.126:3000
// @securityDefinitions.basic  BasicAuth
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
func main() { // http://localhost:3000/swagger/index.html
	// 从配置文件读入配置
	conf.Init()
	// 转载路由 swag init -g main.go
	// 路由配置
	r := routes.NewRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 运行在3000端口
	_ = r.Run(conf.HttpPort)

}
