package middleware

import (
	"time"
	"to-do-list/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析失败",
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
