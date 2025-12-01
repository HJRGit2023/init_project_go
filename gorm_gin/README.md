1、初始化项目：
#gorm、sqlite、mysql、glebarez/sqlite
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u github.com/glebarez/sqlite
go get gorm.io/driver/mysql
#gin
go get -u github.com/gin-gonic/gin
#jwt
go get github.com/dgrijalva/jwt-go
2、项目gorm_gin
----启动文件
--------main.go
----数据库初始化
--------database目录
----模型定义
--------model目录
----业务处理逻辑
--------handler目录
----路由注册
--------routers目录
----工具文件
--------util/errors目录
3、用户认证与授权
● 实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
● 使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。
4、文章管理功能
● 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
● 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
● 实现文章的更新功能，只有文章的作者才能更新自己的文章。
● 实现文章的删除功能，只有文章的作者才能删除自己的文章。
5、评论功能
  a. 实现评论的创建功能，已认证的用户可以对文章发表评论。
  b. 实现评论的读取功能，支持获取某篇文章的所有评论列表。
6、错误处理与日志记录
  a. 对可能出现的错误进行统一处理，如数据库连接错误、用户认证失败、文章或评论不存在等，返回合适的 HTTP 状态码和错误信息。
  b. 使用日志库记录系统的运行信息和错误信息，方便后续的调试和维护。
