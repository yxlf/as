package service

type AdminLogin struct {
	UserName string `form:"user_name" binding:"required"`
	Password string `form:"password" binding:"required"`
}
