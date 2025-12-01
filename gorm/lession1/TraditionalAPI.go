package lession1

import (
	"fmt"

	"github.com/glebarez/sqlite" // 替换原 sqlite 驱动
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Run2(db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})
	// Create
	db.Create(&Product{Code: "L1212", Price: 100})
	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	db.First(&product, "code = ?", "L1212")
	fmt.Println(product)
	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	fmt.Println("updated : ", product)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Code: "L2424", Price: 300})
	fmt.Println("updateds2 : ", product)
	db.Model(&product).Updates(map[string]interface{}{"Code": "L3636", "Price": 400})
	fmt.Println("updateds3 : ", product)
	// Delete
	db.Delete(&product, 1)
	fmt.Println("deleted : ", product)
}
