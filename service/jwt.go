package service

import (
	"JD/model"
	"JD/tool"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func ParseRefreshToken(tokenStr string) (*model.Claims, error) {
	jwtCfg := tool.GetConfig().Jwt
	signKey := []byte(jwtCfg.Sign)
	token, err := jwt.ParseWithClaims(tokenStr, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		if claims.Type == "token" {
			errClaim := new(model.Claims)
			errClaim.Type = "errToken"
			return errClaim, nil
		}
		return claims, nil
	}
	return nil, err
}

func ParseAccessToken(tokenStr string) (*model.Claims, error) {
	jwtCfg := tool.GetConfig().Jwt
	signKey := []byte(jwtCfg.Sign)
	token, err := jwt.ParseWithClaims(tokenStr, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		if claims.Type == "refreshToken" {
			errClaim := new(model.Claims)
			errClaim.Type = "errToken"
			return errClaim, nil
		}
		return claims, nil
	}
	return nil, err
}

func CreateToken(u model.User, ExpireTime int64, tokenType string) (string, error) {
	jwtCfg := tool.GetConfig().Jwt
	signKey := []byte(jwtCfg.Sign)
	myClaims := model.Claims{
		User: u,
		Type: tokenType,
		Time: time.Now(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + ExpireTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	return token.SignedString(signKey)
}
