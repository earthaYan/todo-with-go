package api

import (
	"to-do-list/service"

	"github.com/gin-gonic/gin"
)

// UserRegister
// @Tags USER
// @Summary 用户注册
// @Produce json
// @Accept json
// @Param data body service.UserService true "用户名, 密码"
// @Success 200 {object} serializer.Response "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500  {object} serializer.Response "{"status":400,"data":{},"Msg":{},"Error":"error"}"
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

// UserLogin
// @Tags USER
// @Summary 用户登录
// @Produce json
// @Accept json
// @Param data body service.UserService true "用户名, 密码"
// @Success 200 {object} serializer.Response "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500  {object} serializer.Response "{"status":500,"data":{},"Msg":{},"Error":"error"}"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}
