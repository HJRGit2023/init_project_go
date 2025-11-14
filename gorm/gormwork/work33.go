package gormwork

import (
	"fmt"

	"gorm.io/gorm"
)

var postId int

func (post *Post) AfterCreate(db *gorm.DB) (err error) {
	var count int = 0
	err = db.Debug().Table("Posts").Select("COUNT(*) AS count").Where("user_id = ?", post.UserID).Find(&count).Error
	if err != nil {
		return err
	}
	err = db.Debug().Model(&Userblog{}).Where("id = ?", post.UserID).Update("post_count", count).Error
	if err != nil {
		return err
	}
	return nil
}

func (comment *Comment) AfterDelete(db *gorm.DB) (err error) {
	var count int = 0
	fmt.Println("comment.PostID After", postId)
	err = db.Debug().Table("Comments").Select("COUNT(*) AS count").Where("post_id = ?", postId).Find(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		err = db.Debug().Model(&Post{}).Where("id = ?", postId).Update("Comment_Status", "无评论").Error
		if err != nil {
			return err
		}
	}
	return nil
}

func RunWork33(db *gorm.DB) {
	db.AutoMigrate(&Userblog{})
	db.AutoMigrate(&Post{})
	// post := Post{
	// 	UserID:  1,
	// 	Title:   "Hello World3",
	// 	Content: "This is my third post",
	// }
	// db.Debug().Create(&post)
	comment := Comment{}
	db.Debug().Where("title = ?", "comment3_1").Find(&comment)
	postId = comment.PostID
	db.Debug().Model(&Comment{}).Where("title = ?", "comment3_1").Delete(&comment)
}
