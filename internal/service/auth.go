package service

import (
	"as/internal/dao"
	"errors"
)

type Service struct {
	dao dao.Dao
}

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc *Service) checkAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(
		param.AppKey,
		param.AppSecret,
	)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist.")
}
