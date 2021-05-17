package api

import (
	"as/pkg/app"
	"github.com/gin-gonic/gin"
)

type Admin struct{}

// Login 用户登录
func (Admin) Login(c *gin.Context) {
	app.NewResponse(c).ToResponse("憨批一个")
	return
}

// Logout
// @Summary 用户登录
// @Produce json
// @Param name body string true "用户名"
// @param password body string true "密码"
// @Success 200 {object} errcode.Error "请求成功"
// @Failure 400 {object} errcode.Error "请求失败"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/api/admin/logout [get]
func (Admin) Logout(c *gin.Context) {

	app.NewResponse(c).ToResponse("sfa")
}

// Fresh 刷新token
func (Admin) Fresh(c *gin.Context) {

}

// GetInfo 获取用户信息
func (Admin) GetInfo(c *gin.Context) {

}
func NewAdmin() Admin {
	return Admin{}
}
