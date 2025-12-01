package lession1

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	Name  string
	Email string
}

// Blog 结构体嵌入 Author 结构体
type Blog struct {
	// Author   // Embedded struct 嵌入 Author 结构体
	Author  `gorm:"embedded;embeddedPrefix:author_"` // 指定嵌入 Author 结构体字段的前缀
	ID      int
	Upvotes int32
}

// Blog 等价于 Blog2，只是字段类型不同，ID 字段类型不同，equals
type Blog2 struct {
	ID int64
	// Name    string
	// Email   string
	AuthorName  string
	AuthorEmail string
	// Author  Author `gorm:"embedded"` // 需要使用gorm:"embedded"指定是嵌入的结构体，否则会报错
	Upvotes int32 `gorm:"column:up_votes"` // 指定字段名为 up_votes
}
type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

type UserCol struct {
	gorm.Model
	NameC           string `gorm:"<-:create"`            // 允许读和创建
	NameU           string `gorm:"<-:update"`            // 允许读和更新
	NameCURW        string `gorm:"<-"`                   // 允许读和写（创建和更新）
	NameR           string `gorm:"<-:false"`             // 允许读，禁止写
	NameOnlyR       string `gorm:"->"`                   // 只读（除非有自定义配置，否则禁止写）
	NameRW          string `gorm:"->;<-:create"`         // 允许读和写
	NameOnlyC       string `gorm:"->:false;<-:create"`   // 仅创建（禁止从 db 读）
	Updated         int64  `gorm:"autoUpdateTime:nano"`  // 使用时间戳纳秒数填充更新时间
	Updated2        int64  `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	Created         int64  `gorm:"autoCreateTime"`       // 使用时间戳秒数填充创建时间
	NameSmallLetter string `gorm:"-"`                    // 通过 struct 读写会忽略该字段
	NameRWIgnMig    string `gorm:"-:all"`                // 通过 struct 读写、迁移会忽略该字段
	NameMigration   string `gorm:"-:migration"`          // 通过 struct 迁移会忽略该字段
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	// 先删除旧表
	db.Migrator().DropTable(&Blog{})
	db.AutoMigrate(&Blog{})
	// 先删除旧表
	db.Migrator().DropTable(&Blog2{})
	db.AutoMigrate(&Blog2{})
	fmt.Println("create table Blog and Blog2 success")
	// 先删除旧表
	db.Migrator().DropTable(&UserCol{})
	db.AutoMigrate(&UserCol{}) // 执行db.Create时会自动创建表
	var str *string
	str1 := "johndoe@example.com"
	str = &str1
	now := time.Now()

	// 插入数据
	user := User{
		Name:         "John Doe",
		Email:        str,
		Age:          30,
		Birthday:     &now,
		MemberNumber: sql.NullString{String: "M12345", Valid: true},
		ActivatedAt:  sql.NullTime{Time: time.Now(), Valid: true},
	}
	user1 := User{
		Name: "John Dofg",
		Age:  30,
	}
	user2 := User{
		Age: 30,
	}
	user2.MemberNumber.Valid = true // 验证MemberNumber 字段string，不能存null

	userCol := UserCol{
		NameC:     "NameC",     // 允许读和创建
		NameU:     "NameU",     // 允许读和更新
		NameCURW:  "NameCURW",  // 允许读和写（创建和更新）
		NameR:     "NameR",     // 允许读，禁止写
		NameOnlyR: "NameOnlyR", // 只读（除非有自定义配置，否则禁止写）
		NameRW:    "NameRW",    // 允许读和写
		NameOnlyC: "NameOnlyC", // 仅创建（禁止从 db 读）
		//NameSmallLetter: "NameSmallLetter", // 通过 struct 读写会忽略该字段
		//NameRWIgnMig:    "NameRWIgnMig",    // 通过 struct 读写、迁移会忽略该字段
		//NameMigration:   "NameMigration",   // 通过 struct 迁移会忽略该字段
	}
	db.Create(&user)  // 使用 Create 方法插入数据
	db.Create(&user1) // 验证 *string、*time.Time、sql.NullString、sql.NullTime可以存为null
	db.Create(&user2) // 验证Name 字段string，不能存null
	// 验证字段权限控制，不需要创建到数据库的字段需要注释，显示排除NameMigration
	// 因为该字段使用`gorm:"-:migration"`只是迁移时会忽略，建表就没有该字段，
	// 但是db.First查数据依然会有该字段，会报未知字段错误
	db.Omit("NameMigration").Create(&userCol)
	// 查询数据
	var result UserCol
	//db.First(&result, userCol.ID)
	db.Select("NameC", "NameU").Find(&result)
	// fmt.Println("database value : ", result)
	// fmt.Println("struct value : ", userCol)

}
