package auth

import (
	"github.com/ddh-open/gin/framework/gin"
)

// MiddlewareAuth 登录中间件
func MiddlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
