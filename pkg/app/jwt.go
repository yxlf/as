package app

import (
	"as/global"
	"as/internal/dao/encryption"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key,omitempty"`
	AppSecret string `json:"app_secret,omitempty"`
	jwt.StandardClaims
}

func GetJWTSecret() string {
	return global.JWTSetting.Secret
}

func GenerateToken(appKey, appSecret string) (string, error) {
	now := time.Now()
	expireTime := now.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    encryption.MD5(appKey),
		AppSecret: encryption.MD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(GetJWTSecret()))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetJWTSecret()), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
