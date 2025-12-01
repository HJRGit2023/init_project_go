package handler

import (
	"fmt"

	"github.com/learn/gorm_gin/model"
	"gorm.io/gorm"
)

func UserbCreate(db *gorm.DB, userb *model.Userb) error {
	err := db.Create(&userb).Error
	fmt.Println(err)
	return err
}

func UserbUpdate(db *gorm.DB, userb *model.Userb) error {
	err := db.Updates(&userb).Error
	return err
}

func UserbDelete(db *gorm.DB, userb *model.Userb) error {
	err := db.Delete(&userb).Error
	return err
}

func UserbList(db *gorm.DB, userbs *[]model.Userb) error {
	err := db.Find(&userbs).Error
	return err
}

func UserbDetail(db *gorm.DB, id uint, username string, userb *model.Userb) error {
	if id > 0 {
		err := db.Where("id = ?", id).First(&userb).Error
		return err
	}
	if username != "" {
		err := db.Where("username = ?", username).First(&userb).Error
		return err
	}
	return nil
}

func PostCreate(db *gorm.DB, post *model.Post) error {
	err := db.Create(&post).Error
	return err
}

func PostUpdate(db *gorm.DB, id string, post *model.Post) error {
	err := db.Where("id = ?", id).Updates(&post).Error
	return err
}

func PostDelete(db *gorm.DB, id string) error {
	err := db.Where("id = ?", id).Delete(&model.Post{}).Error
	return err
}

func PostList(db *gorm.DB, posts *[]model.Post, userbID string) error {
	err := db.Where("userb_id = ?", userbID).Find(&posts).Error
	return err
}

func PostDetail(db *gorm.DB, id string, post *model.Post) error {
	err := db.Where("id = ?", id).Find(&post).Error
	return err
}

func CommentCreate(db *gorm.DB, comment *model.Comment) error {
	err := db.Create(&comment).Error
	return err
}

func CommentUpdate(db *gorm.DB, id string, comment *model.Comment) error {
	err := db.Where("id = ?", id).Updates(&comment).Error
	return err
}

func CommentDelete(db *gorm.DB, id string, comment *model.Comment) error {
	err := db.Where("id = ?", id).Delete(&comment).Error
	return err
}

func CommentList(db *gorm.DB, comments *[]model.Comment, postID string) error {
	err := db.Where("post_id = ?", postID).Find(&comments).Error
	return err
}

func CommentDetail(db *gorm.DB, id uint) {
	var comment model.Comment
	db.Where("id = ?", id).First(&comment)
}
