package auth

import (
	"begingo/common/response"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginAuthMiddleware 中间件函数，用于检查用户是否已登录
func LoginAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Session中获取用户的登录状态
		session := sessions.Default(c)
		user := session.Get("currentUser")

		// 如果用户未登录，则重定向到登录页面
		if user == nil {
			//c.Redirect(http.StatusFound, "v1/user/login")
			response.Failed(c, -1, "用户未登录")
			// 停止处理程序链的执行
			c.Abort()
			return
		}

		// 如果用户已登录，则允许通过
		c.Next()
	}
}
