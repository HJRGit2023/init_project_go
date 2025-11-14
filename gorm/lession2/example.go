package lession2

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID           uint           // Standard field for the primary key
	UUID         string         // A custom UUID field
	Role         string         // A regular string field
	Location     Location       // A struct field
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          `gorm:"default:18"` // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
	CreditCard   CreditCard     // A struct field that is a foreign key to CreditCard table
	Count        int            `gorm:"column:count;default:0"`                                         // A regular int field 首字母不能是小写，迁移时会忽略
	DeletedAt    gorm.DeletedAt `gorm:"index"`                                                          // A field that is used for soft delete软删除
	Orders       []Order        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // A struct field that is a foreign key to Order table
	EmailRef     Email          // A struct field that is a foreign key to Email table
	CompanyID    uint           // A regular uint field
	Company      Company        // A struct field that is a foreign key to Company table
	CompanyName  string         // A regular string field
	Active       bool           `gorm:"default:true"` // 自定义字段，默认值为true
	RefreshedAt  time.Time      // 自定义字段，用于记录刷新时间
}

// -------------------------CreditCard 表 是 User关联表，用于存储用户的信用卡信息
type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

// ----------Language 表 是 User关联表，用于存储用户的语言信息
type Language struct {
	ID   uint
	Code string
	Name string
}

// ----------自定义类型result 用于查询结果
type Result struct {
	Date  time.Time
	Total int
}

// ----------自定义类型ResultEmail 用于查询结果
type ResultEmail struct {
	Name  string
	Email string
}

type ResultUser struct {
	Name string
	Age  uint8
}

// -------------------------自定义类型order 用于查询结果----------
type Order struct {
	gorm.Model
	UserID     uint
	Name       string
	Amount     float64
	FinishedAt time.Time `gorm:"autoUpdateTime"`
}

// -------------------------自定义类型email 用于查询结果----------
type Email struct {
	ID     uint
	UserID uint
	Email  string
}

// ----------自定义类型company 用于查询结果----------
type Company struct {
	ID    uint
	Name  string
	Alive bool `gorm:"default:true"` // 自定义字段，默认值为true
}

// -------------------------自定义类型Location----------start
type Location struct {
	X, Y int
}

// Scan implements the sql.Scanner interface
func (loc *Location) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	return nil
}

func (loc Location) GormDataType() string {
	return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
	}
}

// ----------自定义类型Location----------End
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String() // UUID 需要转换成字符串

	if u.Role == "admin" {
		return errors.New("invalid role")
	}
	return
}

