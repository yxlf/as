package api

import "github.com/gin-gonic/gin"

type Admin struct{}

// Login 用户登录
func (Admin) Login(c *gin.Context) {
	c.JSON(200, map[int]int{0: 0})
}

// Logout 用户登出
func (Admin) Logout(c *gin.Context) {

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
