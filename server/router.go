package server

import (
	v1 "begingo/api/v1"
	"begingo/dao/mysql"
	"begingo/middleware/auth"
	"begingo/model"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 连接数据库
	daoIns, _ := mysql.GetMySQLFactory(os.Getenv("MYSQL_DSN"))

	store := cookie.NewStore([]byte("secretbewind"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))
	gob.Register(model.UserVO{})

	// 路由
	routerv1 := r.Group("/v1")
	{
		userv1 := routerv1.Group("/user")
		{
			userHandler := v1.NewUserHandler(daoIns)
			userv1.POST("/register", userHandler.Register)
			userv1.POST("/login", userHandler.Login)
			userv1.GET("/logout", userHandler.Logout)
			userv1.POST("/add", auth.AdminAuthMiddleware(), userHandler.Create)
			userv1.POST("/delete", auth.AdminAuthMiddleware(), userHandler.Delete)
			userv1.POST("/update", auth.AdminAuthMiddleware(), userHandler.Update)
			userv1.GET("/get", auth.LoginAuthMiddleware(), userHandler.Get)
			userv1.GET("/current", userHandler.GetUserLogin)
			userv1.POST("/list/page", userHandler.ListPage)
		}
	}

	return r
}
