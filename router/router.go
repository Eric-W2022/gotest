package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"wxcloudrun-golang/docs"
	"wxcloudrun-golang/service"
)

// Router
// @securityDefinitions.basic  BasicAuth
func Router() *gin.Engine {
	r := gin.Default()

	// 主页
	r.GET("/", service.GetIndex)

	// 路由组
	v1 := r.Group("/api/v1")
	{
		// 用户组
		users := v1.Group("/user")
		{
			// 增删改查
			users.POST("", service.CreateUser)
			users.DELETE(":id", service.GetUserList)
			users.PATCH(":id", service.GetUserList)
			users.GET(":id", service.GetUserList)

			// 获取列表
			users.GET("", service.GetUserList)
			users.POST(":id/images", service.GetUserList)
		}

	}

	// swagger文档
	//docs.SwaggerInfo.Title = "测试"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = "localhost:80"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
