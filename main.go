package main

import (
	"begingo/conf"
	"begingo/server"
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"os"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := server.NewRouter()

	err := r.Run(":8888")
	if err != nil {
		return
	}
}
