package conf

import (
	"begingo/common/log"
	"github.com/go-playground/validator/v10"
	"os"

	"github.com/joho/godotenv"
)

var Validate *validator.Validate

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	log.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 校验初始化

	Validate = validator.New(validator.WithRequiredStructEnabled())
}
