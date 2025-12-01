package routers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/learn/gorm_gin/handler"
	"github.com/learn/gorm_gin/model"
	errors "github.com/learn/gorm_gin/util/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitRouter(r *gin.Engine, db1 *gorm.DB) *gin.Engine {
	// r := gin.Default()
	db = db1
	log.Println("registerRouters db : ", db)
	r.POST("/register", Register)
	r.POST("/login", Login)
	r.GET("/posts/:id", PostDetail)
	r.GET("/posts/list/:userb_id", ListPosts)
	r.GET("/posts/comments/:post_id", ListComments)
	// r.Use(AuthMiddleware())
	// 记录日志
	r.Use(gin.Logger())
	// 创建路由组
	auth := r.Group("/auth")
	auth.Use(AuthMiddleware())
	{
		auth.POST("/posts", CreatePost)
		auth.PUT("/posts/:id", UpdatePost)
		auth.DELETE("/posts/:id", DeletePost)
		auth.POST("/posts/comments", CreateComment)
	}
	return r
}

func Register(c *gin.Context) {
	var userb *model.Userb
	if err := c.ShouldBindJSON(&userb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userb.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFailedHashingPassword})
		return
	}
	userb.Password = string(hashedPassword)

	if err := handler.UserbCreate(db, userb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrUserCreate})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var userb *model.Userb
	if err := c.ShouldBindJSON(&userb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("registerRouters login userb : ", userb)
	var storedUser = model.Userb{}
	log.Println("registerRouters login db : ", db)
	if err := handler.UserbDetail(db, userb.ID, userb.Username, &storedUser); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUserNotFound})
		return
	}
	log.Println("registerRouters login storedUser : ", storedUser)
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(userb.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUserInvalidNameOrPassword})
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	// 剩下的逻辑...
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access: " + errors.ErrTokenIsNull})
			c.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):] // 移除 "Bearer " 前缀
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("your_secret_key"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access: " + errors.ErrTokenSignatureInvalid})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access: " + errors.ErrInvalidToken})
			c.Abort()
			return
		}
		// claims := token.Claims.(jwt.MapClaims)
		// userID := uint(claims["id"].(float64))
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userbID", claims["id"])
			c.Next()
		} else {
			c.Error(fmt.Errorf("invalid token"))
			c.Abort()
		}
	}
}

func CreatePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userb model.Userb
	if err := handler.UserbDetail(db, post.UserbID, "", &userb); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrUserNotFound})
		return
	}

	post.Userb = userb
	if err := handler.PostCreate(db, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrPostCreate})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

func UpdatePost(c *gin.Context) {
	// 验证用户权限
	userbID, _ := c.Get("userbID")
	log.Println("registerRouters UpdatePost userbID : ", userbID)
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// id , _:= strconv.ParseUint(c.Param("id"), 10, 64)
	// fmt.Println(id)
	// 验证用户权限
	log.Println("registerRouters UpdatePost post : ", post.UserbID)
	log.Println("registerRouters UpdatePost 比较 : ", (int)(post.UserbID) != (int)(userbID.(float64)))
	if (int)(post.UserbID) != (int)(userbID.(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUserIDNotMatchPostUserID})
		return
	}

	if err := handler.PostUpdate(db, c.Param("id"), &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func PostDetail(c *gin.Context) {
	var post = model.Post{}
	if err := handler.PostDetail(db, c.Param("id"), &post); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrPostNotFound})
		return
	}
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrPostNotFound})
		return
	}
	c.JSON(http.StatusOK, post)
}

func ListPosts(c *gin.Context) {
	var posts []model.Post
	if err := handler.PostList(db, &posts, c.Param("userb_id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func DeletePost(c *gin.Context) {
	// 验证用户权限
	id := c.Param("id")
	userbID, _ := c.Get("userbID")
	log.Println("registerRouters DeletePost userbID : ", userbID)
	// json里的post
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 数据库里的post
	var storePost = model.Post{}
	if err := handler.PostDetail(db, id, &storePost); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrPostNotFound})
		return
	}
	log.Println("registerRouters DeletePost post : ", post.UserbID)
	if (int)(post.UserbID) != (int)(userbID.(float64)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUserIDNotMatchPostUserID})
		return
	}
	if err := handler.PostDelete(db, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrPostDeleteFailed})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func CreateComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userb model.Userb
	if err := handler.UserbDetail(db, comment.UserbID, "", &userb); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrUserNotFound})
		return
	}

	var post = model.Post{}
	if err := handler.PostDetail(db, strconv.Itoa((int)(comment.PostID)), &post); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrPostNotFound})
		return
	}
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.ErrPostNotFound})
		return
	}

	if err := handler.CommentCreate(db, &comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrCommentCreate})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully"})
}

func ListComments(c *gin.Context) {
	var comments []model.Comment
	if err := handler.CommentList(db, &comments, c.Param("post_id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list comments"})
		return
	}
	c.JSON(http.StatusOK, comments)
}
