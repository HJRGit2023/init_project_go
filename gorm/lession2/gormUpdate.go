package lession2

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       uint
	Quantity    uint
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Role == "admin" {
		fmt.Println("admin user not allowed to update")
		return errors.New("admin user not allowed to update")
	}
	if tx.Statement.Changed("Name", "Admin") { // if Name or Role changed
		tx.Statement.SetColumn("Age", 18)
	}

	// if any fields changed
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("RefreshedAt", time.Now())
	}
	return nil
}

func RunUpdate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Product{})
	// ---------------------------Save()会保存所有的字段，即使字段是零值----------------------
	// var user User
	// db.First(&user) // find user with id = 1
	// user.Name = "jinzhu 2"
	// user.Age = 100
	// db.Debug().Save(&user)
	// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
	// db.Save(&User{Name: "jinzhuInsert", Age: 100})
	// INSERT INTO `users` (`name`,`age`,`birthday`,`update_at`) VALUES ("jinzhu",100,"0000-00-00 00:00:00","0000-00-00 00:00:00")
	// db.Save(&User{ID: 1, Name: "jinzhuUpdate", Age: 100})
	// UPDATE `users` SET `name`="jinzhu",`age`=100,`birthday`="0000-00-00 00:00:00",`update_at`="0000-00-00 00:00:00" WHERE `id` = 1
	// ---------------------------更新单列字段Generics API And Traditional API----------------------
	// ctx := context.Background()
	// Update with conditions
	// num, err := gorm.G[User](db).Where("id = ?", 11).Update(ctx, "name", "helloCondition")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;
	// Update with ID condition
	// num1, err := gorm.G[User](db).Where("id = ?", 12).Update(ctx, "name", "helloCondition")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;
	// Update with multiple conditions
	// num2, err := gorm.G[User](db).Where("id = ?", 13).Update(ctx, "name", "helloCondition")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

	// 根据条件更新
	// var user User
	// db.First(&user, 14)
	// db.Model(&User{}).Where("name = ?", "Tom").Update("name", "hello111")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

	// User 的 ID 是 `14`
	// db.Model(&user).Update("name", "hello111")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 根据条件和 model 的值进行更新
	// db.Model(&user).Where("name = ?", "Tom").Update("name", "hello111")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

	// ---------------------------更新多列字段Generics API And Traditional API-----------------------
	// ctx := context.Background()
	// Update attributes with `struct`, will only update non-zero fields
	// rows, err := gorm.G[User](db).Where("id = ?", 15).Updates(ctx, User{Name: "helloCond15", Age: 18, Active: false})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;
	// ---------这里不支持map Update attributes with `map`
	// rows1, err := gorm.G[User](db).Where("id = ?", 16).Updates(ctx, map[string]interface{}{"name": "helloCond16", "age": 18, "active": false})
	// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 根据 `struct` 更新属性，只会更新非零值的字段
	// var user User
	// db.First(&user, 17)
	// db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	// 根据 `map` 更新属性
	// db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	// ---------------------------更新选定字段Generics API And Traditional API---------
	// ctx := context.Background()

	// --------------泛型方式，不支持map Select with Map
	// rows, err := gorm.G[User](db).Where("id = ?", 18).Select("name").Updates(ctx, map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello' WHERE id=111;
	// --------------泛型方式，支持map Select with Map
	// rows, err := gorm.G[User](db).Where("id = ?", 18).Omit("name").Updates(ctx, map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
	// ------支持 Struct 方式，Select with Struct (select zero value fields)
	// rows, err := gorm.G[User](db).Where("id = ?", 18).Select("Name", "Age").Updates(ctx, User{Name: "new_name", Age: 0})
	// UPDATE users SET name='new_name', age=0 WHERE id=111;
	// Select all fields (select all fields include zero value fields)
	// rows, err := gorm.G[User](db).Where("id = ?", 18).Select("*").Updates(ctx, User{Name: "jinzhu", Role: "admin", Age: 0})
	// Select all fields but omit Role (select all fields include zero value fields)
	// rows1, err := gorm.G[User](db).Where("id = ?", 18).Select("*").Omit("Role").Updates(ctx, User{Name: "jinzhu", Role: "admin", Age: 0})

	// 选择 Map 的字段
	// User 的 ID 是 `111`:
	// var user User = User{ID: 19, Name: "Tom19", Age: 19, Active: true}
	// db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "helloSelect", "age": 18, "active": false})
	// UPDATE users SET name='hello' WHERE id=111;
	// db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "helloSelect2", "age": 18, "active": false})
	// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
	// 选择 Struct 的字段（会选中零值的字段）
	// user = User{ID: 20, Name: "Tom20", Age: 20, Active: true}
	// db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
	// UPDATE users SET name='new_name', age=0 WHERE id=111;
	// 选择所有字段（选择包括零值字段的所有字段）
	// user = User{ID: 21, Name: "Tom21", Age: 21, Active: true, Role: "admin1"}
	// db.Debug().Model(&user).Select("*").Updates(User{Name: "jinzhu", Role: "admin", Age: 0, ID: 21})
	// 选择除 Role 外的所有字段（包括零值字段的所有字段）
	// user = User{ID: 22, Name: "Tom22", Age: 22, Active: true}
	// db.Debug().Model(&user).Select("*").Omit("Role", "ID").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})

	// ----------------------------批量更新---------------------------------
	// Update with struct
	// db.Debug().Model(&User{}).Where("name = ?", "helloBatch").Updates(User{Name: "helloBatch23", Age: 24})
	// UPDATE users SET name='hello', age=18 WHERE role = 'admin';
	// Update with map
	// db.Debug().Table("users").Where("id IN ?", []int{24, 25}).Updates(map[string]interface{}{"name": "helloBatch2", "age": 23})
	// UPDATE users SET name='hello', age=18 WHERE id IN (24, 25);

	// ---------------------------阻止全局更新------------------------------
	// db.Debug().Model(&User{}).Update("name", "jinzhu") // gorm.ErrMissingWhereClause
	// db.Debug().Model(&User{}).Where("1 = 1").Update("name", "jinzhu")
	// UPDATE users SET `name` = "jinzhu" WHERE 1=1
	// db.Debug().Exec("UPDATE users SET name = ?", "jinzhu")
	// UPDATE users SET name = "jinzhu"
	// 使用session会话，开启全局更新AllowGlobalUpdate: true
	// db.Debug().Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("name", "jinzhu")
	// UPDATE users SET `name` = "jinzhu"

	// --------------------------------更新记录数-----------------------------
	// Get updated records count with `RowsAffected`
	// result := db.Debug().Model(&User{}).Where("role = ?", "admin").Updates(User{Name: "helloRowsAffected", Age: 18})
	// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

	// result.RowsAffected// returns updated records count
	// result.Error        // returns updating error
	// fmt.Println(result.RowsAffected, result.Error)

	// ---------------------------用 SQL 表达式更新列------------------------
	// db.Create(&Product{Name: "Product1", Description: "Product1Description", Price: 100, Quantity: 50})
	// db.Create(&Product{Name: "Product2", Description: "Product2Description", Price: 102, Quantity: 60})
	// db.Create(&Product{Name: "Product3", Description: "Product3Description", Price: 103, Quantity: 70})
	// var product Product
	// db.Debug().Find(&product, 3)
	// product's ID is `3`
	// db.Debug().Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
	// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;
	// product.ID = 4
	// db.Debug().Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
	// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;
	// product.ID = 5
	// db.Debug().Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
	// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3;
	// product.ID = 6
	// db.Debug().Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
	// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3 AND quantity > 1;

	// ---------------------------使用子查询更新一个表---------------------------------
	// user := User{ID: 26, Name: "Tom26", Age: 26, Active: true, CompanyID: 1}
	// db.Debug().Model(&user).Update("company_name", db.Model(&Company{}).Select("name").Where("companies.id = users.company_id"))
	// UPDATE "users" SET "company_name" = (SELECT name FROM companies WHERE companies.id = users.company_id);
	// db.Debug().Table("users as u").Where("name = ?", "jinzhu").Update("company_name", db.Table("companies as c").Select("name").Where("c.id = u.company_id"))
	// db.Debug().Table("users as u").Where("name = ?", "jinzhu").Updates(map[string]interface{}{"company_name": db.Table("companies as c").Select("name").Where("c.id = u.company_id")})

	// ---------------------------不使用 Hook 和时间追踪---------------------------
	/* 如果你希望更新时跳过 Hook 方法，并且不追踪更新的时间，你可以使用
	UpdateColumn, UpdateColumns, 它们的用法类似于 Update, Updates */
	// Update single column 更新单列字段
	// user := User{ID: 27}
	// db.Debug().Model(&user).UpdateColumn("name", "helloSkipHook")
	// UPDATE users SET name='hello' WHERE id = 111;
	// Update multiple columns更新多个字段
	// user = User{ID: 28}
	// db.Debug().Model(&user).UpdateColumns(User{Name: "helloSkipHook", Age: 28})
	// UPDATE users SET name='hello', age=18 WHERE id = 111;
	// Update selected columns 更新选定字段
	// user = User{ID: 29}
	// db.Debug().Model(&user).Select("name", "age").UpdateColumns(User{Name: "helloSkipHook", Age: 0})
	// UPDATE users SET name='hello', age=0 WHERE id = 111;

	// ---------------------------BeforeUpdate检查字段是否有变更
	// db.Debug().Model(&User{ID: 30, Name: "jinzhu"}).Updates(map[string]interface{}{"name": "jinzhu2"})
	// ----------------------------Delete--------------------
	// Email 的 ID 是 `10`
	// var email Email = Email{ID: 6}
	// db.Debug().Delete(&email)
	// DELETE from emails where id = 10;

	// 带额外条件的删除---软删除
	// db.Debug().Where("user_id = ?", 2328).Delete(&email)
	// DELETE from emails where id = 10 AND name = "jinzhu";
	// 查询被软删除的数据 // email.DeletedAt != nil, 如果 emails 表中没有 DeletedAt 字段，则无法查询被软删除的数据
	// db.Debug().Unscoped().Where("ID = 6").Find(&email)
	// ------------------------------批量删除------------------------------------
	// db.Debug().Where("name like ?", "order%").Delete(&Order{})
	orders := []Order{}
	// db.Debug().Unscoped().Where("name like ?", "order%").Find(&Order{}).Scan(&orders)
	// for _, order := range orders {
	// 	fmt.Println("delete order id:", order.ID)
	// }
	// -------------------------------永久删除 users---------------------------
	db.Debug().Unscoped().Where("name like ?", "order%").Delete(&orders)
	db.Debug().Unscoped().Where("name like ?", "order%").Find(&Order{}).Scan(&orders)
	for _, order := range orders {
		fmt.Println("delete user id after2:", order.ID)
	}
	// DELETE FROM orders WHERE id=10;
}
