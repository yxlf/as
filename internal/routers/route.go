package routers

import (
	_ "as/docs"
	"as/internal/routers/api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	backend := r.Group("api")
	{
		adminControl := backend.Group("admin")
		{
			admin := api.NewAdmin()
			adminControl.POST("login", admin.Login)
			adminControl.POST("logout", admin.Logout)
			adminControl.POST("fresh", admin.Fresh)
			adminControl.GET("test", admin.Login)
			adminControl.GET("auth", api.GetAuth)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
