package api

import (
	"as/global"
	"as/internal/model"
	"as/pkg/app"
	"as/pkg/convert"
	"as/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Work struct {
}

func NewWork() *Work {
	return &Work{}
}

func (w Work) Store(c *gin.Context) {
	var work model.Work
	response := app.NewResponse(c)
	err := c.ShouldBind(&work)
	if err != nil {
		response.ToErrorResponse(errcode.InvilidParams.WithDetails(err.Error()))
		return
	}

	err = global.DBEngine.Create(&work).Error
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails(err.Error()))
		return
	}
	response.ToResponse("创建成功")
}

func (w Work) Index(c *gin.Context) {
	var works []*model.Work
	response := app.NewResponse(c)
	err := global.DBEngine.Find(&works).Error
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails(err.Error()))
		return
	}
	response.ToResponse(works)
}

func (w Work) Delete(c *gin.Context) {
	param := c.Param("id")
	response := app.NewResponse(c)
	if param == "" {
		response.ToErrorResponse(errcode.InvilidParams.WithDetails("参数错误"))
		return
	}
	res, _ := convert.StrTo(param).Uint64()
	work := model.Work{Model: &model.Model{ID: res}}
	err := global.DBEngine.Delete(work).Error
	if err != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails(err.Error()))
		return
	}
	response.ToResponse("删除成功")
}
