package gormwork

import "gorm.io/gorm"

// 用户
type Userblog struct {
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(255);not null"`
	Age       int    `gorm:"type:int"`
	Posts     []Post `gorm:"foreignkey:UserID"`
	PostCount int    `gorm:"column:post_count"`
}

// 文章
type Post struct {
	ID            int       `gorm:"primary_key"`
	Title         string    `gorm:"type:varchar(255);not null"`
	Content       string    `gorm:"type:text;not null"`
	UserID        int       `gorm:"not null"`
	Comments      []Comment `gorm:"foreignkey:PostID"`
	CommentStatus string
}

// 评论
type Comment struct {
	ID      int `gorm:"primary_key"`
	Title   string
	Content string
	PostID  int
}

func RunWork3(db *gorm.DB) {
	db.AutoMigrate(&Userblog{}, &Post{}, &Comment{})
	// 创建用户
	users := []Userblog{
		{Name: "user1", Age: 20},
		{Name: "user2", Age: 25},
		{Name: "user3", Age: 30},
	}
	db.Debug().Create(&users)
	// 创建文章
	posts := []Post{
		{Title: "post1_1", Content: "content1_1", UserID: 1},
		{Title: "post1_2", Content: "content1_2", UserID: 1},
		{Title: "post2_1", Content: "content2_1", UserID: 2},
		{Title: "post2_2", Content: "content2_2", UserID: 2},
		{Title: "post3", Content: "content3", UserID: 3},
	}
	db.Debug().Create(&posts)
	// 创建评论
	comments := []Comment{
		{Title: "comment1_1", Content: "content_1", PostID: 1},
		{Title: "comment1_2", Content: "content_2", PostID: 1},
		{Title: "comment1_3", Content: "content_3", PostID: 1},
		{Title: "comment2_1", Content: "content2_1", PostID: 2},
		{Title: "comment2_2", Content: "content2_2", PostID: 2},
		{Title: "comment2_3", Content: "content2_3", PostID: 2},
		{Title: "comment3_1", Content: "content3_1", PostID: 3},
	}
	db.Debug().Create(&comments)
}
