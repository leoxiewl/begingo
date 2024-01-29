package server

import (
	v1 "begingo/api/v1"
	"begingo/dao/mysql"
	"github.com/gin-gonic/gin"
	"os"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 连接数据库
	daoIns, _ := mysql.GetMySQLFactory(os.Getenv("MYSQL_DSN"))

	// 路由
	routerv1 := r.Group("/v1")
	{
		userv1 := routerv1.Group("/user")
		{
			userHandler := v1.NewUserHandler(daoIns)
			userv1.POST("/register", userHandler.Register)
		}
	}

	return r
}
