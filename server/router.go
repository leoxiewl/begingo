package server

import (
	"begingo/api"
	"begingo/dao/mysql"
	"github.com/gin-gonic/gin"
	"os"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 连接数据库
	daoIns, _ := mysql.GetMySQLFactory(os.Getenv("MYSQL_DSN"))

	// 路由
	v1 := r.Group("/v1")
	{
		userv1 := v1.Group("/user")
		{
			userHandler := api.NewUserHandler(daoIns)
			userv1.POST("/add", userHandler.Create)
		}
	}

	return r
}
