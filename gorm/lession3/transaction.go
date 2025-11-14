package lession3

import (
	"gorm.io/gorm"
)

func RunTransaction(db *gorm.DB, callback func(tx *gorm.DB) error) error {
	tx := db.Begin()
	tx.Set("gorm:table_options", "ENGINE=InnoDB")
	if tx.Error != nil {
		return tx.Error
	}
	defer tx.Commit()

	// var user User
	// db.Create(&User{Name: "u1", Age: 10})
	// db.First(&user)

	// 禁用事务
	// tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	// tx.First(&user, 1)
	// tx.Model(&user).Update("Age", 18)

	// db.Transaction(func(tx *gorm.DB) error {
	// 	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	// 	if err := tx.Create(&User{Name: "Giraffe"}).Error; err != nil {
	// 		// 返回任何错误都会回滚事务
	// 		return err
	// 	}
	//
	// 	if err := tx.Create(&User{Name: "Lion"}).Error; err != nil {
	// 		return err
	// 	}
	//
	// 	// 返回 nil 提交事务
	// 	// return nil
	// 	return errors.New("transaction create error")
	// })

	// 嵌套事务
	// db.Transaction(func(tx *gorm.DB) error {
	// 	tx.Create(&User{Name: "u1"})
	//
	// 	tx.Transaction(func(tx2 *gorm.DB) error {
	// 		tx2.Create(&User{Name: "u2"})
	// 		return errors.New("rollback user2") // Rollback user2
	// 	})
	//
	// 	tx.Transaction(func(tx3 *gorm.DB) error {
	// 		tx3.Create(&User{Name: "u3"})
	// 		return nil
	// 	})
	//
	// 	return nil
	// })

	// 手动事务
	// 开始事务
	// tx := db.Begin()
	//
	// // 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	// tx.Create(&User{Name: "u1"})
	// tx.Create(&User{Name: "u2"})
	//
	// // 遇到错误时回滚事务
	// tx.Rollback()
	//
	// // 否则，提交事务
	// tx.Commit()

	// tx := db.Begin()
	// tx.Create(&User{Name: "u1"})
	//
	// tx.SavePoint("sp1")
	// tx.Create(&User{Name: "u2"})
	// tx.RollbackTo("sp1") // Rollback user2
	//
	// tx.Commit() // Commit user1

	return callback(tx)
}
