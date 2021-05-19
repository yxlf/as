package api

import (
	"as/internal/service"
	"as/pkg/app"
	"as/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// GetAuth 获取token
// @Summary 获取token
// @Produce json
// @Param app_key body string true "key"
// @Param app_secret body string true "secret"
// @Success 200 {object} errcode.Error "请求成功"
// @Failure 500 {object} errcode.Error "服务器内部错误"
// @Router /api/admin/auth [get]
func GetAuth(c *gin.Context) {
	response := app.NewResponse(c)
	param := service.AuthRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		response.ToErrorResponse(errcode.InvilidParams.WithDetails(err.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.CheckAuth(&param)
	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist.WithDetails(err.Error()))
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{"token": token})
}
