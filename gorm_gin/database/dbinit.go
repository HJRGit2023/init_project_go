package db

import (
	"log"

	errors "github.com/learn/gorm_gin/util/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func LoadDB() *gorm.DB {
	dsn := "gorm:gorm1234@tcp(127.0.0.1:3306)/gorm_config?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Info),
		CreateBatchSize: 1000,
	})
	if err != nil {
		panic(errors.ErrDBConnection)
	}
	log.Println("数据库连接成功", db)
	return db
}
