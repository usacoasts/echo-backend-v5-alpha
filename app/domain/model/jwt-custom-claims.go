package model

import (
	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	Id    int  `json:"id"`
	Admin bool `json:"admin"`
	jwt.StandardClaims
}