func Run(db *gorm.DB) {
	// 迁移 User 表
	db.AutoMigrate(&User{})
	// 迁移 CreditCard 表
	db.AutoMigrate(&CreditCard{})
	// 迁移 Language 表
	db.AutoMigrate(&Language{})
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Email{})
	db.AutoMigrate(&Company{})
	// 迁移 Location 表
	// email := "custom_email@example.com"
	// 插入数据
	// user := User{Name: "Tom", Age: 25, Email: &email}
	// fmt.Println(&user)
	// 需要使用指针，否则不会插入数据，因为ID是自动生成的,
	// 使用结构体的话，gorm会自动忽略ID，没有ID(主键)，所以不会插入数据
	// -------------------------使用结构体插入数据，失败----------
	// db.Create(user) //GORM 插入数据时因值不可寻址导致的反射异常
	// -------------------------插入单条数据----------
	// db.Create(&user)
	// 而使用指针的话，gorm会自动插入ID
	// db.Create(&User{Name: "Alice", Age: 18})
	// db.Create(&User{Name: "Bob", Age: 20})
	// -------------------------插入多条数据------------------------
	// users := []User{
	// 	{Name: "Cindy", Age: 22},
	// 	{Name: "David", Age: 23},
	// }
	// db.Create(&users)
	// -------------------------指定字段插入数据---------------------
	// db.Select("Name", "Age", "CreatedAt").Create(&user)
	// -------------------------创建记录并忽略传递给 ‘Omit’ 的字段值-----------
	// db.Omit("Name", "Age", "CreatedAt").Create(&User{Name: "Eve", Age: 21})
	// -------------------------批量插入数据返回记录的ID和Name---------------
	// var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	// db.Create(&users)
	// for _, user := range users {
	// 	// user.ID // 1,2,3 // 批量插入数据后，返回的记录的ID
	// 	// user.Name // jinzhu1,jinzhu2,jinzhu3 // 批量插入数据后，返回的记录的Name
	// 	fmt.Println(user.ID, user.Name)
	// }
	// ----------批量插入数据，指定batch size，使用CreateInBatches()---------------
	// var users = []User{}
	// for i := 0; i < 2000; i++ {
	// 	users = append(users, User{Name: fmt.Sprintf("dbBatchSize_%d", i)})
	// }
	// // batch size 100
	// db.CreateInBatches(users, 100)
	// -------------------------db.Create()使用CreateBatchSize()设置批量插入的大小----------
	// for i := 0; i < 2000; i++ {
	// 	users = append(users, User{Name: fmt.Sprintf("zhangsan_%d", i)})
	// }
	// db.Create(&users)
	// -------------------------db.BeforCreate()设置BeforeCreate()函数，在插入数据之前执行函数----------
	// db.Create(&User{Name: "admin", Age: 18})
	// -------------------------db.Session()设置Session，跳过钩子函数----------
	// db.Session(&gorm.Session{SkipHooks: true}).Create(&user)
	// -------------------------Map创建记录-------------------------------
	// db.Model(&User{}).Create(map[string]interface{}{
	// 	"Name": "LilyMap",
	// 	"Age":  24,
	// })
	// // batch insert from `[]map[string]interface{}{}`
	// db.Model(&User{}).Create([]map[string]interface{}{
	// 	{"Name": "lilyMap_1", "Age": 18},
	// 	{"Name": "lilyMap_2", "Age": 20},
	// })
	// -------------------------使用SQL插入数据-----------------------------
	// db.Exec("INSERT INTO users (name, age) VALUES ('SQL', 25)")
	// 方法一：Create from map
	// db.Model(User{}).Create(map[string]interface{}{
	// 	"Name":     "jinzhu",
	// 	"Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
	// })
	// 方法二：Create from struct
	// db.Model(User{}).Create(&User{
	// 	Name:     "jinzhu",
	// 	Location: Location{X: 100, Y: 100},
	// })
	// --------------------------关联表的插入------------------------
	// INSERT INTO `users` ...
	// INSERT INTO `credit_cards` ...
	// db.Create(&User{
	// 	Name:       "jinzhu",
	// 	CreditCard: CreditCard{Number: "411111111111"},
	// })
	// ------------------------Clauses使用--------------------------
	// ------------------------Clauses使用 Do nothing on conflict------------------
	// db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	// Update columns to default value on `id` conflict
	// 当插入 user 时，若 name 字段与已有数据冲突，则自动更新该记录的 age 和 email 字段，
	// 实现 “存在则更新，不存在则插入” 的逻辑（即 Upsert 操作）
	// email1 := "sameName1@example.com"
	// email2 := "sameName2@example.com"
	// email3 := "sameName3@example.com"
	// var users = []User{
	// 	{ID: 1, Name: "sameID1", Age: 35, Email: &email1, Count: 10},
	// 	{ID: 2, Name: "sameID", Age: 36, Email: &email2, Count: 2},
	// 	{ID: 2, Name: "sameID2", Age: 37, Email: &email3, Count: 3}}
	// ------------------------Clauses使用 On conflict 2-way update--------------
	// db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "name"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"age", "email"}),
	// }).Create(&users)
	//------------------------Clauses使用 Use SQL expression----------
	// db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "id"}},
	// 	DoUpdates: clause.Assignments(map[string]interface{}{"count": gorm.Expr("GREATEST(count, VALUES(count))")}),
	// }).Create(&users)
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `count`=GREATEST(count, VALUES(count));
	// ------------------------Clauses使用 id冲突更新 name和age字段----------
	// db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "id"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	// }).Create(&users)
	// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET "name"="excluded"."name"; SQL Server
	// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age"; PostgreSQL
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age`=VALUES(age); MySQL
	// ------------------------Clauses使用 更新All字段值----------
	// Update all columns to new value on conflict except primary keys and those columns having default values from sql func
	// db.Clauses(clause.OnConflict{
	// 	UpdateAll: true,
	// }).Create(&users)
	// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age", ...;
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age`=VALUES(age), ...; MySQL
	// --------------------------查询数据-------------------------------
	// ------------------------查询单条数据--------------------------
	// 泛型方式查询
	// ctx := context.Background()
	// user, err := gorm.G[User](db).First(ctx) // SELECT * FROM users ORDER BY id LIMIT 1;
	// fmt.Println("Generics first user : ", user)
	// user, err = gorm.G[User](db).Take(ctx) // SELECT * FROM users LIMIT 1;
	// fmt.Println("Generics take user : ", user)
	// user, err = gorm.G[User](db).Last(ctx) // SELECT * FROM users ORDER BY id DESC LIMIT 1;
	// fmt.Println("Generics last user : ", user)

	// 传统方式查询
	// db.First(&user) // SELECT * FROM users ORDER BY id LIMIT 1;
	// fmt.Println("Traditional first user : ", user)
	// db.Take(&user) // SELECT * FROM users LIMIT 1;
	// fmt.Println("Traditional take user : ", user)
	// db.Last(&user) // SELECT * FROM users ORDER BY id DESC LIMIT 1;
	// fmt.Println("Traditional last user : ", user)
	// ------------------------查询单条数据 First和Take的区别-
	// var user User
	// var users []User
	// ===works=== because destination struct is passed in
	// db.Debug().First(&user) // SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	// ===works=== because model is specified using `db.Model()`
	// result := map[string]interface{}{}
	// db.Debug().Model(&User{}).First(&result) // SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	// ===doesn't work===
	// result := map[string]interface{}{}
	// db.Debug().Table("users").First(&result)
	// ===works=== with Take
	// result := map[string]interface{}{}
	// db.Debug().Table("users").Take(&result) // no primary key defined, results will be ordered by first field (i.e., `Code`)

	// db.First(&Language{})
	// SELECT * FROM `languages` ORDER BY `languages`.`code` LIMIT 1
	// ------------------------主键查询---------------------
	// db.First(&user, 1) // SELECT * FROM users WHERE id = 1 LIMIT 1;
	// ctx := context.Background()
	// Using numeric primary key
	// user, err := gorm.G[User](db).Where("id = ?", 10).First(ctx) // SELECT * FROM users WHERE id = 10;
	// Using string primary key
	// user, err := gorm.G[User](db).Where("id = ?", "10").First(ctx) // SELECT * FROM users WHERE id = 10;
	// Using multiple primary keys
	// users, err := gorm.G[User](db).Where("id IN ?", []int{1, 2, 3}).Find(ctx) // SELECT * FROM users WHERE id IN (1,2,3);
	// If the primary key is a string (for example, like a uuid)
	// user, err := gorm.G[User](db).Where("uuid = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a").First(ctx)
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";
	// fmt.Println("err : ", err)
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	fmt.Println("Record not found")
	// 	return
	// } else if err != nil {
	// 	fmt.Println("Error occurred : ", err)
	// 	return
	// }
	// Traditional 主键 查询单条数据
	// user := User{}
	// users := []User{}
	// db.First(&user, 10) // SELECT * FROM users WHERE id = 10;
	// db.Debug().First(&user, "10") // SELECT * FROM users WHERE id = 10;
	// db.Find(&users, []int{1, 2, 3}) // SELECT * FROM users WHERE id IN (1,2,3);
	// db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a") // SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	// ------------------------查询多条数据-----------------------
	// Get all records
	// result := db.Find(&users) // SELECT * FROM users;
	// --------------------------查询数据,String条件查询----------------------
	// Get first matched record
	// result1 := db.Where("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	// Get all matched records
	// result2 := db.Where("name <> ?", "jinzhu").Find(&users)
	// SELECT * FROM users WHERE name <> 'jinzhu';
	// IN
	// result3 := db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');
	// LIKE
	// result4 := db.Where("name LIKE ?", "%jin%").Find(&users)
	// SELECT * FROM users WHERE name LIKE '%jin%';
	// AND
	// result5 := db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	// Time
	// lastWeek := time.Now().Add(-time.Hour * 24 * 7)
	// today := time.Now()
	// result6 := db.Where("updated_at > ?", lastWeek).Find(&users)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
	// BETWEEN
	// result7 := db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
	// fmt.Println("result1 : 条数: ", result1.RowsAffected) // returns found records count, equals `len(users)`
	// fmt.Println("result1 : Error: ", result1.Error)     // returns error
	// fmt.Println("result2 : 条数: ", result2.RowsAffected)
	// fmt.Println("result3 : 条数: ", result3.RowsAffected)
	// fmt.Println("result4 : 条数: ", result4.RowsAffected)
	// fmt.Println("result5 : 条数: ", result5.RowsAffected)
	// fmt.Println("result6 : 条数: ", result6.RowsAffected)
	// fmt.Println("result7 : 条数: ", result7.RowsAffected)
	// 如果设置了ID，下面的语句是使用and 拼接 id = 20 and id = 10
	// var user = User{ID: 10}
	// db.Debug().Where("id = ?", 20).First(&user) // SELECT * FROM users WHERE id = 10 and id = 20 ORDER BY id ASC LIMIT 1

	// ------------------------查询数据,Struct & Map & 切片条件--------------------
	// var user User
	// var users []User
	// Struct
	// db.Debug().Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

	// Map
	// db.Debug().Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// Slice of primary keys
	// db.Debug().Where([]int64{20, 21, 22}).Find(&users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);
	// 如果是Struct，则会忽略掉零值字段，比如Age字段为0
	// db.Debug().Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu";
	// 想要查询条件包含零值条件，可以使用Map条件查询
	// db.Debug().Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
	// 使用结构体字段名或者数据库字段名进行条件查询，指定具体值
	// db.Debug().Where(&User{Name: "jinzhu"}, "name", "Age").Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
	// GORM 零值匹配会使Name = "jinzhu"失效
	// db.Debug().Where(&User{Name: "jinzhu"}, "Age").Find(&users)
	// SELECT * FROM users WHERE age = 0;

	// ------------------------查询数据,内联条件------------
	// Get by primary key if it were a non-integer type
	// db.Debug().First(&user, "id = ?", "string_primary_key")	// SELECT * FROM users WHERE id = 'string_primary_key';
	// Plain SQL // 简单sql查询
	// db.Debug().Find(&user, "name = ?", "jinzhu")                   // SELECT * FROM users WHERE name = "jinzhu";
	// db.Debug().Find(&users, "name <> ? AND age > ?", "jinzhu", 20) // SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;
	// Struct
	// db.Debug().Find(&users, User{Age: 20}) // SELECT * FROM users WHERE age = 20;
	// Map
	// db.Debug().Find(&users, map[string]interface{}{"age": 20}) // SELECT * FROM users WHERE age = 20;

	// ------------------------查询数据,Not条件--------------------------
	// db.Debug().Not("name = ?", "jinzhu").First(&user) // SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;
	// Not In
	// db.Debug().Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");
	// Struct
	// db.Debug().Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;
	// Not In slice of primary keys
	// db.Debug().Not([]int64{1, 2, 3}).First(&user) // SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;

	// ------------------------查询数据,OR条件-----------------------------
	// db.Debug().Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

	// Struct
	// db.Debug().Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	// Map
	// db.Debug().Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	//--------------------------------查询数据，指定特定字段----------
	// db.Debug().Select("name", "age").Find(&users)
	// SELECT name, age FROM users;
	// db.Debug().Select([]string{"name", "age"}).Find(&users)
	// SELECT name, age FROM users;
	// db.Debug().Table("users").Select("COALESCE(age,?)", 42).Rows()
	// SELECT COALESCE(age,'42') FROM users;

	//----------查询数据，排序----------
	// db.Debug().Order("age desc, name").Find(&users) // SELECT * FROM users ORDER BY age desc, name;
	// Multiple orders
	// db.Debug().Order("age desc").Order("name").Find(&users) // SELECT * FROM users ORDER BY age desc, name;
	// 按 id 字段的指定顺序排序，优先级为 id=1→id=2→id=3，其余未在列表中的 id 会排在最后
	// db.Debug().Clauses(clause.OrderBy{
	// 	Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	// }).Find(&User{})
	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)

	// ------------------------
	// db.Debug().Limit(3).Find(&users) // SELECT * FROM users LIMIT 3;
	// var users, users1, users2 []User
	// Cancel limit condition with -1
	// db.Debug().Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)
	// db.Debug().Limit(10).Offset(3).Find(&users) // SELECT * FROM users limit 10 OFFSET 3;
	// db.Debug().Limit(10).Offset(5).Find(&users)     // SELECT * FROM users LIMIT 10 OFFSET 5 ;
	// Cancel offset condition with -1
	// db.Debug().Limit(20).Offset(10).Find(&users1).Offset(-1).Find(&users2)
	// SELECT * FROM users OFFSET 10; (users1)
	// SELECT * FROM users; (users2)

	// ------------------------查询数据,Group By & Having---------------------
	// result := map[string]interface{}{}
	// results := []map[string]interface{}{}
	// db.Debug().Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "jin%").Group("name").First(&result)
	// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1

	// db.Debug().Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
	// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"
	// db.Create(&User{Name: "group 1", Age: 22, Order: Order{Name: "order 1", Amount: 100}})
	// db.Create(&User{Name: "group 1", Age: 22, Order: Order{Name: "order 2", Amount: 101}})
	// db.Create(&User{Name: "group 1", Age: 22, Order: Order{Name: "order 3", Amount: 102}})
	// rows, err := db.Debug().Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
	// defer rows.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for rows.Next() {
	// 	var result2 Result
	// 	err := rows.Scan(&result2.Date, &result2.Total)
	// 	if err != nil {
	// 		fmt.Println("error : ", err)
	// 	}
	// 	fmt.Println("result2 : ", result2)
	// }

	// rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
	// defer rows.Close()
	// for rows.Next() {
	// 	var result3 Result
	// 	err := rows.Scan(&result3.Date, &result3.Total)
	// 	if err != nil {
	// 		fmt.Println("error : ", err)
	// 	}
	// 	fmt.Println("result3 : ", result3)
	// }

	// db.Debug().Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)
	// for _, result := range results {
	// 	fmt.Println(result)
	// }

	// ------------------------查询数据,Distinct---------------------------
	// users := []User{}
	// db.Debug().Distinct("name", "age").Order("name, age desc").Find(&users)
	// ------------------------查询数据，Joins------------------------------
	// resultEmail := ResultEmail{}
	// resultEmails := []ResultEmail{}
	// db.Create(&User{Name: "email 1", Age: 23, EmailRef: Email{Email: "email1@example.org"}})
	// db.Create(&User{Name: "email 1", Age: 23, EmailRef: Email{Email: "email2@example.org"}})
	// db.Create(&User{Name: "email 1", Age: 23, EmailRef: Email{Email: "email3@example.org"}})
	// db.Debug().Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Where("users.name = ?", "email 1").Scan(&ResultEmail{})
	// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

	// rows, err := db.Debug().Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Where("users.name = ?", "email 1").Rows()
	// fmt.Println(err)
	// defer rows.Close()
	// for rows.Next() {
	// 	var name string
	// 	var email string
	// 	err2 := rows.Scan(&name, &email)
	// 	if err2 != nil {
	// 		fmt.Println("error : ", err2)
	// 	}
	// 	fmt.Println("name : ", name, " email : ", email)
	// }

	// db.Debug().Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Where("users.name = ?", "email 1").Scan(&resultEmails)
	// fmt.Println(resultEmails)
	// multiple joins with parameter
	// db.Debug().Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&User{})

	// ------------------------查询数据，Joins预加载---------------------------
	// db.Create(&Company{Name: "company 1"})
	// db.Create(&Company{Name: "company 2"})
	// db.Create(&User{Name: "user 1", Age: 22, CompanyID: 1})
	// db.Create(&User{Name: "user 2", Age: 23, CompanyID: 2})
	// db.Create(&User{Name: "user 3", Age: 24, CompanyID: 1})
	// users := []User{}
	// db.Debug().Select([]string{"id", "name", "age", "company_id"}).Where("name like ?", "user%").Find(&users)
	// fmt.Println(users)
	// db.Debug().Joins("Company").Find(&users)
	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;

	// inner join
	// db.Debug().InnerJoins("Company").Find(&users)
	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` INNER JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;
	// db.Debug().Joins("Company", db.Where(&Company{Alive: true})).Find(&users)
	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id` AND `Company`.`alive` = true;
	// joins 表示：使用join 关联一个衍生表
	// orders := []Order{}
	// query := db.Debug().Table("orders").Select("MAX(orders.finished_at) as latest").Joins("left join users user on orders.user_id = user.id").Where("user.age > ?", 18).Group("orders.user_id")
	// db.Debug().Model(&Order{}).Joins("join (?) q on orders.finished_at = q.latest", query).Scan(&orders)
	// SELECT `order`.`user_id`,`order`.`finished_at` FROM `order` join (SELECT MAX(order.finished_at) as latest FROM `order` left join user user on order.user_id = user.id WHERE user.age > 18 GROUP BY `order`.`user_id`) q on order.finished_at = q.latest

	// ---------------------------查询数据，Scan-----------------------------------
	var result ResultUser
	db.Debug().Table("users").Select("name", "age").Where("name = ?", "Antonio").Scan(&result)

	// Raw SQL
	db.Debug().Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)
	fmt.Println("Raw SQL query result : ", result)
	// fmt.Println("Primary key query user : ", user)
}
