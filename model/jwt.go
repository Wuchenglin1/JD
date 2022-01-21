package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	User User
	Type string //"refreshToken"表示一个refreshToken,"token"表示一个token,"errToken"代表是错误的token
	Time time.Time
	jwt.StandardClaims
}

type Jwt struct {
	Sign string `json:"sign"`
}
