package model

import "gorm.io/gorm"

type Userb struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required,email"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title" binding:"required"`
	Content string `gorm:"not null" json:"content" binding:"required"`
	UserbID uint   `json:"userb_id"`
	Userb   Userb  `json:"userb" binding:"-"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content" binding:"required"`
	UserbID uint   `json:"userb_id"`
	Userb   Userb  `json:"userb" binding:"-"`
	PostID  uint   `json:"post_id"`
	Post    Post   `json:"post" binding:"-"`
}
