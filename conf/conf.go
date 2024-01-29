package conf

import (
	"begingo/common/log"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	log.BuildLogger(os.Getenv("LOG_LEVEL"))
}
