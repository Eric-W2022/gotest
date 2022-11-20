package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"wxcloudrun-golang/docs"
	"wxcloudrun-golang/service"
	"wxcloudrun-golang/utils"
)

// Router
// @securityDefinitions.basic  BasicAuth
func Router() *gin.Engine {
	userAuthMiddleware := utils.UserJwt()
	r := gin.Default()
	// 主页
	r.GET("/", service.GetIndex)

	// 路由组
	v1 := r.Group("/api/v1")
	{
		// 注册
		v1.POST("/register", service.CreateUser)
		// 登录
		v1.POST("/login", userAuthMiddleware.LoginHandler)
		// 刷新
		v1.GET("/refresh_token", userAuthMiddleware.RefreshHandler)

		// 用户组
		users := v1.Group("/users")
		users.Use(userAuthMiddleware.MiddlewareFunc())
		{
			// 获取用户信息
			users.GET("", service.GetUserInfo)
			// 修改用户信息
			users.PATCH("", service.UpdateUser)
			// 删除用户
			users.DELETE("", service.DeleteUser)

			//users.POST(":id/images", service.GetUserList)
		}

	}

	// 空路由
	r.NoRoute(userAuthMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// swagger文档
	//docs.SwaggerInfo.Title = "测试"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = "localhost:80"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
