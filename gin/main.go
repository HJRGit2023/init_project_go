// @title Gin Web API
// @version 1.0
// @description RESTful API 文档
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 用户认证：模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

type User struct {
	Username string `form:"u" json:"u" xml:"u" binding:"required"`
	Password string `form:"p" json:"p" xml:"p" binding:"required"`
	Phone    string `form:"phone" binding:"required,phone"`
}

type User2 struct {
	Sex    string  `form:"sex" json:"sex" xml:"sex"`
	Age    int     `form:"age" json:"age" xml:"age"`
	Height float32 `form:"height" json:"height" xml:"height"`
}

// @Summary 用户登录
// @Tags auth
// @Accept  json
// @Produce json
// @Param   login body Login true "登录凭证"
// @Success 200 {object} Response
// @Router /login [post]
type LoginInfo struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"number"`
	Email    string `form:"email" json:"email" xml:"email" binding:"email"`
}

type Booking struct {
	CheckIn  time.Time `form:"check_in"  binding:"required,bookabledate" time_format:"2006-01-02""`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

// 自定义验证函数
var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

// 自定义验证函数
func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Booking is available!"})
	}
}

func getPhone(c *gin.Context) {
	var user User
	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"phone": user.Phone})
	}
}

// 自定义中间件
func middleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware1 before request")
		c.Next()
		fmt.Println("middleware1 after request")
	}
}

func middlesware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middlesware2 before request")
		c.Next()
		fmt.Println("middlesware2 after request")
	}
}

