package lession3

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name            string
	CompanyID       int
	Company         Company      `gorm:"foreignKey:CompanyID"`                                             // belong to Company
	Orders          []Order      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 配置外键及级联删除"` 一对多关系
	Languages       []Language   `gorm:"many2many:user_languages;"`                                        // 多对多关系
	MemberNumber    string       `gorm:"index:idx_users_member_number"`
	CreditCards     []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Friends         []*User      `gorm:"many2many:user_friends"` // 自引用-多对多关系
	BillingAddress  Address
	ShippingAddress Address
	Emails          []Email `gorm:"foreignKey:UserID"`
}

type Company struct {
	ID   int
	Name string
}

// -------------------------自定义类型order 用于查询结果----------
type Order struct {
	gorm.Model
	UserID     uint
	Name       string
	Amount     float64
	FinishedAt time.Time `gorm:"autoUpdateTime"`
}

type Language struct {
	gorm.Model
	Name string
}

type CreditCard struct {
	gorm.Model
	Number     string
	UserNumber string
}

// -------------------------自定义类型email 用于查询结果----------
type Email struct {
	ID     uint
	UserID uint
	Email  string
}

type Address struct {
	ID       uint
	UserID   uint
	Address1 string
}

// 检索用户列表并预加载信用卡
func GetAllCreditCards(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Debug().Model(&User{}).Preload("CreditCard").Find(&users).Error
	return users, err
}

// 检索 User 列表并预加载 Language
func GetAllLanUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Debug().Model(&User{}).Preload("Languages").Find(&users).Error
	return users, err
}

// 检索 Language 列表并预加载 User
func GetAllUserLanguages(db *gorm.DB) ([]Language, error) {
	var languages []Language
	err := db.Debug().Model(&Language{}).Preload("Users").Find(&languages).Error
	return languages, err
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Company{}, &Order{}, &Language{}, &CreditCard{}, &Email{}, &Address{})
	// db.Create(&Company{Name: "Company0", ID: 4})
	// var users = []User{
	// 	{Name: "John Smith", CompanyID: 1},
	// 	{Name: "John Doee", CompanyID: 2},
	// 	{Name: "Jane Doff", CompanyID: 3},
	// }
	// db.Create(&users)
	// var orders = []Order{
	// 	{UserID: 1, Name: "order1", Amount: 100.0},
	// 	{UserID: 2, Name: "order2", Amount: 202.0},
	// 	{UserID: 3, Name: "order3", Amount: 303.0},
	// }
	// db.Create(&orders)
	// 预加载 Orders
	// db.Debug().Preload("Orders").Find(&users)
	// 预加载 Languages
	// go GetAllCreditCards(db)
	// go GetAllLanUsers(db)
	// go GetAllUserLanguages(db)
	// ---------------------------在创建时自动保存关联数据----------------------------
	// 创建一条新的记录时，GORM会自动保存它的关联数据。 这个过程包括向关联表插入数据以及维护外键引用。
	// user := User{
	// 	Name:            "jinzhuSkip2",
	// 	BillingAddress:  Address{Address1: "Billing Address - Address 1"},
	// 	ShippingAddress: Address{Address1: "Shipping Address - Address 1"},
	// 	Emails: []Email{
	// 		{Email: "jinzhuSkip2@example.com"},
	// 		{Email: "jinzhuSkip2-2@example.com"},
	// 	},
	// 	Languages: []Language{
	// 		{Name: "ZH"},
	// 		{Name: "EN"},
	// 		{Name: "JP"},
	// 		{Name: "US"},
	// 	},
	// 	CompanyID: 2,
	// }
	// db.Debug().Create(&user)
	// db.Debug().Save(&user)
	// ---------------------------FullSaveAssociations来更新关联数据----------------
	// 更新用户并完全更新其所有关联 (不只外键)
	// user := User{}
	// db.Debug().Where("id = ?", 2354).First(&user)
	// db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	// SQL：完全更新地址、用户、电子邮件表，包括现有的关联记录
	// ----------------------------跳过自动创建、更新-------------------------------
	// 当插入用户的时候仅包含“Name”字段
	// db.Select("Name").Create(&user)
	// SQL: INSERT INTO "users" (name) VALUES ("jinzhu");
	// 跳过自动创建、更新关联数据
	// 创建用户时跳过字段“BillingAddress”
	// db.Omit("BillingAddress").Create(&user)

	// 创建用户时跳过全部关联关系
	// db.Debug().Omit(clause.Associations).Create(&user)

	// 跳过更新"Languages"关联
	// user := User{}
	// db.Debug().Where("id = ?", 2355).First(&user)
	// user.Languages = []Language{
	// 	{Name: "ZH"},
	// 	{Name: "EN"},
	// 	{Name: "JP"},
	// 	{Name: "US"},
	// }
	// db.Omit("Languages.*").Create(&user) // 创建时会报错，因为Languages字段是空的
	// db.Omit("Languages.*").Save(&user)
	// 跳过创建 'Languages' 关联及其引用
	// db.Debug().Omit("Languages").Save(&user)
	// ---------------------------删除关联-------------------------------
	// 删除 user 时，也删除 user 的 Orders、CreditCards 关联记录
	// user := User{}
	// db.Debug().Select("Orders", "CreditCards").Where("id = ?", 1).Delete(&user)

	// 删除用户 时，也删除用户的所有一对一、一对多和多对多关联
	// db.Debug().Select(clause.Associations).Where("id = ?", 1).Delete(&user)
	// ---------------------------启动关联模式-----------------------
	var user User
	var languages []Language
	db.Debug().Model(&languages).Where("id in ?", []int{4, 5}).Find(&languages)
	// var newLanguage Language = Language{Name: "ja-JP"}
	db.Debug().Where("id = ?", 2354).First(&user)
	// db.Debug().Model(&user).Association("Languages").Find(&languages)
	// db.Debug().Model(&user).Association("Languages").Append(&newLanguage)
	// db.Model(&user).Association("Languages").Replace(&newLanguage, &Language{Name: "JP"})
	// db.Debug().Model(&user).Association("Languages").Delete([]Language{newLanguage, Language{Name: "JP"}})
	db.Debug().Model(&user).Association("Languages").Delete(&languages)
	// 检查error
	// error := db.Debug().Model(&user).Association("Languages").Error
	// fmt.Println(error)
	// Find with conditions
	codes := []string{"ZH", "EN", "ja-JP", "JP"}
	// db.Debug().Model(&user).Where("code IN ?", codes).Association("Languages").Find(&languages)
	count1 := db.Debug().Model(&user).Where("name IN ?", codes).Association("Languages").Count()
	fmt.Println(count1)

}
