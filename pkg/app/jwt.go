package app

import (
	"github.com/golang-jwt/jwt/v4"
	global "github/h1deOnBush/dousheng/gloabal"
	"github/h1deOnBush/dousheng/pkg/errcode"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Id       int64  `json:"id"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// 根据用户名生成jwt
func GenToken(username string, id int64) (string, error) {
	now := time.Now()
	expire := now.Add(global.JWTSetting.Expire)
	claims := Claims{
		Username: username,
		Id:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	// 过期时间、签发者、时间等等的验证
	if ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, errcode.UnauthorizedTokenError
}
