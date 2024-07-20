package models

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTトークンのクレーム（ペイロード）を格納する構造体
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
