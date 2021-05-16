package routers

import (
	"as/internal/routers/api"
	"github.com/gin-gonic/gin"
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
		}
	}
	return r
}
