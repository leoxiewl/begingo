package mysql

import (
	log2 "begingo/common/log"
	"begingo/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type datastore struct {
	db *gorm.DB
}

func newDatastore(db *gorm.DB) *datastore {
	return &datastore{
		db: db,
	}
}

func (ds *datastore) Users() dao.UserDao {
	return newUsers(ds)
}

var (
	mysqlFactory dao.Factory
)

// GetMySQLFactory 在中间件中初始化mysql链接
func GetMySQLFactory(connString string) (dao.Factory, error) {
	// 初始化GORM日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound myerrors for log
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})
	// Error
	if connString == "" || err != nil {
		log2.Log().Error("mysql lost: %v", err)
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log2.Log().Error("mysql lost: %v", err)
		panic(err)
	}

	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(20)

	mysqlFactory = newDatastore(db)
	return mysqlFactory, nil
}
