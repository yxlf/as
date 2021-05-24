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
		// 获取token
		backend.GET("admin/auth", api.GetAuth)

		//用户管理
		adminControl := backend.Group("admin") /*.Use(middleware.JWT())*/
		{
			{
				store := api.NewStore()
				// 新增门店
				adminControl.POST("store", store.Store)
				// 删除门店
				adminControl.DELETE("store/:id", store.Delete)
				// 获取门店
				adminControl.GET("store/:id", store.Get)
				// 门店列表
				adminControl.GET("store", store.Index)
				// 编辑门店
				adminControl.PUT("store/:id", store.Update)
			}
			{
				work := api.NewWork()
				adminControl.POST("work", work.Store)
				adminControl.GET("work", work.Index)
				adminControl.DELETE("work/:id", work.Delete)
			}
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
