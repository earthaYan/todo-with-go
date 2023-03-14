package routes

import (
	"to-do-list/api"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	router.Use(sessions.Sessions("mysession", store))
	// 基础路由
	v1 := router.Group("api/v1")
	{
		// 用户操作
		// 注册
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)
	}
	return router
}
