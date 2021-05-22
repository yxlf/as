package routers

import (
	_ "as/docs"
	"as/internal/middleware"
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
		// 获取token
		backend.GET("admin/auth", api.GetAuth)

		//用户管理
		adminControl := backend.Group("admin").Use(middleware.JWT())
		{
			// 新增门店
			store := api.NewStore()
			adminControl.POST("store", store.Store)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
