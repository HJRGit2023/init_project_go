package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	database "github.com/learn/gorm_gin/database"
	"github.com/learn/gorm_gin/model"
	"github.com/learn/gorm_gin/routers"
	"gorm.io/gorm"
)

func main() {
	//1. Connect to database
	var db *gorm.DB
	db = database.LoadDB()
	// defer db.DB().Close()
	//2. load models
	// migrate models
	db.AutoMigrate(&model.Userb{}, &model.Post{}, &model.Comment{})
	// 3. 初始化Gin引擎
	r := gin.Default()

	// 4. 注册路由（调用router包封装的路由逻辑）
	routers.InitRouter(r, db)

	// 5. 启动Gin服务
	err := r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		fmt.Printf("Gin服务启动失败：%v\n", err)
		panic(err)
	}
	fmt.Println("Gin服务启动成功")
}