func main() {
	router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// -----------------------------Router GET/POST/Any-------------------------------
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello world Get")
	// })
	// router.POST("/login", func(c *gin.Context) {
	// 	c.String(200, "Hello world Post")
	// })
	// router.Any("/login", func(c *gin.Context) {
	// 	c.String(200, "Hello world Any")
	// })
	// -----------------------------Router Group---------------------------
	// grp1 := router.Group("/v1")
	// {
	// 	grp1.GET("/", func(c *gin.Context) {
	// 		c.String(200, "v1 Hello world Get")
	// 	})
	// }

	// grp2 := router.Group("/v2")
	// {
	// 	grp2.GET("/", func(c *gin.Context) {
	// 		c.String(200, "v2 Hello world Get")
	// 	})
	// }
	// -----------------------------Restful API--------------------
	// router.GET("/user", func(c *gin.Context) { c.String(200, "user get") })       // 获得
	// router.POST("/user", func(c *gin.Context) { c.String(200, "user post") })     // 新增
	// router.PUT("/user", func(c *gin.Context) { c.String(200, "user put") })       // 更新
	// router.DELETE("/user", func(c *gin.Context) { c.String(200, "user delete") }) // 删除
	// router.PATCH("/user", func(c *gin.Context) { c.String(200, "user patch") })   // 修改部分字段
	// -----------------------------Redirect-----------------------------
	// 重定向到外部
	// router.GET("/test", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	// })

	// // 重定向到内部
	// router.POST("/test", func(c *gin.Context) {
	// 	c.Redirect(http.StatusFound, "/foo")
	// })

	// router.GET("/test1", func(c *gin.Context) {
	// 	c.Request.URL.Path = "/test2"
	// 	router.HandleContext(c)
	// })
	// router.GET("/test2", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"hello": "world"})
	// })
	// -----------------------------Static----------------------
	// router.Static("/static", "./static") // 文件目录
	// router.StaticFS("/static", http.Dir("static")) // 文件目录
	// router.StaticFile("/f1", "./static/test1.txt") // 单独的文件
	// -----------------------------Output 输出XML/JSON/TOML/YAML/ProtoBuf-------------------
	// router.GET("/user", func(c *gin.Context) {
	// 	var user User = User{
	// 		Username: "admin",
	// 		Password: "123456",
	// 	}
	// c.JSON(200, &user)
	// c.XML(200, user)
	// 输出为 toml 格式，调用了postman,生成了user.toml文件，因为没有安装toml库
	// c.TOML(200, user)
	// 输出为 yaml 格式,会生成文件，浏览器下载，打开后不是yaml格式，因为没有安装yaml库
	// c.YAML(200, user)
	// c.ProtoBuf(200, user) // 输出为 protobuf 格式,报错
	// })
	// -----------------------------HTML模板--------------------
	// router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	// router.GET("/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "user.tmpl", gin.H{
	// 		"title": "Main website，Html测试",
	// 	})
	// })
	// -------templates/subs子目录
	// router.LoadHTMLGlob("templates/**/*")
	// router.GET("/posts/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "subs/posts.tmpl", gin.H{
	// 		"title": "Posts",
	// 	})
	// })
	// xxx.tmpl 必须在 templates 子目录下，不然会找不到模板文件，
	// 因为上面LoadHTMLGlob("templates/**/*")会加载子目录下的模板
	// router.GET("top", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "top/user.tmpl", gin.H{
	// 		"title": "Top User",
	// 	})
	// })
	// -----------------------------参数绑定----------------------
	// router.GET("/user/:username/:password", func(c *gin.Context) {
	// 	username := c.Param("username")
	// 	password := c.Param("password")
	// 	c.JSON(200, gin.H{
	// 		"username": username,
	// 		"password": password,
	// 	})
	// })
	// router.GET("/user", func(c *gin.Context) {
	// 	username := c.Query("username")
	// 	password := c.Query("password")
	// 	c.JSON(200, gin.H{
	// 		"username": username,
	// 		"password": password,
	// 	})
	// })
	// router.POST("/userForm", func(c *gin.Context) {
	// 	username := c.PostForm("username") // post表单，使用c.Query获取不到参数，使用c.PostForm获取参数
	// 	password := c.PostForm("password")
	// 	c.JSON(200, gin.H{
	// 		"username": username,
	// 		"password": password,
	// 	})
	// })
	// 这里:id([0-9]+)没啥用，不起作用，gin不支持这种写法
	// router.GET("/users/:id([0-9]+)", func(c *gin.Context) { //
	// 	id := c.Param("id")
	// 	// 处理数字ID...
	// 	c.JSON(200, gin.H{
	// 		"id": id,
	// 	})
	// })
	// 校验id是否为数字
	// router.GET("/users/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	if _, err := strconv.Atoi(id); err != nil {
	// 		c.JSON(400, gin.H{"error": "id must be a number"})
	// 		return
	// 	}
	// 	c.JSON(200, gin.H{"id": id})
	// })
	// ------------shoudBind 绑定参数--------------------------
	// router.POST("/user/", func(c *gin.Context) {
	// 	var user User
	// 	if c.ShouldBind(&user) == nil { // 绑定参数到结构体
	// 		c.JSON(200, gin.H{
	// 			"username": user.Username,
	// 			"password": user.Password,
	// 		})
	// 	} else {
	// 		c.JSON(400, gin.H{
	// 			"error": "参数绑定失败",
	// 		})
	// 	}
	// })
	// ---------------ShouldBindWith 绑定多个结构体参数-
	// router.POST("/user/", func(c *gin.Context) {
	// 	var user User
	// 	if c.ShouldBind(&user) == nil { // 绑定参数到结构体
	// 		c.JSON(200, gin.H{
	// 			"username": user.Username,
	// 			"password": user.Password,
	// 		})
	// 	} else {
	// 		c.JSON(400, gin.H{
	// 			"error": "参数绑定失败",
	// 		})
	// 	}
	// })
	// router.POST("/user/", func(c *gin.Context) {
	// 	user := User{}
	// 	user2 := User2{}
	// 	// 读取 c.Request.Body 并将结果存入上下文。接收多种结构体参数
	// 	if errUser := c.ShouldBindBodyWith(&user, binding.JSON); errUser == nil {
	// 		c.String(http.StatusOK, `the body should be user`)
	// 		c.JSON(http.StatusOK, user)
	// 		// 这时, 复用存储在上下文中的 body。
	// 	} else if errUser2 := c.ShouldBindBodyWith(&user2, binding.JSON); errUser2 == nil {
	// 		c.String(http.StatusOK, `the body should be user2 JSON`)
	// 		c.JSON(http.StatusOK, user2)
	// 		// 可以接受其他格式
	// 	} else if errUser22 := c.ShouldBindBodyWith(&user2, binding.XML); errUser22 == nil {
	// 		c.String(http.StatusOK, `the body should be user2 XML`)
	// 		c.XML(http.StatusOK, user2)
	// 	} else {
	// 		c.String(http.StatusBadRequest, `the body should be user or user2 JSON or formB XML`)
	// 	}
	// })
	// -----------------------------模型验证-----------------------
	// router.GET("/", func(c *gin.Context) {
	// 	login := LoginInfo{}
	// 	err := c.ShouldBind(&login)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, login)
	// })
	// -----------------------------自定义验证器-----------------
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("bookabledate", bookableDate) // 注册自定义验证器
	// }
	// router.GET("/bookable", getBookable)
	// ----------Phone验证器
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
	// 		return regexp.MustCompile(`^1[3-9]\d{9}$`).MatchString(fl.Field().String())
	// 	})
	// }
	// router.GET("/phone", getPhone)
	// -----------------------------自定义中间件-----------------------------
	// router.Use(middleware1()) // 注册全局中间件
	// router.Use(middlesware2())
	// router.GET("/ping", func(c *gin.Context) {
	// 	fmt.Println("self")
	// 	c.String(200, "server is running") // 响应请求，响应到页面或postman结果里，不会在控制台打印
	// })
	// -----------用户认证
	// authorized路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	// authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
	// 	"foo":    "bar",
	// 	"austin": "1234",
	// 	"lena":   "hello2",
	// 	"manu":   "4321",
	// }))
	// authorized.GET("/secret", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)
	// 	if secret, ok := secrets[user]; ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	// 	} else {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"user": user, "secret": "not secret"})
	// 	}
	// })
	// -----------------------------HTTPS-------------------------
	// router.RunTLS(":443", "server.crt", "server.key")
	// router.RunTLS(":8080", "./cert/server.crt", "./cert/server.key")
	// Ping handler
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.String(200, "pong")
	// })
	// m := autocert.Manager{
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
	// 	Cache:      autocert.DirCache("/var/www/.cache"),
	// }
	// log.Fatal(autotls.RunWithManager(router, &m))
	// -----------------------------Swagger 集成规范------------------------
	// 添加docs路由（main函数内）
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 配置swagger.json路径（初始化路由前添加）
	router.StaticFile("/swagger.json", "./docs/swagger.json")

	// 添加文档访问中间件
	authMiddleware := gin.BasicAuth(gin.Accounts{
		"admin": "swagger123",
	})
	// 文档访问路径，中间件authMiddleware可以控制文档的访问权限，
	router.GET("/docs", authMiddleware, ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		panic(err)
	}

}
