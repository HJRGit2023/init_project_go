package gormwork

import (
	"fmt"

	"gorm.io/gorm"
)

func RunWork32(db *gorm.DB) {
	//
	var user Userblog = Userblog{ID: 1}
	var conments []Comment
	// 预加载user.ID = 1 对应的Posts
	db.Debug().Model(&Userblog{}).Preload("Posts").Find(&user)
	err := db.Debug().Joins("inner join posts on comments.post_id = posts.id").Joins("inner join userblogs on posts.user_id = userblogs.id").Where("userblogs.id = ?", user.ID).Find(&conments).Error
	if err != nil {
		panic(err)
	}
	db.Debug().Model(&Comment{}).Select("count(*) as count, post_id").Group("post_id").Order("count desc").Limit(1).Offset(0).Find(&conments)
	// 通过comments.post_id 找到对应的post，
	var post Post
	db.Debug().Model(&Post{}).Where("id in (?)", conments[0].PostID).Find(&post)
	fmt.Println("评论最多的文章", post)
}
