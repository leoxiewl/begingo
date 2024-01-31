package auth

import (
	"begingo/common/response"
	"begingo/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AdminAuthMiddleware 中间件函数，用于检查用户是否是管理员
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Session中获取用户的登录状态
		session := sessions.Default(c)
		user := session.Get("currentUser")

		// 如果用户未登录，则重定向到登录页面
		if user == nil || user.(model.UserVO).UserRole != "admin" {
			response.Failed(c, -1, "用户未登录或者不是管理员")
			c.Abort()
			return
		}

		// 如果用户已登录，则允许通过
		c.Next()
	}
}
