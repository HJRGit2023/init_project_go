package lession1

import (
	"context"
	"fmt"

	//"gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite" // 替换原 sqlite 驱动
	"gorm.io/gorm"
)

type Product1 struct {
	gorm.Model
	Code  string
	Price uint
}

func Run1(db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// context.Background() 创建的是一个无超时、无取消信号的 “根上下文”，
	// 它是所有其他上下文的起点，本身不包含超时设置。
	ctx := context.Background()
	// Migrate the schema 迁移架构/模式
	db.AutoMigrate(&Product1{})
	// Create
	// db.WithContext(ctx).Create(&Product{Code: "L1212", Price: 1000})
	err = gorm.G[Product1](db).Create(ctx, &Product1{Code: "D42", Price: 1000})

	// Read
	// Find product by id
	product, err := gorm.G[Product1](db).Where("id = ?", "1").First(ctx)
	fmt.Println("product：", product)
	// Find product by code
	products, err := gorm.G[Product1](db).Where("Code = ? ", "D42").Find(ctx)
	fmt.Println("products：", products)
	// Update
	_, err = gorm.G[Product1](db).Where("id = ?", product.ID).Update(ctx, "Price", 2000)
	fmt.Println("update result：", err)
	fmt.Println("updated product：", product)
	_, err = gorm.G[Product1](db).Where("id = ?", product.ID).Updates(ctx, Product1{Code: "D42", Price: 2000})
	fmt.Println("updates result：", err)
	fmt.Println("updated2 product：", product)
	// Delete
	_, err = gorm.G[Product1](db).Where("id = ?", product.ID).Delete(ctx)
	fmt.Println("delete result：", err)
	fmt.Println("deleted product：", product)
}
