package main

import (
	"fmt"
	// "github.com/learn/gorm/lession1"
	// "github.com/learn/gorm/lession2"
	// "github.com/learn/gorm/lession3"
	"github.com/learn/gorm/gormwork"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
	**注意：**想要正确的处理 time.Time ，您需要带上 parseTime 参数，

(更多参数) 要支持完整的 UTF-8 编码，您需要将 charset=utf8 更改为 charset=utf8mb4
*/
func main() {
	// 方式一：连接数据库
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "gorm:gorm1234@tcp(127.0.0.1:3306)/gorm_config?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 方式二：连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		CreateBatchSize: 1000, // 初始化 batch size，默认1000 ，之后的批量插入将使用该值
	})
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return
	}
	fmt.Println("连接数据库成功")
	// lession1.Run(db)
	// lession2.Run(db)
	// lession2.RunUpdate(db)
	// lession3.Run(db)
	gormwork.RunWork33(db)
}
