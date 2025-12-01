package gormwork

import "gorm.io/gorm"

/*
- 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
  - 要求 ：
    - 定义一个 Book 结构体，包含与 books 表对应的字段。
    - 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，
	并将结果映射到 Book 结构体切片中，确保类型安全。
*/

type Book struct {
	ID     int `gorm:"primary_key"`
	Title  string
	Author string
	Price  float32
}

func RunWork2(db *gorm.DB) {
	db.AutoMigrate(&Book{})
	// db.Create(&Book{Title: "The Catcher in the Rye", Author: "J.D. Salinger", Price: 19.99})
	// db.Create(&Book{Title: "To Kill a Mockingbird", Author: "Harper Lee", Price: 12.99})
	// db.Create(&Book{Title: "1984", Author: "George Orwell", Price: 60.99})
	// db.Create(&Book{Title: "Pride and Prejudice", Author: "Jane Austen", Price: 77.99})
	var books []Book
	db.Debug().Model(&Book{}).Where("price > ?", 50).Find(&books)
	for _, book := range books {
		println("Title : ", book.Title, " Author : ", book.Author, "Price : ", book.Price)
	}
}
