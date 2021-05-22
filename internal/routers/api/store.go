package api

import (
	"as/global"
	"as/internal/model"
	"as/pkg/app"
	"as/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Store struct {
}

func (s Store) Store(c *gin.Context) {
	response := app.NewResponse(c)
	var store model.Store
	err := c.ShouldBind(&store)
	if err != nil {
		response.ToErrorResponse(errcode.InvilidParams.WithDetails(err.Error()))
		return
	}
	err = global.DBEngine.Create(&store).Error
	if err != nil {
		response.ToErrorResponse(errcode.NewError(10008, "创建失败"))
		return
	}
	response.ToResponse("创建成功")

}
func NewStore() *Store {
	return &Store{}
}
